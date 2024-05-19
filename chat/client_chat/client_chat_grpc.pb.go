// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: client_chat.proto

package client_chat

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
	ClientChatService_CreateRoom_FullMethodName           = "/chat.ClientChatService/CreateRoom"
	ClientChatService_GetListOfRooms_FullMethodName       = "/chat.ClientChatService/GetListOfRooms"
	ClientChatService_AddClientToRoom_FullMethodName      = "/chat.ClientChatService/AddClientToRoom"
	ClientChatService_RemoveClientFromRoom_FullMethodName = "/chat.ClientChatService/RemoveClientFromRoom"
)

// ClientChatServiceClient is the client API for ClientChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientChatServiceClient interface {
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	GetListOfRooms(ctx context.Context, in *GetListOfRoomsRequest, opts ...grpc.CallOption) (*GetListOfRoomsResponse, error)
	AddClientToRoom(ctx context.Context, in *AddClientToRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveClientFromRoom(ctx context.Context, in *RemoveClientFromRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientChatServiceClient(cc grpc.ClientConnInterface) ClientChatServiceClient {
	return &clientChatServiceClient{cc}
}

func (c *clientChatServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, ClientChatService_CreateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientChatServiceClient) GetListOfRooms(ctx context.Context, in *GetListOfRoomsRequest, opts ...grpc.CallOption) (*GetListOfRoomsResponse, error) {
	out := new(GetListOfRoomsResponse)
	err := c.cc.Invoke(ctx, ClientChatService_GetListOfRooms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientChatServiceClient) AddClientToRoom(ctx context.Context, in *AddClientToRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientChatService_AddClientToRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientChatServiceClient) RemoveClientFromRoom(ctx context.Context, in *RemoveClientFromRoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ClientChatService_RemoveClientFromRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientChatServiceServer is the server API for ClientChatService service.
// All implementations must embed UnimplementedClientChatServiceServer
// for forward compatibility
type ClientChatServiceServer interface {
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	GetListOfRooms(context.Context, *GetListOfRoomsRequest) (*GetListOfRoomsResponse, error)
	AddClientToRoom(context.Context, *AddClientToRoomRequest) (*emptypb.Empty, error)
	RemoveClientFromRoom(context.Context, *RemoveClientFromRoomRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientChatServiceServer()
}

// UnimplementedClientChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientChatServiceServer struct {
}

func (UnimplementedClientChatServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedClientChatServiceServer) GetListOfRooms(context.Context, *GetListOfRoomsRequest) (*GetListOfRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfRooms not implemented")
}
func (UnimplementedClientChatServiceServer) AddClientToRoom(context.Context, *AddClientToRoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClientToRoom not implemented")
}
func (UnimplementedClientChatServiceServer) RemoveClientFromRoom(context.Context, *RemoveClientFromRoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveClientFromRoom not implemented")
}
func (UnimplementedClientChatServiceServer) mustEmbedUnimplementedClientChatServiceServer() {}

// UnsafeClientChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientChatServiceServer will
// result in compilation errors.
type UnsafeClientChatServiceServer interface {
	mustEmbedUnimplementedClientChatServiceServer()
}

func RegisterClientChatServiceServer(s grpc.ServiceRegistrar, srv ClientChatServiceServer) {
	s.RegisterService(&ClientChatService_ServiceDesc, srv)
}

func _ClientChatService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientChatServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientChatService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientChatServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientChatService_GetListOfRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOfRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientChatServiceServer).GetListOfRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientChatService_GetListOfRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientChatServiceServer).GetListOfRooms(ctx, req.(*GetListOfRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientChatService_AddClientToRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClientToRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientChatServiceServer).AddClientToRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientChatService_AddClientToRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientChatServiceServer).AddClientToRoom(ctx, req.(*AddClientToRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientChatService_RemoveClientFromRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveClientFromRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientChatServiceServer).RemoveClientFromRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientChatService_RemoveClientFromRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientChatServiceServer).RemoveClientFromRoom(ctx, req.(*RemoveClientFromRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientChatService_ServiceDesc is the grpc.ServiceDesc for ClientChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ClientChatService",
	HandlerType: (*ClientChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _ClientChatService_CreateRoom_Handler,
		},
		{
			MethodName: "GetListOfRooms",
			Handler:    _ClientChatService_GetListOfRooms_Handler,
		},
		{
			MethodName: "AddClientToRoom",
			Handler:    _ClientChatService_AddClientToRoom_Handler,
		},
		{
			MethodName: "RemoveClientFromRoom",
			Handler:    _ClientChatService_RemoveClientFromRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_chat.proto",
}
