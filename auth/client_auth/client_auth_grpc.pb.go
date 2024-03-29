// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: client_auth.proto

package client_auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClientAuthService_SignUp_FullMethodName           = "/auth.ClientAuthService/SignUp"
	ClientAuthService_SignIn_FullMethodName           = "/auth.ClientAuthService/SignIn"
	ClientAuthService_GetClientDetails_FullMethodName = "/auth.ClientAuthService/GetClientDetails"
)

// ClientAuthServiceClient is the client API for ClientAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientAuthServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	GetClientDetails(ctx context.Context, in *GetClientDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
}

type clientAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientAuthServiceClient(cc grpc.ClientConnInterface) ClientAuthServiceClient {
	return &clientAuthServiceClient{cc}
}

func (c *clientAuthServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientAuthService_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAuthServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, ClientAuthService_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientAuthServiceClient) GetClientDetails(ctx context.Context, in *GetClientDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	out := new(GetUserDetailsResponse)
	err := c.cc.Invoke(ctx, ClientAuthService_GetClientDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientAuthServiceServer is the server API for ClientAuthService service.
// All implementations must embed UnimplementedClientAuthServiceServer
// for forward compatibility
type ClientAuthServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*emptypb.Empty, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	GetClientDetails(context.Context, *GetClientDetailsRequest) (*GetUserDetailsResponse, error)
	mustEmbedUnimplementedClientAuthServiceServer()
}

// UnimplementedClientAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientAuthServiceServer struct {
}

func (UnimplementedClientAuthServiceServer) SignUp(context.Context, *SignUpRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedClientAuthServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedClientAuthServiceServer) GetClientDetails(context.Context, *GetClientDetailsRequest) (*GetUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientDetails not implemented")
}
func (UnimplementedClientAuthServiceServer) mustEmbedUnimplementedClientAuthServiceServer() {}

// UnsafeClientAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientAuthServiceServer will
// result in compilation errors.
type UnsafeClientAuthServiceServer interface {
	mustEmbedUnimplementedClientAuthServiceServer()
}

func RegisterClientAuthServiceServer(s grpc.ServiceRegistrar, srv ClientAuthServiceServer) {
	s.RegisterService(&ClientAuthService_ServiceDesc, srv)
}

func _ClientAuthService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAuthServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAuthService_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAuthServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAuthService_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAuthServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientAuthService_GetClientDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientAuthServiceServer).GetClientDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientAuthService_GetClientDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientAuthServiceServer).GetClientDetails(ctx, req.(*GetClientDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientAuthService_ServiceDesc is the grpc.ServiceDesc for ClientAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.ClientAuthService",
	HandlerType: (*ClientAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _ClientAuthService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _ClientAuthService_SignIn_Handler,
		},
		{
			MethodName: "GetClientDetails",
			Handler:    _ClientAuthService_GetClientDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_auth.proto",
}
