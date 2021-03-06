// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: api/sigma/v1/wechat_user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WechatUserClient is the client API for WechatUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WechatUserClient interface {
	CreateWechatUser(ctx context.Context, in *CreateWechatUserRequest, opts ...grpc.CallOption) (*CreateWechatUserReply, error)
	// rpc UpdateWechatUser(UpdateWechatUserRequest) returns (UpdateWechatUserReply);
	GetWechatUser(ctx context.Context, in *GetWechatUserRequest, opts ...grpc.CallOption) (*GetWechatUserReply, error)
	ListWechatUser(ctx context.Context, in *ListWechatUserRequest, opts ...grpc.CallOption) (*ListWechatUserReply, error)
}

type wechatUserClient struct {
	cc grpc.ClientConnInterface
}

func NewWechatUserClient(cc grpc.ClientConnInterface) WechatUserClient {
	return &wechatUserClient{cc}
}

func (c *wechatUserClient) CreateWechatUser(ctx context.Context, in *CreateWechatUserRequest, opts ...grpc.CallOption) (*CreateWechatUserReply, error) {
	out := new(CreateWechatUserReply)
	err := c.cc.Invoke(ctx, "/api.sigma.v1.WechatUser/CreateWechatUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatUserClient) GetWechatUser(ctx context.Context, in *GetWechatUserRequest, opts ...grpc.CallOption) (*GetWechatUserReply, error) {
	out := new(GetWechatUserReply)
	err := c.cc.Invoke(ctx, "/api.sigma.v1.WechatUser/GetWechatUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatUserClient) ListWechatUser(ctx context.Context, in *ListWechatUserRequest, opts ...grpc.CallOption) (*ListWechatUserReply, error) {
	out := new(ListWechatUserReply)
	err := c.cc.Invoke(ctx, "/api.sigma.v1.WechatUser/ListWechatUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WechatUserServer is the server API for WechatUser service.
// All implementations must embed UnimplementedWechatUserServer
// for forward compatibility
type WechatUserServer interface {
	CreateWechatUser(context.Context, *CreateWechatUserRequest) (*CreateWechatUserReply, error)
	// rpc UpdateWechatUser(UpdateWechatUserRequest) returns (UpdateWechatUserReply);
	GetWechatUser(context.Context, *GetWechatUserRequest) (*GetWechatUserReply, error)
	ListWechatUser(context.Context, *ListWechatUserRequest) (*ListWechatUserReply, error)
	mustEmbedUnimplementedWechatUserServer()
}

// UnimplementedWechatUserServer must be embedded to have forward compatible implementations.
type UnimplementedWechatUserServer struct {
}

func (UnimplementedWechatUserServer) CreateWechatUser(context.Context, *CreateWechatUserRequest) (*CreateWechatUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWechatUser not implemented")
}
func (UnimplementedWechatUserServer) GetWechatUser(context.Context, *GetWechatUserRequest) (*GetWechatUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWechatUser not implemented")
}
func (UnimplementedWechatUserServer) ListWechatUser(context.Context, *ListWechatUserRequest) (*ListWechatUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWechatUser not implemented")
}
func (UnimplementedWechatUserServer) mustEmbedUnimplementedWechatUserServer() {}

// UnsafeWechatUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WechatUserServer will
// result in compilation errors.
type UnsafeWechatUserServer interface {
	mustEmbedUnimplementedWechatUserServer()
}

func RegisterWechatUserServer(s grpc.ServiceRegistrar, srv WechatUserServer) {
	s.RegisterService(&WechatUser_ServiceDesc, srv)
}

func _WechatUser_CreateWechatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWechatUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatUserServer).CreateWechatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sigma.v1.WechatUser/CreateWechatUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatUserServer).CreateWechatUser(ctx, req.(*CreateWechatUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatUser_GetWechatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWechatUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatUserServer).GetWechatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sigma.v1.WechatUser/GetWechatUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatUserServer).GetWechatUser(ctx, req.(*GetWechatUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatUser_ListWechatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWechatUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatUserServer).ListWechatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sigma.v1.WechatUser/ListWechatUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatUserServer).ListWechatUser(ctx, req.(*ListWechatUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WechatUser_ServiceDesc is the grpc.ServiceDesc for WechatUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WechatUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sigma.v1.WechatUser",
	HandlerType: (*WechatUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWechatUser",
			Handler:    _WechatUser_CreateWechatUser_Handler,
		},
		{
			MethodName: "GetWechatUser",
			Handler:    _WechatUser_GetWechatUser_Handler,
		},
		{
			MethodName: "ListWechatUser",
			Handler:    _WechatUser_ListWechatUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sigma/v1/wechat_user.proto",
}
