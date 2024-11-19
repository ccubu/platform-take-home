// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/api.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TakeHomeService_GetItems_FullMethodName   = "/skip.platform.api.TakeHomeService/GetItems"
	TakeHomeService_GetItem_FullMethodName    = "/skip.platform.api.TakeHomeService/GetItem"
	TakeHomeService_CreateItem_FullMethodName = "/skip.platform.api.TakeHomeService/CreateItem"
)

// TakeHomeServiceClient is the client API for TakeHomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TakeHomeServiceClient interface {
	GetItems(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetItemsResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
}

type takeHomeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTakeHomeServiceClient(cc grpc.ClientConnInterface) TakeHomeServiceClient {
	return &takeHomeServiceClient{cc}
}

func (c *takeHomeServiceClient) GetItems(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, TakeHomeService_GetItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *takeHomeServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, TakeHomeService_GetItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *takeHomeServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, TakeHomeService_CreateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TakeHomeServiceServer is the server API for TakeHomeService service.
// All implementations must embed UnimplementedTakeHomeServiceServer
// for forward compatibility.
type TakeHomeServiceServer interface {
	GetItems(context.Context, *EmptyRequest) (*GetItemsResponse, error)
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	mustEmbedUnimplementedTakeHomeServiceServer()
}

// UnimplementedTakeHomeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTakeHomeServiceServer struct{}

func (UnimplementedTakeHomeServiceServer) GetItems(context.Context, *EmptyRequest) (*GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedTakeHomeServiceServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedTakeHomeServiceServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedTakeHomeServiceServer) mustEmbedUnimplementedTakeHomeServiceServer() {}
func (UnimplementedTakeHomeServiceServer) testEmbeddedByValue()                         {}

// UnsafeTakeHomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TakeHomeServiceServer will
// result in compilation errors.
type UnsafeTakeHomeServiceServer interface {
	mustEmbedUnimplementedTakeHomeServiceServer()
}

func RegisterTakeHomeServiceServer(s grpc.ServiceRegistrar, srv TakeHomeServiceServer) {
	// If the following call pancis, it indicates UnimplementedTakeHomeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TakeHomeService_ServiceDesc, srv)
}

func _TakeHomeService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TakeHomeServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TakeHomeService_GetItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TakeHomeServiceServer).GetItems(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TakeHomeService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TakeHomeServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TakeHomeService_GetItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TakeHomeServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TakeHomeService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TakeHomeServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TakeHomeService_CreateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TakeHomeServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TakeHomeService_ServiceDesc is the grpc.ServiceDesc for TakeHomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TakeHomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skip.platform.api.TakeHomeService",
	HandlerType: (*TakeHomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItems",
			Handler:    _TakeHomeService_GetItems_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _TakeHomeService_GetItem_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _TakeHomeService_CreateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
