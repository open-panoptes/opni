// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragù               v0.2.3
// source: pkg/gateway/stream/stream.proto

package stream

import (
	context "context"
	totem "github.com/kralicky/totem"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Stream_ConnectClient, error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Stream_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stream_ServiceDesc.Streams[0], "/stream.Stream/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamConnectClient{stream}
	return x, nil
}

type Stream_ConnectClient interface {
	Send(*totem.RPC) error
	Recv() (*totem.RPC, error)
	grpc.ClientStream
}

type streamConnectClient struct {
	grpc.ClientStream
}

func (x *streamConnectClient) Send(m *totem.RPC) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamConnectClient) Recv() (*totem.RPC, error) {
	m := new(totem.RPC)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
// All implementations must embed UnimplementedStreamServer
// for forward compatibility
type StreamServer interface {
	Connect(Stream_ConnectServer) error
	mustEmbedUnimplementedStreamServer()
}

// UnimplementedStreamServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (UnimplementedStreamServer) Connect(Stream_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedStreamServer) mustEmbedUnimplementedStreamServer() {}

// UnsafeStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServer will
// result in compilation errors.
type UnsafeStreamServer interface {
	mustEmbedUnimplementedStreamServer()
}

func RegisterStreamServer(s grpc.ServiceRegistrar, srv StreamServer) {
	s.RegisterService(&Stream_ServiceDesc, srv)
}

func _Stream_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServer).Connect(&streamConnectServer{stream})
}

type Stream_ConnectServer interface {
	Send(*totem.RPC) error
	Recv() (*totem.RPC, error)
	grpc.ServerStream
}

type streamConnectServer struct {
	grpc.ServerStream
}

func (x *streamConnectServer) Send(m *totem.RPC) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamConnectServer) Recv() (*totem.RPC, error) {
	m := new(totem.RPC)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Stream_ServiceDesc is the grpc.ServiceDesc for Stream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Stream_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/gateway/stream/stream.proto",
}
