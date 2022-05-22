// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragù               v0.2.3
// source: pkg/agent/remote.proto

package agent

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

// RemoteControlClient is the client API for RemoteControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteControlClient interface {
	GetHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AgentHealth, error)
}

type remoteControlClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteControlClient(cc grpc.ClientConnInterface) RemoteControlClient {
	return &remoteControlClient{cc}
}

func (c *remoteControlClient) GetHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AgentHealth, error) {
	out := new(AgentHealth)
	err := c.cc.Invoke(ctx, "/agent.RemoteControl/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteControlServer is the server API for RemoteControl service.
// All implementations must embed UnimplementedRemoteControlServer
// for forward compatibility
type RemoteControlServer interface {
	GetHealth(context.Context, *emptypb.Empty) (*AgentHealth, error)
	mustEmbedUnimplementedRemoteControlServer()
}

// UnimplementedRemoteControlServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteControlServer struct {
}

func (UnimplementedRemoteControlServer) GetHealth(context.Context, *emptypb.Empty) (*AgentHealth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedRemoteControlServer) mustEmbedUnimplementedRemoteControlServer() {}

// UnsafeRemoteControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteControlServer will
// result in compilation errors.
type UnsafeRemoteControlServer interface {
	mustEmbedUnimplementedRemoteControlServer()
}

func RegisterRemoteControlServer(s grpc.ServiceRegistrar, srv RemoteControlServer) {
	s.RegisterService(&RemoteControl_ServiceDesc, srv)
}

func _RemoteControl_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteControlServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.RemoteControl/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteControlServer).GetHealth(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteControl_ServiceDesc is the grpc.ServiceDesc for RemoteControl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteControl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.RemoteControl",
	HandlerType: (*RemoteControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _RemoteControl_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/agent/remote.proto",
}
