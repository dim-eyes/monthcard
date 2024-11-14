// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: api/monthCard/v1/monthCard.proto

package v1

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
	MonthCard_OpenMonthCard_FullMethodName     = "/api.monthCard.v1.MonthCard/OpenMonthCard"
	MonthCard_GetMonthCardRward_FullMethodName = "/api.monthCard.v1.MonthCard/GetMonthCardRward"
	MonthCard_GetMonthCardList_FullMethodName  = "/api.monthCard.v1.MonthCard/GetMonthCardList"
)

// MonthCardClient is the client API for MonthCard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonthCardClient interface {
	OpenMonthCard(ctx context.Context, in *OpenMonthCardRequest, opts ...grpc.CallOption) (*OpenMonthCardReply, error)
	GetMonthCardRward(ctx context.Context, in *GetMonthCardRewardRequest, opts ...grpc.CallOption) (*GetMonthCardRewardReply, error)
	GetMonthCardList(ctx context.Context, in *GetMonthCardListRequest, opts ...grpc.CallOption) (*GetMonthCardListReply, error)
}

type monthCardClient struct {
	cc grpc.ClientConnInterface
}

func NewMonthCardClient(cc grpc.ClientConnInterface) MonthCardClient {
	return &monthCardClient{cc}
}

func (c *monthCardClient) OpenMonthCard(ctx context.Context, in *OpenMonthCardRequest, opts ...grpc.CallOption) (*OpenMonthCardReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpenMonthCardReply)
	err := c.cc.Invoke(ctx, MonthCard_OpenMonthCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monthCardClient) GetMonthCardRward(ctx context.Context, in *GetMonthCardRewardRequest, opts ...grpc.CallOption) (*GetMonthCardRewardReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMonthCardRewardReply)
	err := c.cc.Invoke(ctx, MonthCard_GetMonthCardRward_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monthCardClient) GetMonthCardList(ctx context.Context, in *GetMonthCardListRequest, opts ...grpc.CallOption) (*GetMonthCardListReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMonthCardListReply)
	err := c.cc.Invoke(ctx, MonthCard_GetMonthCardList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonthCardServer is the server API for MonthCard service.
// All implementations must embed UnimplementedMonthCardServer
// for forward compatibility.
type MonthCardServer interface {
	OpenMonthCard(context.Context, *OpenMonthCardRequest) (*OpenMonthCardReply, error)
	GetMonthCardRward(context.Context, *GetMonthCardRewardRequest) (*GetMonthCardRewardReply, error)
	GetMonthCardList(context.Context, *GetMonthCardListRequest) (*GetMonthCardListReply, error)
	mustEmbedUnimplementedMonthCardServer()
}

// UnimplementedMonthCardServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMonthCardServer struct{}

func (UnimplementedMonthCardServer) OpenMonthCard(context.Context, *OpenMonthCardRequest) (*OpenMonthCardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenMonthCard not implemented")
}
func (UnimplementedMonthCardServer) GetMonthCardRward(context.Context, *GetMonthCardRewardRequest) (*GetMonthCardRewardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthCardRward not implemented")
}
func (UnimplementedMonthCardServer) GetMonthCardList(context.Context, *GetMonthCardListRequest) (*GetMonthCardListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthCardList not implemented")
}
func (UnimplementedMonthCardServer) mustEmbedUnimplementedMonthCardServer() {}
func (UnimplementedMonthCardServer) testEmbeddedByValue()                   {}

// UnsafeMonthCardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonthCardServer will
// result in compilation errors.
type UnsafeMonthCardServer interface {
	mustEmbedUnimplementedMonthCardServer()
}

func RegisterMonthCardServer(s grpc.ServiceRegistrar, srv MonthCardServer) {
	// If the following call pancis, it indicates UnimplementedMonthCardServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MonthCard_ServiceDesc, srv)
}

func _MonthCard_OpenMonthCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenMonthCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonthCardServer).OpenMonthCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonthCard_OpenMonthCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonthCardServer).OpenMonthCard(ctx, req.(*OpenMonthCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonthCard_GetMonthCardRward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonthCardRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonthCardServer).GetMonthCardRward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonthCard_GetMonthCardRward_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonthCardServer).GetMonthCardRward(ctx, req.(*GetMonthCardRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonthCard_GetMonthCardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonthCardListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonthCardServer).GetMonthCardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MonthCard_GetMonthCardList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonthCardServer).GetMonthCardList(ctx, req.(*GetMonthCardListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MonthCard_ServiceDesc is the grpc.ServiceDesc for MonthCard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MonthCard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.monthCard.v1.MonthCard",
	HandlerType: (*MonthCardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenMonthCard",
			Handler:    _MonthCard_OpenMonthCard_Handler,
		},
		{
			MethodName: "GetMonthCardRward",
			Handler:    _MonthCard_GetMonthCardRward_Handler,
		},
		{
			MethodName: "GetMonthCardList",
			Handler:    _MonthCard_GetMonthCardList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/monthCard/v1/monthCard.proto",
}