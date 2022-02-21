package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"

	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	v2refs "github.com/nspcc-dev/neofs-api-go/v2/refs"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/client"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	cid "github.com/nspcc-dev/neofs-sdk-go/container/id"
	"github.com/nspcc-dev/neofs-sdk-go/object"
	oid "github.com/nspcc-dev/neofs-sdk-go/object/id"
	"github.com/nspcc-dev/neofs-sdk-go/session"
	"github.com/nspcc-dev/neofs-sdk-go/token"
)

// PrmObjectSearch groups parameters of ObjectSearch operation.
type PrmObjectSearch struct {
	local bool

	sessionSet bool
	session    session.Token

	bearerSet bool
	bearer    token.BearerToken

	cnrSet bool
	cnr    cid.ID

	filters object.SearchFilters
}

// MarkLocal tells the server to execute the operation locally.
func (x *PrmObjectSearch) MarkLocal() {
	x.local = true
}

// WithinSession specifies session within which the search query must be executed.
//
// Creator of the session acquires the authorship of the request.
// This may affect the execution of an operation (e.g. access control).
//
// Must be signed.
func (x *PrmObjectSearch) WithinSession(t session.Token) {
	x.session = t
	x.sessionSet = true
}

// WithBearerToken attaches bearer token to be used for the operation.
//
// If set, underlying eACL rules will be used in access control.
//
// Must be signed.
func (x *PrmObjectSearch) WithBearerToken(t token.BearerToken) {
	x.bearer = t
	x.bearerSet = true
}

// InContainer specifies the container in which to look for objects.
// Required parameter.
func (x *PrmObjectSearch) InContainer(id cid.ID) {
	x.cnr = id
	x.cnrSet = true
}

// SetFilters sets filters by which to select objects. All container objects
// match unset/empty filters.
func (x *PrmObjectSearch) SetFilters(filters object.SearchFilters) {
	x.filters = filters
}

// ResObjectSearch groups the final result values of ObjectSearch operation.
type ResObjectSearch struct {
	statusRes
}

// ObjectListReader is designed to read list of object identifiers from NeoFS system.
//
// Must be initialized using Client.ObjectSearch, any other usage is unsafe.
type ObjectListReader struct {
	cancelCtxStream context.CancelFunc

	ctxCall contextCall

	reqWritten bool

	// initially bound to contextCall
	bodyResp v2object.SearchResponseBody

	tail []*v2refs.ObjectID
}

// UseKey specifies private key to sign the requests.
// If key is not provided, then Client default key is used.
func (x *ObjectListReader) UseKey(key ecdsa.PrivateKey) {
	x.ctxCall.key = key
}

// Read reads another list of the object identifiers. Works similar to
// io.Reader.Read but copies oid.ID and returns success flag instead of error.
//
// Failure reason can be received via Close.
//
// Panics if buf has zero length.
func (x *ObjectListReader) Read(buf []oid.ID) (int, bool) {
	if len(buf) == 0 {
		panic("empty buffer in ObjectListReader.ReadList")
	}

	if !x.reqWritten {
		if !x.ctxCall.writeRequest() {
			return 0, false
		}

		x.reqWritten = true
	}

	// read remaining tail
	read := len(x.tail)
	if read > len(buf) {
		read = len(buf)
	}

	for i := 0; i < read; i++ {
		buf[i] = *oid.NewIDFromV2(x.tail[i]) // need smth better
	}

	x.tail = x.tail[read:]

	if len(buf) == read {
		return read, true
	}

	// receive next message
	ok := x.ctxCall.readResponse()
	if !ok {
		return read, false
	}

	// read new chunk of objects
	ids := x.bodyResp.GetIDList()
	ln := len(ids)
	if ln == 0 {
		x.ctxCall.err = io.EOF
		return read, false
	}

	buf = buf[read:]
	if ln > len(buf) {
		ln = len(buf)
	}

	for i := 0; i < ln; i++ {
		buf[i] = *oid.NewIDFromV2(ids[i]) // need smth better
	}

	read += ln

	// save the tail
	x.tail = append(x.tail, ids[ln:]...)

	return read, true
}

// Close ends reading list of the matched objects and returns the result of the operation
// along with the final results. Must be called after using the ObjectListReader.
//
// Exactly one return value is non-nil. By default, server status is returned in res structure.
// Any client's internal or transport errors are returned as Go built-in error.
// If Client is tuned to resolve NeoFS API statuses, then NeoFS failures
// codes are returned as error.
//
// Return statuses:
//   global (see Client docs).
func (x *ObjectListReader) Close() (*ResObjectSearch, error) {
	defer x.cancelCtxStream()

	if x.ctxCall.err != nil && !errors.Is(x.ctxCall.err, io.EOF) {
		return nil, x.ctxCall.err
	}

	return x.ctxCall.statusRes.(*ResObjectSearch), nil
}

// ObjectSearchInit initiates object selection through a remote server using NeoFS API protocol.
//
// The call only opens the transmission channel, explicit fetching of matched objects
// is done using the ObjectListReader. Exactly one return value is non-nil.
// Resulting reader must be finally closed.
//
// Immediately panics if parameters are set incorrectly (see PrmObjectSearch docs).
// Context is required and must not be nil. It is used for network communication.
func (c *Client) ObjectSearchInit(ctx context.Context, prm PrmObjectSearch) (*ObjectListReader, error) {
	// check parameters
	switch {
	case ctx == nil:
		panic(panicMsgMissingContext)
	case !prm.cnrSet:
		panic(panicMsgMissingContainer)
	}

	// form request body
	var body v2object.SearchRequestBody

	body.SetVersion(1)
	body.SetContainerID(prm.cnr.ToV2())
	body.SetFilters(prm.filters.ToV2())

	// form meta header
	var meta v2session.RequestMetaHeader

	if prm.local {
		meta.SetTTL(1)
	}

	if prm.bearerSet {
		meta.SetBearerToken(prm.bearer.ToV2())
	}

	if prm.sessionSet {
		meta.SetSessionToken(prm.session.ToV2())
	}

	// form request
	var req v2object.SearchRequest

	req.SetBody(&body)
	req.SetMetaHeader(&meta)

	// init reader
	var (
		r      ObjectListReader
		resp   v2object.SearchResponse
		stream *rpcapi.SearchResponseReader
	)

	ctx, r.cancelCtxStream = context.WithCancel(ctx)

	resp.SetBody(&r.bodyResp)

	// init call context
	c.initCallContext(&r.ctxCall)
	r.ctxCall.req = &req
	r.ctxCall.statusRes = new(ResObjectSearch)
	r.ctxCall.resp = &resp
	r.ctxCall.wReq = func() error {
		var err error

		stream, err = rpcapi.SearchObjects(c.Raw(), &req, client.WithContext(ctx))
		if err != nil {
			return fmt.Errorf("open stream: %w", err)
		}

		return nil
	}
	r.ctxCall.rResp = func() error {
		return stream.Read(&resp)
	}

	return &r, nil
}