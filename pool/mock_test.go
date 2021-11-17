// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nspcc-dev/neofs-sdk-go/pool (interfaces: Client)

// Package pool is a generated GoMock package.
package pool

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client0 "github.com/nspcc-dev/neofs-sdk-go/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	client0.Client
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

// AnnounceContainerUsedSpace indicates an expected call of AnnounceContainerUsedSpace.
func (mr *MockClientMockRecorder) AnnounceContainerUsedSpace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceContainerUsedSpace", reflect.TypeOf((*MockClient)(nil).AnnounceContainerUsedSpace), varargs...)
}

// AnnounceIntermediateTrust indicates an expected call of AnnounceIntermediateTrust.
func (mr *MockClientMockRecorder) AnnounceIntermediateTrust(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceIntermediateTrust", reflect.TypeOf((*MockClient)(nil).AnnounceIntermediateTrust), varargs...)
}

// AnnounceLocalTrust indicates an expected call of AnnounceLocalTrust.
func (mr *MockClientMockRecorder) AnnounceLocalTrust(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceLocalTrust", reflect.TypeOf((*MockClient)(nil).AnnounceLocalTrust), varargs...)
}

// Conn indicates an expected call of Conn.
func (mr *MockClientMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockClient)(nil).Conn))
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockClientMockRecorder) CreateSession(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockClient)(nil).CreateSession), varargs...)
}

// DeleteContainer indicates an expected call of DeleteContainer.
func (mr *MockClientMockRecorder) DeleteContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockClient)(nil).DeleteContainer), varargs...)
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockClientMockRecorder) DeleteObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockClient)(nil).DeleteObject), varargs...)
}

// EndpointInfo indicates an expected call of EndpointInfo.
func (mr *MockClientMockRecorder) EndpointInfo(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointInfo", reflect.TypeOf((*MockClient)(nil).EndpointInfo), varargs...)
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockClientMockRecorder) GetBalance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockClient)(nil).GetBalance), varargs...)
}

// GetContainer indicates an expected call of GetContainer.
func (mr *MockClientMockRecorder) GetContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainer", reflect.TypeOf((*MockClient)(nil).GetContainer), varargs...)
}

// GetEACL indicates an expected call of GetEACL.
func (mr *MockClientMockRecorder) GetEACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EACL", reflect.TypeOf((*MockClient)(nil).EACL), varargs...)
}

// GetObject indicates an expected call of GetObject.
func (mr *MockClientMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockClient)(nil).GetObject), varargs...)
}

// GetObjectHeader indicates an expected call of GetObjectHeader.
func (mr *MockClientMockRecorder) GetObjectHeader(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadObject", reflect.TypeOf((*MockClient)(nil).HeadObject), varargs...)
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockClientMockRecorder) ListContainers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockClient)(nil).ListContainers), varargs...)
}

// NetworkInfo indicates an expected call of NetworkInfo.
func (mr *MockClientMockRecorder) NetworkInfo(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkInfo", reflect.TypeOf((*MockClient)(nil).NetworkInfo), varargs...)
}

// ObjectPayloadRangeData indicates an expected call of ObjectPayloadRangeData.
func (mr *MockClientMockRecorder) ObjectPayloadRangeData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeData", reflect.TypeOf((*MockClient)(nil).ObjectPayloadRangeData), varargs...)
}

// ObjectPayloadRangeSHA256 indicates an expected call of ObjectPayloadRangeSHA256.
func (mr *MockClientMockRecorder) ObjectPayloadRangeSHA256(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashObjectPayloadRanges", reflect.TypeOf((*MockClient)(nil).HashObjectPayloadRanges), varargs...)
}

// ObjectPayloadRangeTZ indicates an expected call of ObjectPayloadRangeTZ.
func (mr *MockClientMockRecorder) ObjectPayloadRangeTZ(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashObjectPayloadRanges", reflect.TypeOf((*MockClient)(nil).HashObjectPayloadRanges), varargs...)
}

// PutContainer indicates an expected call of PutContainer.
func (mr *MockClientMockRecorder) PutContainer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContainer", reflect.TypeOf((*MockClient)(nil).PutContainer), varargs...)
}

// PutObject indicates an expected call of PutObject.
func (mr *MockClientMockRecorder) PutObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockClient)(nil).PutObject), varargs...)
}

// Raw indicates an expected call of Raw.
func (mr *MockClientMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockClient)(nil).Raw))
}

// SearchObject indicates an expected call of SearchObject.
func (mr *MockClientMockRecorder) SearchObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchObjects", reflect.TypeOf((*MockClient)(nil).SearchObjects), varargs...)
}

// SetEACL indicates an expected call of SetEACL.
func (mr *MockClientMockRecorder) SetEACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEACL", reflect.TypeOf((*MockClient)(nil).SetEACL), varargs...)
}
