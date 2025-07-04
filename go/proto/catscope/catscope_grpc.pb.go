// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: catscope.proto

package catscope

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
	Graph_Get_FullMethodName           = "/catscopestate.Graph/Get"
	Graph_Subscribe_FullMethodName     = "/catscopestate.Graph/Subscribe"
	Graph_Chain_FullMethodName         = "/catscopestate.Graph/Chain"
	Graph_RentExemption_FullMethodName = "/catscopestate.Graph/RentExemption"
)

// GraphClient is the client API for Graph service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraphClient interface {
	// Do graph subset subscritptions. Cannot have system program account or token account as root.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Graph_SubscribeClient, error)
	// Get blockhash and slot+status updates.
	Chain(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Graph_ChainClient, error)
	RentExemption(ctx context.Context, in *RentRequest, opts ...grpc.CallOption) (*RentResponse, error)
}

type graphClient struct {
	cc grpc.ClientConnInterface
}

func NewGraphClient(cc grpc.ClientConnInterface) GraphClient {
	return &graphClient{cc}
}

func (c *graphClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Graph_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Graph_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Graph_ServiceDesc.Streams[0], Graph_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &graphSubscribeClient{stream}
	return x, nil
}

type Graph_SubscribeClient interface {
	Send(*SubscriptionRequest) error
	Recv() (*SubscriptionResponse, error)
	grpc.ClientStream
}

type graphSubscribeClient struct {
	grpc.ClientStream
}

func (x *graphSubscribeClient) Send(m *SubscriptionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *graphSubscribeClient) Recv() (*SubscriptionResponse, error) {
	m := new(SubscriptionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *graphClient) Chain(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Graph_ChainClient, error) {
	stream, err := c.cc.NewStream(ctx, &Graph_ServiceDesc.Streams[1], Graph_Chain_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &graphChainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Graph_ChainClient interface {
	Recv() (*ChainUpdate, error)
	grpc.ClientStream
}

type graphChainClient struct {
	grpc.ClientStream
}

func (x *graphChainClient) Recv() (*ChainUpdate, error) {
	m := new(ChainUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *graphClient) RentExemption(ctx context.Context, in *RentRequest, opts ...grpc.CallOption) (*RentResponse, error) {
	out := new(RentResponse)
	err := c.cc.Invoke(ctx, Graph_RentExemption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraphServer is the server API for Graph service.
// All implementations must embed UnimplementedGraphServer
// for forward compatibility
type GraphServer interface {
	// Do graph subset subscritptions. Cannot have system program account or token account as root.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Subscribe(Graph_SubscribeServer) error
	// Get blockhash and slot+status updates.
	Chain(*Empty, Graph_ChainServer) error
	RentExemption(context.Context, *RentRequest) (*RentResponse, error)
	mustEmbedUnimplementedGraphServer()
}

// UnimplementedGraphServer must be embedded to have forward compatible implementations.
type UnimplementedGraphServer struct {
}

func (UnimplementedGraphServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGraphServer) Subscribe(Graph_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedGraphServer) Chain(*Empty, Graph_ChainServer) error {
	return status.Errorf(codes.Unimplemented, "method Chain not implemented")
}
func (UnimplementedGraphServer) RentExemption(context.Context, *RentRequest) (*RentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RentExemption not implemented")
}
func (UnimplementedGraphServer) mustEmbedUnimplementedGraphServer() {}

// UnsafeGraphServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraphServer will
// result in compilation errors.
type UnsafeGraphServer interface {
	mustEmbedUnimplementedGraphServer()
}

func RegisterGraphServer(s grpc.ServiceRegistrar, srv GraphServer) {
	s.RegisterService(&Graph_ServiceDesc, srv)
}

func _Graph_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Graph_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Graph_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GraphServer).Subscribe(&graphSubscribeServer{stream})
}

type Graph_SubscribeServer interface {
	Send(*SubscriptionResponse) error
	Recv() (*SubscriptionRequest, error)
	grpc.ServerStream
}

type graphSubscribeServer struct {
	grpc.ServerStream
}

func (x *graphSubscribeServer) Send(m *SubscriptionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *graphSubscribeServer) Recv() (*SubscriptionRequest, error) {
	m := new(SubscriptionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Graph_Chain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GraphServer).Chain(m, &graphChainServer{stream})
}

type Graph_ChainServer interface {
	Send(*ChainUpdate) error
	grpc.ServerStream
}

type graphChainServer struct {
	grpc.ServerStream
}

func (x *graphChainServer) Send(m *ChainUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _Graph_RentExemption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphServer).RentExemption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Graph_RentExemption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphServer).RentExemption(ctx, req.(*RentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Graph_ServiceDesc is the grpc.ServiceDesc for Graph service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Graph_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catscopestate.Graph",
	HandlerType: (*GraphServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Graph_Get_Handler,
		},
		{
			MethodName: "RentExemption",
			Handler:    _Graph_RentExemption_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Graph_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Chain",
			Handler:       _Graph_Chain_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "catscope.proto",
}
