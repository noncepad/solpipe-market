// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: bot.proto

package bot

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
	Bot_Upload_FullMethodName    = "/bot.Bot/Upload"
	Bot_GetStatus_FullMethodName = "/bot.Bot/GetStatus"
	Bot_Start_FullMethodName     = "/bot.Bot/Start"
	Bot_Stop_FullMethodName      = "/bot.Bot/Stop"
)

// BotClient is the client API for Bot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[BotUploadRequest, BotStatus], error)
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BotStatus, error)
	Start(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Stdin, BotStatus], error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BotStatus, error)
}

type botClient struct {
	cc grpc.ClientConnInterface
}

func NewBotClient(cc grpc.ClientConnInterface) BotClient {
	return &botClient{cc}
}

func (c *botClient) Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[BotUploadRequest, BotStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Bot_ServiceDesc.Streams[0], Bot_Upload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BotUploadRequest, BotStatus]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Bot_UploadClient = grpc.ClientStreamingClient[BotUploadRequest, BotStatus]

func (c *botClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BotStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BotStatus)
	err := c.cc.Invoke(ctx, Bot_GetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) Start(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Stdin, BotStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Bot_ServiceDesc.Streams[1], Bot_Start_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Stdin, BotStatus]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Bot_StartClient = grpc.ClientStreamingClient[Stdin, BotStatus]

func (c *botClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BotStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BotStatus)
	err := c.cc.Invoke(ctx, Bot_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServer is the server API for Bot service.
// All implementations must embed UnimplementedBotServer
// for forward compatibility.
type BotServer interface {
	Upload(grpc.ClientStreamingServer[BotUploadRequest, BotStatus]) error
	GetStatus(context.Context, *Empty) (*BotStatus, error)
	Start(grpc.ClientStreamingServer[Stdin, BotStatus]) error
	Stop(context.Context, *Empty) (*BotStatus, error)
	mustEmbedUnimplementedBotServer()
}

// UnimplementedBotServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBotServer struct{}

func (UnimplementedBotServer) Upload(grpc.ClientStreamingServer[BotUploadRequest, BotStatus]) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedBotServer) GetStatus(context.Context, *Empty) (*BotStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedBotServer) Start(grpc.ClientStreamingServer[Stdin, BotStatus]) error {
	return status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedBotServer) Stop(context.Context, *Empty) (*BotStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedBotServer) mustEmbedUnimplementedBotServer() {}
func (UnimplementedBotServer) testEmbeddedByValue()             {}

// UnsafeBotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotServer will
// result in compilation errors.
type UnsafeBotServer interface {
	mustEmbedUnimplementedBotServer()
}

func RegisterBotServer(s grpc.ServiceRegistrar, srv BotServer) {
	// If the following call pancis, it indicates UnimplementedBotServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Bot_ServiceDesc, srv)
}

func _Bot_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotServer).Upload(&grpc.GenericServerStream[BotUploadRequest, BotStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Bot_UploadServer = grpc.ClientStreamingServer[BotUploadRequest, BotStatus]

func _Bot_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bot_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bot_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotServer).Start(&grpc.GenericServerStream[Stdin, BotStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Bot_StartServer = grpc.ClientStreamingServer[Stdin, BotStatus]

func _Bot_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bot_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Bot_ServiceDesc is the grpc.ServiceDesc for Bot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bot.Bot",
	HandlerType: (*BotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Bot_GetStatus_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Bot_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Bot_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Start",
			Handler:       _Bot_Start_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "bot.proto",
}