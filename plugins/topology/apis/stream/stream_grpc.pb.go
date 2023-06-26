// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/topology/apis/stream/stream.proto

package stream

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
	RemoteTopology_Push_FullMethodName         = "/stream.topology.RemoteTopology/Push"
	RemoteTopology_SyncTopology_FullMethodName = "/stream.topology.RemoteTopology/SyncTopology"
)

// RemoteTopologyClient is the client API for RemoteTopology service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteTopologyClient interface {
	Push(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SyncTopology(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type remoteTopologyClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteTopologyClient(cc grpc.ClientConnInterface) RemoteTopologyClient {
	return &remoteTopologyClient{cc}
}

func (c *remoteTopologyClient) Push(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RemoteTopology_Push_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteTopologyClient) SyncTopology(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RemoteTopology_SyncTopology_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteTopologyServer is the server API for RemoteTopology service.
// All implementations must embed UnimplementedRemoteTopologyServer
// for forward compatibility
type RemoteTopologyServer interface {
	Push(context.Context, *Payload) (*emptypb.Empty, error)
	SyncTopology(context.Context, *Payload) (*emptypb.Empty, error)
	mustEmbedUnimplementedRemoteTopologyServer()
}

// UnimplementedRemoteTopologyServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteTopologyServer struct {
}

func (UnimplementedRemoteTopologyServer) Push(context.Context, *Payload) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedRemoteTopologyServer) SyncTopology(context.Context, *Payload) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncTopology not implemented")
}
func (UnimplementedRemoteTopologyServer) mustEmbedUnimplementedRemoteTopologyServer() {}

// UnsafeRemoteTopologyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteTopologyServer will
// result in compilation errors.
type UnsafeRemoteTopologyServer interface {
	mustEmbedUnimplementedRemoteTopologyServer()
}

func RegisterRemoteTopologyServer(s grpc.ServiceRegistrar, srv RemoteTopologyServer) {
	s.RegisterService(&RemoteTopology_ServiceDesc, srv)
}

func _RemoteTopology_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteTopologyServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteTopology_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteTopologyServer).Push(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteTopology_SyncTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteTopologyServer).SyncTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteTopology_SyncTopology_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteTopologyServer).SyncTopology(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteTopology_ServiceDesc is the grpc.ServiceDesc for RemoteTopology service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteTopology_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.topology.RemoteTopology",
	HandlerType: (*RemoteTopologyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _RemoteTopology_Push_Handler,
		},
		{
			MethodName: "SyncTopology",
			Handler:    _RemoteTopology_SyncTopology_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/topology/apis/stream/stream.proto",
}