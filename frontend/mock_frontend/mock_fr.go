// Code generated by MockGen. DO NOT EDIT.
// Source: edgekv/frontend/frontend (interfaces: FrontendClient)

// Package mock_frontend is a generated GoMock package.
package mock_frontend

import (
	context "context"
	frontend "edgekv/frontend/frontend"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockFrontendClient is a mock of FrontendClient interface
type MockFrontendClient struct {
	ctrl     *gomock.Controller
	recorder *MockFrontendClientMockRecorder
}

// MockFrontendClientMockRecorder is the mock recorder for MockFrontendClient
type MockFrontendClientMockRecorder struct {
	mock *MockFrontendClient
}

// NewMockFrontendClient creates a new mock instance
func NewMockFrontendClient(ctrl *gomock.Controller) *MockFrontendClient {
	mock := &MockFrontendClient{ctrl: ctrl}
	mock.recorder = &MockFrontendClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFrontendClient) EXPECT() *MockFrontendClientMockRecorder {
	return m.recorder
}

// Del mocks base method
func (m *MockFrontendClient) Del(arg0 context.Context, arg1 *frontend.DeleteRequest, arg2 ...grpc.CallOption) (*frontend.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*frontend.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Del indicates an expected call of Del
func (mr *MockFrontendClientMockRecorder) Del(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockFrontendClient)(nil).Del), varargs...)
}

// Get mocks base method
func (m *MockFrontendClient) Get(arg0 context.Context, arg1 *frontend.GetRequest, arg2 ...grpc.CallOption) (*frontend.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*frontend.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockFrontendClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFrontendClient)(nil).Get), varargs...)
}

// Put mocks base method
func (m *MockFrontendClient) Put(arg0 context.Context, arg1 *frontend.PutRequest, arg2 ...grpc.CallOption) (*frontend.PutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*frontend.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockFrontendClientMockRecorder) Put(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockFrontendClient)(nil).Put), varargs...)
}
