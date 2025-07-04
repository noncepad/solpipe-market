// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bot_Upload_FullMethodName    = "/bot.Bot/Upload"
	Bot_GetStatus_FullMethodName = "/bot.Bot/GetStatus"
	Bot_Run_FullMethodName       = "/bot.Bot/Run"
)

// BotClient is the client API for Bot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotClient interface {
	// Upload a bot. If more than N bots exist, the oldest one will be overwritten and deleted.
	Upload(ctx context.Context, opts ...grpc.CallOption) (Bot_UploadClient, error)
	// Get a list of currently uploaded bots. There is a limit of N.
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	// The bot must have already been uploaded.  The Bot is frozen upon disconnect.
	Run(ctx context.Context, opts ...grpc.CallOption) (Bot_RunClient, error)
}

type botClient struct {
	cc grpc.ClientConnInterface
}

func NewBotClient(cc grpc.ClientConnInterface) BotClient {
	return &botClient{cc}
}

func (c *botClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Bot_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bot_ServiceDesc.Streams[0], Bot_Upload_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &botUploadClient{stream}
	return x, nil
}

type Bot_UploadClient interface {
	Send(*BotUploadRequest) error
	CloseAndRecv() (*BotStatus, error)
	grpc.ClientStream
}

type botUploadClient struct {
	grpc.ClientStream
}

func (x *botUploadClient) Send(m *BotUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *botUploadClient) CloseAndRecv() (*BotStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BotStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *botClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Bot_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botClient) Run(ctx context.Context, opts ...grpc.CallOption) (Bot_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bot_ServiceDesc.Streams[1], Bot_Run_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &botRunClient{stream}
	return x, nil
}

type Bot_RunClient interface {
	Send(*BotIn) error
	Recv() (*BotOut, error)
	grpc.ClientStream
}

type botRunClient struct {
	grpc.ClientStream
}

func (x *botRunClient) Send(m *BotIn) error {
	return x.ClientStream.SendMsg(m)
}

func (x *botRunClient) Recv() (*BotOut, error) {
	m := new(BotOut)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BotServer is the server API for Bot service.
// All implementations must embed UnimplementedBotServer
// for forward compatibility
type BotServer interface {
	// Upload a bot. If more than N bots exist, the oldest one will be overwritten and deleted.
	Upload(Bot_UploadServer) error
	// Get a list of currently uploaded bots. There is a limit of N.
	GetStatus(context.Context, *Empty) (*StatusResponse, error)
	// The bot must have already been uploaded.  The Bot is frozen upon disconnect.
	Run(Bot_RunServer) error
	mustEmbedUnimplementedBotServer()
}

// UnimplementedBotServer must be embedded to have forward compatible implementations.
type UnimplementedBotServer struct {
}

func (UnimplementedBotServer) Upload(Bot_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedBotServer) GetStatus(context.Context, *Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedBotServer) Run(Bot_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedBotServer) mustEmbedUnimplementedBotServer() {}

// UnsafeBotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotServer will
// result in compilation errors.
type UnsafeBotServer interface {
	mustEmbedUnimplementedBotServer()
}

func RegisterBotServer(s grpc.ServiceRegistrar, srv BotServer) {
	s.RegisterService(&Bot_ServiceDesc, srv)
}

func _Bot_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotServer).Upload(&botUploadServer{stream})
}

type Bot_UploadServer interface {
	SendAndClose(*BotStatus) error
	Recv() (*BotUploadRequest, error)
	grpc.ServerStream
}

type botUploadServer struct {
	grpc.ServerStream
}

func (x *botUploadServer) SendAndClose(m *BotStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *botUploadServer) Recv() (*BotUploadRequest, error) {
	m := new(BotUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

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

func _Bot_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotServer).Run(&botRunServer{stream})
}

type Bot_RunServer interface {
	Send(*BotOut) error
	Recv() (*BotIn, error)
	grpc.ServerStream
}

type botRunServer struct {
	grpc.ServerStream
}

func (x *botRunServer) Send(m *BotOut) error {
	return x.ServerStream.SendMsg(m)
}

func (x *botRunServer) Recv() (*BotIn, error) {
	m := new(BotIn)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Bot_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Run",
			Handler:       _Bot_Run_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bot.proto",
}
