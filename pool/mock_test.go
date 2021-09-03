// Code generated by MockGen. DO NOT EDIT.
// Source: ./pool.go

// Package pool is a generated GoMock package.
package pool

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	accounting "github.com/nspcc-dev/neofs-api-go/pkg/accounting"
	eacl "github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	client "github.com/nspcc-dev/neofs-api-go/pkg/client"
	container "github.com/nspcc-dev/neofs-api-go/pkg/container"
	id "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	netmap "github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	object "github.com/nspcc-dev/neofs-api-go/pkg/object"
	owner "github.com/nspcc-dev/neofs-api-go/pkg/owner"
	session "github.com/nspcc-dev/neofs-api-go/pkg/session"
	client0 "github.com/nspcc-dev/neofs-api-go/rpc/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AnnounceContainerUsedSpace mocks base method.
func (m *MockClient) AnnounceContainerUsedSpace(arg0 context.Context, arg1 []container.UsedSpaceAnnouncement, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AnnounceContainerUsedSpace", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnounceContainerUsedSpace indicates an expected call of AnnounceContainerUsedSpace.
func (mr *MockClientMockRecorder) AnnounceContainerUsedSpace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceContainerUsedSpace", reflect.TypeOf((*MockClient)(nil).AnnounceContainerUsedSpace), varargs...)
}

// AnnounceIntermediateTrust mocks base method.
func (m *MockClient) AnnounceIntermediateTrust(arg0 context.Context, arg1 client.AnnounceIntermediateTrustPrm, arg2 ...client.CallOption) (*client.AnnounceIntermediateTrustRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AnnounceIntermediateTrust", varargs...)
	ret0, _ := ret[0].(*client.AnnounceIntermediateTrustRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceIntermediateTrust indicates an expected call of AnnounceIntermediateTrust.
func (mr *MockClientMockRecorder) AnnounceIntermediateTrust(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceIntermediateTrust", reflect.TypeOf((*MockClient)(nil).AnnounceIntermediateTrust), varargs...)
}

// AnnounceLocalTrust mocks base method.
func (m *MockClient) AnnounceLocalTrust(arg0 context.Context, arg1 client.AnnounceLocalTrustPrm, arg2 ...client.CallOption) (*client.AnnounceLocalTrustRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AnnounceLocalTrust", varargs...)
	ret0, _ := ret[0].(*client.AnnounceLocalTrustRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceLocalTrust indicates an expected call of AnnounceLocalTrust.
func (mr *MockClientMockRecorder) AnnounceLocalTrust(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceLocalTrust", reflect.TypeOf((*MockClient)(nil).AnnounceLocalTrust), varargs...)
}

// Conn mocks base method.
func (m *MockClient) Conn() io.Closer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(io.Closer)
	return ret0
}

// Conn indicates an expected call of Conn.
func (mr *MockClientMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockClient)(nil).Conn))
}

// CreateSession mocks base method.
func (m *MockClient) CreateSession(arg0 context.Context, arg1 uint64, arg2 ...client.CallOption) (*session.Token, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSession", varargs...)
	ret0, _ := ret[0].(*session.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockClientMockRecorder) CreateSession(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockClient)(nil).CreateSession), varargs...)
}

// DeleteContainer mocks base method.
func (m *MockClient) DeleteContainer(arg0 context.Context, arg1 *id.ID, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteContainer", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContainer indicates an expected call of DeleteContainer.
func (mr *MockClientMockRecorder) DeleteContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockClient)(nil).DeleteContainer), varargs...)
}

// DeleteObject mocks base method.
func (m *MockClient) DeleteObject(arg0 context.Context, arg1 *client.DeleteObjectParams, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObject", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockClientMockRecorder) DeleteObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockClient)(nil).DeleteObject), varargs...)
}

// EndpointInfo mocks base method.
func (m *MockClient) EndpointInfo(arg0 context.Context, arg1 ...client.CallOption) (*client.EndpointInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EndpointInfo", varargs...)
	ret0, _ := ret[0].(*client.EndpointInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndpointInfo indicates an expected call of EndpointInfo.
func (mr *MockClientMockRecorder) EndpointInfo(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointInfo", reflect.TypeOf((*MockClient)(nil).EndpointInfo), varargs...)
}

// GetBalance mocks base method.
func (m *MockClient) GetBalance(arg0 context.Context, arg1 *owner.ID, arg2 ...client.CallOption) (*accounting.Decimal, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBalance", varargs...)
	ret0, _ := ret[0].(*accounting.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockClientMockRecorder) GetBalance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockClient)(nil).GetBalance), varargs...)
}

// GetContainer mocks base method.
func (m *MockClient) GetContainer(arg0 context.Context, arg1 *id.ID, arg2 ...client.CallOption) (*container.Container, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContainer", varargs...)
	ret0, _ := ret[0].(*container.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainer indicates an expected call of GetContainer.
func (mr *MockClientMockRecorder) GetContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainer", reflect.TypeOf((*MockClient)(nil).GetContainer), varargs...)
}

// GetEACL mocks base method.
func (m *MockClient) GetEACL(arg0 context.Context, arg1 *id.ID, arg2 ...client.CallOption) (*client.EACLWithSignature, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEACL", varargs...)
	ret0, _ := ret[0].(*client.EACLWithSignature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEACL indicates an expected call of GetEACL.
func (mr *MockClientMockRecorder) GetEACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEACL", reflect.TypeOf((*MockClient)(nil).GetEACL), varargs...)
}

// GetObject mocks base method.
func (m *MockClient) GetObject(arg0 context.Context, arg1 *client.GetObjectParams, arg2 ...client.CallOption) (*object.Object, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*object.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockClientMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockClient)(nil).GetObject), varargs...)
}

// GetObjectHeader mocks base method.
func (m *MockClient) GetObjectHeader(arg0 context.Context, arg1 *client.ObjectHeaderParams, arg2 ...client.CallOption) (*object.Object, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectHeader", varargs...)
	ret0, _ := ret[0].(*object.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectHeader indicates an expected call of GetObjectHeader.
func (mr *MockClientMockRecorder) GetObjectHeader(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectHeader", reflect.TypeOf((*MockClient)(nil).GetObjectHeader), varargs...)
}

// ListContainers mocks base method.
func (m *MockClient) ListContainers(arg0 context.Context, arg1 *owner.ID, arg2 ...client.CallOption) ([]*id.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainers", varargs...)
	ret0, _ := ret[0].([]*id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockClientMockRecorder) ListContainers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockClient)(nil).ListContainers), varargs...)
}

// NetworkInfo mocks base method.
func (m *MockClient) NetworkInfo(arg0 context.Context, arg1 ...client.CallOption) (*netmap.NetworkInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NetworkInfo", varargs...)
	ret0, _ := ret[0].(*netmap.NetworkInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkInfo indicates an expected call of NetworkInfo.
func (mr *MockClientMockRecorder) NetworkInfo(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkInfo", reflect.TypeOf((*MockClient)(nil).NetworkInfo), varargs...)
}

// ObjectPayloadRangeData mocks base method.
func (m *MockClient) ObjectPayloadRangeData(arg0 context.Context, arg1 *client.RangeDataParams, arg2 ...client.CallOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectPayloadRangeData", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPayloadRangeData indicates an expected call of ObjectPayloadRangeData.
func (mr *MockClientMockRecorder) ObjectPayloadRangeData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeData", reflect.TypeOf((*MockClient)(nil).ObjectPayloadRangeData), varargs...)
}

// ObjectPayloadRangeSHA256 mocks base method.
func (m *MockClient) ObjectPayloadRangeSHA256(arg0 context.Context, arg1 *client.RangeChecksumParams, arg2 ...client.CallOption) ([][32]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectPayloadRangeSHA256", varargs...)
	ret0, _ := ret[0].([][32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPayloadRangeSHA256 indicates an expected call of ObjectPayloadRangeSHA256.
func (mr *MockClientMockRecorder) ObjectPayloadRangeSHA256(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeSHA256", reflect.TypeOf((*MockClient)(nil).ObjectPayloadRangeSHA256), varargs...)
}

// ObjectPayloadRangeTZ mocks base method.
func (m *MockClient) ObjectPayloadRangeTZ(arg0 context.Context, arg1 *client.RangeChecksumParams, arg2 ...client.CallOption) ([][64]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectPayloadRangeTZ", varargs...)
	ret0, _ := ret[0].([][64]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPayloadRangeTZ indicates an expected call of ObjectPayloadRangeTZ.
func (mr *MockClientMockRecorder) ObjectPayloadRangeTZ(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeTZ", reflect.TypeOf((*MockClient)(nil).ObjectPayloadRangeTZ), varargs...)
}

// PutContainer mocks base method.
func (m *MockClient) PutContainer(arg0 context.Context, arg1 *container.Container, arg2 ...client.CallOption) (*id.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutContainer", varargs...)
	ret0, _ := ret[0].(*id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutContainer indicates an expected call of PutContainer.
func (mr *MockClientMockRecorder) PutContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContainer", reflect.TypeOf((*MockClient)(nil).PutContainer), varargs...)
}

// PutObject mocks base method.
func (m *MockClient) PutObject(arg0 context.Context, arg1 *client.PutObjectParams, arg2 ...client.CallOption) (*object.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*object.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockClientMockRecorder) PutObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockClient)(nil).PutObject), varargs...)
}

// Raw mocks base method.
func (m *MockClient) Raw() *client0.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(*client0.Client)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *MockClientMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockClient)(nil).Raw))
}

// SearchObject mocks base method.
func (m *MockClient) SearchObject(arg0 context.Context, arg1 *client.SearchObjectParams, arg2 ...client.CallOption) ([]*object.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchObject", varargs...)
	ret0, _ := ret[0].([]*object.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchObject indicates an expected call of SearchObject.
func (mr *MockClientMockRecorder) SearchObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchObject", reflect.TypeOf((*MockClient)(nil).SearchObject), varargs...)
}

// SetEACL mocks base method.
func (m *MockClient) SetEACL(arg0 context.Context, arg1 *eacl.Table, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetEACL", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEACL indicates an expected call of SetEACL.
func (mr *MockClientMockRecorder) SetEACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEACL", reflect.TypeOf((*MockClient)(nil).SetEACL), varargs...)
}

// MockPool is a mock of Pool interface.
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool.
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance.
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// AnnounceContainerUsedSpace mocks base method.
func (m *MockPool) AnnounceContainerUsedSpace(arg0 context.Context, arg1 []container.UsedSpaceAnnouncement, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AnnounceContainerUsedSpace", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnounceContainerUsedSpace indicates an expected call of AnnounceContainerUsedSpace.
func (mr *MockPoolMockRecorder) AnnounceContainerUsedSpace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceContainerUsedSpace", reflect.TypeOf((*MockPool)(nil).AnnounceContainerUsedSpace), varargs...)
}

// Connection mocks base method.
func (m *MockPool) Connection() (client.Client, *session.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connection")
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(*session.Token)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Connection indicates an expected call of Connection.
func (mr *MockPoolMockRecorder) Connection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connection", reflect.TypeOf((*MockPool)(nil).Connection))
}

// DeleteContainer mocks base method.
func (m *MockPool) DeleteContainer(arg0 context.Context, arg1 *id.ID, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteContainer", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContainer indicates an expected call of DeleteContainer.
func (mr *MockPoolMockRecorder) DeleteContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockPool)(nil).DeleteContainer), varargs...)
}

// DeleteObject mocks base method.
func (m *MockPool) DeleteObject(arg0 context.Context, arg1 *client.DeleteObjectParams, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObject", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockPoolMockRecorder) DeleteObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockPool)(nil).DeleteObject), varargs...)
}

// GetContainer mocks base method.
func (m *MockPool) GetContainer(arg0 context.Context, arg1 *id.ID, arg2 ...client.CallOption) (*container.Container, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContainer", varargs...)
	ret0, _ := ret[0].(*container.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainer indicates an expected call of GetContainer.
func (mr *MockPoolMockRecorder) GetContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainer", reflect.TypeOf((*MockPool)(nil).GetContainer), varargs...)
}

// GetEACL mocks base method.
func (m *MockPool) GetEACL(arg0 context.Context, arg1 *id.ID, arg2 ...client.CallOption) (*client.EACLWithSignature, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEACL", varargs...)
	ret0, _ := ret[0].(*client.EACLWithSignature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEACL indicates an expected call of GetEACL.
func (mr *MockPoolMockRecorder) GetEACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEACL", reflect.TypeOf((*MockPool)(nil).GetEACL), varargs...)
}

// GetObject mocks base method.
func (m *MockPool) GetObject(arg0 context.Context, arg1 *client.GetObjectParams, arg2 ...client.CallOption) (*object.Object, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*object.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockPoolMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockPool)(nil).GetObject), varargs...)
}

// GetObjectHeader mocks base method.
func (m *MockPool) GetObjectHeader(arg0 context.Context, arg1 *client.ObjectHeaderParams, arg2 ...client.CallOption) (*object.Object, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectHeader", varargs...)
	ret0, _ := ret[0].(*object.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectHeader indicates an expected call of GetObjectHeader.
func (mr *MockPoolMockRecorder) GetObjectHeader(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectHeader", reflect.TypeOf((*MockPool)(nil).GetObjectHeader), varargs...)
}

// ListContainers mocks base method.
func (m *MockPool) ListContainers(arg0 context.Context, arg1 *owner.ID, arg2 ...client.CallOption) ([]*id.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainers", varargs...)
	ret0, _ := ret[0].([]*id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockPoolMockRecorder) ListContainers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockPool)(nil).ListContainers), varargs...)
}

// ObjectPayloadRangeData mocks base method.
func (m *MockPool) ObjectPayloadRangeData(arg0 context.Context, arg1 *client.RangeDataParams, arg2 ...client.CallOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectPayloadRangeData", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPayloadRangeData indicates an expected call of ObjectPayloadRangeData.
func (mr *MockPoolMockRecorder) ObjectPayloadRangeData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeData", reflect.TypeOf((*MockPool)(nil).ObjectPayloadRangeData), varargs...)
}

// ObjectPayloadRangeSHA256 mocks base method.
func (m *MockPool) ObjectPayloadRangeSHA256(arg0 context.Context, arg1 *client.RangeChecksumParams, arg2 ...client.CallOption) ([][32]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectPayloadRangeSHA256", varargs...)
	ret0, _ := ret[0].([][32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPayloadRangeSHA256 indicates an expected call of ObjectPayloadRangeSHA256.
func (mr *MockPoolMockRecorder) ObjectPayloadRangeSHA256(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeSHA256", reflect.TypeOf((*MockPool)(nil).ObjectPayloadRangeSHA256), varargs...)
}

// ObjectPayloadRangeTZ mocks base method.
func (m *MockPool) ObjectPayloadRangeTZ(arg0 context.Context, arg1 *client.RangeChecksumParams, arg2 ...client.CallOption) ([][64]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectPayloadRangeTZ", varargs...)
	ret0, _ := ret[0].([][64]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPayloadRangeTZ indicates an expected call of ObjectPayloadRangeTZ.
func (mr *MockPoolMockRecorder) ObjectPayloadRangeTZ(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeTZ", reflect.TypeOf((*MockPool)(nil).ObjectPayloadRangeTZ), varargs...)
}

// OwnerID mocks base method.
func (m *MockPool) OwnerID() *owner.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OwnerID")
	ret0, _ := ret[0].(*owner.ID)
	return ret0
}

// OwnerID indicates an expected call of OwnerID.
func (mr *MockPoolMockRecorder) OwnerID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OwnerID", reflect.TypeOf((*MockPool)(nil).OwnerID))
}

// PutContainer mocks base method.
func (m *MockPool) PutContainer(arg0 context.Context, arg1 *container.Container, arg2 ...client.CallOption) (*id.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutContainer", varargs...)
	ret0, _ := ret[0].(*id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutContainer indicates an expected call of PutContainer.
func (mr *MockPoolMockRecorder) PutContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContainer", reflect.TypeOf((*MockPool)(nil).PutContainer), varargs...)
}

// PutObject mocks base method.
func (m *MockPool) PutObject(arg0 context.Context, arg1 *client.PutObjectParams, arg2 ...client.CallOption) (*object.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*object.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockPoolMockRecorder) PutObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockPool)(nil).PutObject), varargs...)
}

// SearchObject mocks base method.
func (m *MockPool) SearchObject(arg0 context.Context, arg1 *client.SearchObjectParams, arg2 ...client.CallOption) ([]*object.ID, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchObject", varargs...)
	ret0, _ := ret[0].([]*object.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchObject indicates an expected call of SearchObject.
func (mr *MockPoolMockRecorder) SearchObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchObject", reflect.TypeOf((*MockPool)(nil).SearchObject), varargs...)
}

// SetEACL mocks base method.
func (m *MockPool) SetEACL(arg0 context.Context, arg1 *eacl.Table, arg2 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetEACL", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEACL indicates an expected call of SetEACL.
func (mr *MockPoolMockRecorder) SetEACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEACL", reflect.TypeOf((*MockPool)(nil).SetEACL), varargs...)
}

// WaitForContainerPresence mocks base method.
func (m *MockPool) WaitForContainerPresence(arg0 context.Context, arg1 *id.ID, arg2 *ContainerPollingParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForContainerPresence", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForContainerPresence indicates an expected call of WaitForContainerPresence.
func (mr *MockPoolMockRecorder) WaitForContainerPresence(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForContainerPresence", reflect.TypeOf((*MockPool)(nil).WaitForContainerPresence), arg0, arg1, arg2)
}