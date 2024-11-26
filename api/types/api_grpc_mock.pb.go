// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: api/api.proto

package types

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTakeHomeServiceClient is a mock of TakeHomeServiceClient interface.
type MockTakeHomeServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTakeHomeServiceClientMockRecorder
}

// MockTakeHomeServiceClientMockRecorder is the mock recorder for MockTakeHomeServiceClient.
type MockTakeHomeServiceClientMockRecorder struct {
	mock *MockTakeHomeServiceClient
}

// NewMockTakeHomeServiceClient creates a new mock instance.
func NewMockTakeHomeServiceClient(ctrl *gomock.Controller) *MockTakeHomeServiceClient {
	mock := &MockTakeHomeServiceClient{ctrl: ctrl}
	mock.recorder = &MockTakeHomeServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTakeHomeServiceClient) EXPECT() *MockTakeHomeServiceClientMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockTakeHomeServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateItem", varargs...)
	ret0, _ := ret[0].(*CreateItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockTakeHomeServiceClientMockRecorder) CreateItem(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockTakeHomeServiceClient)(nil).CreateItem), varargs...)
}

// GetItem mocks base method.
func (m *MockTakeHomeServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItem", varargs...)
	ret0, _ := ret[0].(*GetItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockTakeHomeServiceClientMockRecorder) GetItem(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockTakeHomeServiceClient)(nil).GetItem), varargs...)
}

// GetItems mocks base method.
func (m *MockTakeHomeServiceClient) GetItems(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItems", varargs...)
	ret0, _ := ret[0].(*GetItemsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockTakeHomeServiceClientMockRecorder) GetItems(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockTakeHomeServiceClient)(nil).GetItems), varargs...)
}

// MockTakeHomeServiceServer is a mock of TakeHomeServiceServer interface.
type MockTakeHomeServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTakeHomeServiceServerMockRecorder
}

// MockTakeHomeServiceServerMockRecorder is the mock recorder for MockTakeHomeServiceServer.
type MockTakeHomeServiceServerMockRecorder struct {
	mock *MockTakeHomeServiceServer
}

// NewMockTakeHomeServiceServer creates a new mock instance.
func NewMockTakeHomeServiceServer(ctrl *gomock.Controller) *MockTakeHomeServiceServer {
	mock := &MockTakeHomeServiceServer{ctrl: ctrl}
	mock.recorder = &MockTakeHomeServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTakeHomeServiceServer) EXPECT() *MockTakeHomeServiceServerMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockTakeHomeServiceServer) CreateItem(ctx context.Context, in *CreateItemRequest) (*CreateItemResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", ctx, in)
	ret0, _ := ret[0].(*CreateItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockTakeHomeServiceServerMockRecorder) CreateItem(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockTakeHomeServiceServer)(nil).CreateItem), ctx, in)
}

// GetItem mocks base method.
func (m *MockTakeHomeServiceServer) GetItem(ctx context.Context, in *GetItemRequest) (*GetItemResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, in)
	ret0, _ := ret[0].(*GetItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockTakeHomeServiceServerMockRecorder) GetItem(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockTakeHomeServiceServer)(nil).GetItem), ctx, in)
}

// GetItems mocks base method.
func (m *MockTakeHomeServiceServer) GetItems(ctx context.Context, in *EmptyRequest) (*GetItemsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ctx, in)
	ret0, _ := ret[0].(*GetItemsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockTakeHomeServiceServerMockRecorder) GetItems(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockTakeHomeServiceServer)(nil).GetItems), ctx, in)
}