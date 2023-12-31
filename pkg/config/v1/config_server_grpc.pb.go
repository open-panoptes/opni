// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/config/v1/config_server.proto

package configv1

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	driverutil "github.com/rancher/opni/pkg/plugins/driverutil"
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
	GatewayConfig_GetDefaultConfiguration_FullMethodName   = "/config.v1.GatewayConfig/GetDefaultConfiguration"
	GatewayConfig_SetDefaultConfiguration_FullMethodName   = "/config.v1.GatewayConfig/SetDefaultConfiguration"
	GatewayConfig_GetConfiguration_FullMethodName          = "/config.v1.GatewayConfig/GetConfiguration"
	GatewayConfig_SetConfiguration_FullMethodName          = "/config.v1.GatewayConfig/SetConfiguration"
	GatewayConfig_ResetDefaultConfiguration_FullMethodName = "/config.v1.GatewayConfig/ResetDefaultConfiguration"
	GatewayConfig_ResetConfiguration_FullMethodName        = "/config.v1.GatewayConfig/ResetConfiguration"
	GatewayConfig_DryRun_FullMethodName                    = "/config.v1.GatewayConfig/DryRun"
	GatewayConfig_ConfigurationHistory_FullMethodName      = "/config.v1.GatewayConfig/ConfigurationHistory"
	GatewayConfig_WatchReactive_FullMethodName             = "/config.v1.GatewayConfig/WatchReactive"
)

// GatewayConfigClient is the client API for GatewayConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayConfigClient interface {
	GetDefaultConfiguration(ctx context.Context, in *driverutil.GetRequest, opts ...grpc.CallOption) (*GatewayConfigSpec, error)
	SetDefaultConfiguration(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetConfiguration(ctx context.Context, in *driverutil.GetRequest, opts ...grpc.CallOption) (*GatewayConfigSpec, error)
	SetConfiguration(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResetDefaultConfiguration(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ResetConfiguration(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DryRun(ctx context.Context, in *DryRunRequest, opts ...grpc.CallOption) (*DryRunResponse, error)
	ConfigurationHistory(ctx context.Context, in *driverutil.ConfigurationHistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error)
	WatchReactive(ctx context.Context, in *v1.ReactiveWatchRequest, opts ...grpc.CallOption) (GatewayConfig_WatchReactiveClient, error)
}

type gatewayConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayConfigClient(cc grpc.ClientConnInterface) GatewayConfigClient {
	return &gatewayConfigClient{cc}
}

func (c *gatewayConfigClient) GetDefaultConfiguration(ctx context.Context, in *driverutil.GetRequest, opts ...grpc.CallOption) (*GatewayConfigSpec, error) {
	out := new(GatewayConfigSpec)
	err := c.cc.Invoke(ctx, GatewayConfig_GetDefaultConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) SetDefaultConfiguration(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayConfig_SetDefaultConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) GetConfiguration(ctx context.Context, in *driverutil.GetRequest, opts ...grpc.CallOption) (*GatewayConfigSpec, error) {
	out := new(GatewayConfigSpec)
	err := c.cc.Invoke(ctx, GatewayConfig_GetConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) SetConfiguration(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayConfig_SetConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) ResetDefaultConfiguration(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayConfig_ResetDefaultConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) ResetConfiguration(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GatewayConfig_ResetConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) DryRun(ctx context.Context, in *DryRunRequest, opts ...grpc.CallOption) (*DryRunResponse, error) {
	out := new(DryRunResponse)
	err := c.cc.Invoke(ctx, GatewayConfig_DryRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) ConfigurationHistory(ctx context.Context, in *driverutil.ConfigurationHistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error) {
	out := new(HistoryResponse)
	err := c.cc.Invoke(ctx, GatewayConfig_ConfigurationHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayConfigClient) WatchReactive(ctx context.Context, in *v1.ReactiveWatchRequest, opts ...grpc.CallOption) (GatewayConfig_WatchReactiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &GatewayConfig_ServiceDesc.Streams[0], GatewayConfig_WatchReactive_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayConfigWatchReactiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GatewayConfig_WatchReactiveClient interface {
	Recv() (*v1.ReactiveEvents, error)
	grpc.ClientStream
}

type gatewayConfigWatchReactiveClient struct {
	grpc.ClientStream
}

func (x *gatewayConfigWatchReactiveClient) Recv() (*v1.ReactiveEvents, error) {
	m := new(v1.ReactiveEvents)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayConfigServer is the server API for GatewayConfig service.
// All implementations should embed UnimplementedGatewayConfigServer
// for forward compatibility
type GatewayConfigServer interface {
	GetDefaultConfiguration(context.Context, *driverutil.GetRequest) (*GatewayConfigSpec, error)
	SetDefaultConfiguration(context.Context, *SetRequest) (*emptypb.Empty, error)
	GetConfiguration(context.Context, *driverutil.GetRequest) (*GatewayConfigSpec, error)
	SetConfiguration(context.Context, *SetRequest) (*emptypb.Empty, error)
	ResetDefaultConfiguration(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	ResetConfiguration(context.Context, *ResetRequest) (*emptypb.Empty, error)
	DryRun(context.Context, *DryRunRequest) (*DryRunResponse, error)
	ConfigurationHistory(context.Context, *driverutil.ConfigurationHistoryRequest) (*HistoryResponse, error)
	WatchReactive(*v1.ReactiveWatchRequest, GatewayConfig_WatchReactiveServer) error
}

// UnimplementedGatewayConfigServer should be embedded to have forward compatible implementations.
type UnimplementedGatewayConfigServer struct {
}

func (UnimplementedGatewayConfigServer) GetDefaultConfiguration(context.Context, *driverutil.GetRequest) (*GatewayConfigSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultConfiguration not implemented")
}
func (UnimplementedGatewayConfigServer) SetDefaultConfiguration(context.Context, *SetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultConfiguration not implemented")
}
func (UnimplementedGatewayConfigServer) GetConfiguration(context.Context, *driverutil.GetRequest) (*GatewayConfigSpec, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedGatewayConfigServer) SetConfiguration(context.Context, *SetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfiguration not implemented")
}
func (UnimplementedGatewayConfigServer) ResetDefaultConfiguration(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetDefaultConfiguration not implemented")
}
func (UnimplementedGatewayConfigServer) ResetConfiguration(context.Context, *ResetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetConfiguration not implemented")
}
func (UnimplementedGatewayConfigServer) DryRun(context.Context, *DryRunRequest) (*DryRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DryRun not implemented")
}
func (UnimplementedGatewayConfigServer) ConfigurationHistory(context.Context, *driverutil.ConfigurationHistoryRequest) (*HistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigurationHistory not implemented")
}
func (UnimplementedGatewayConfigServer) WatchReactive(*v1.ReactiveWatchRequest, GatewayConfig_WatchReactiveServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchReactive not implemented")
}

// UnsafeGatewayConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayConfigServer will
// result in compilation errors.
type UnsafeGatewayConfigServer interface {
	mustEmbedUnimplementedGatewayConfigServer()
}

func RegisterGatewayConfigServer(s grpc.ServiceRegistrar, srv GatewayConfigServer) {
	s.RegisterService(&GatewayConfig_ServiceDesc, srv)
}

func _GatewayConfig_GetDefaultConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(driverutil.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).GetDefaultConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_GetDefaultConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).GetDefaultConfiguration(ctx, req.(*driverutil.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_SetDefaultConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).SetDefaultConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_SetDefaultConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).SetDefaultConfiguration(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(driverutil.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_GetConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).GetConfiguration(ctx, req.(*driverutil.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_SetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).SetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_SetConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).SetConfiguration(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_ResetDefaultConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).ResetDefaultConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_ResetDefaultConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).ResetDefaultConfiguration(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_ResetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).ResetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_ResetConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).ResetConfiguration(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_DryRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DryRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).DryRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_DryRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).DryRun(ctx, req.(*DryRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_ConfigurationHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(driverutil.ConfigurationHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayConfigServer).ConfigurationHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayConfig_ConfigurationHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayConfigServer).ConfigurationHistory(ctx, req.(*driverutil.ConfigurationHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayConfig_WatchReactive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(v1.ReactiveWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayConfigServer).WatchReactive(m, &gatewayConfigWatchReactiveServer{stream})
}

type GatewayConfig_WatchReactiveServer interface {
	Send(*v1.ReactiveEvents) error
	grpc.ServerStream
}

type gatewayConfigWatchReactiveServer struct {
	grpc.ServerStream
}

func (x *gatewayConfigWatchReactiveServer) Send(m *v1.ReactiveEvents) error {
	return x.ServerStream.SendMsg(m)
}

// GatewayConfig_ServiceDesc is the grpc.ServiceDesc for GatewayConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.v1.GatewayConfig",
	HandlerType: (*GatewayConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDefaultConfiguration",
			Handler:    _GatewayConfig_GetDefaultConfiguration_Handler,
		},
		{
			MethodName: "SetDefaultConfiguration",
			Handler:    _GatewayConfig_SetDefaultConfiguration_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _GatewayConfig_GetConfiguration_Handler,
		},
		{
			MethodName: "SetConfiguration",
			Handler:    _GatewayConfig_SetConfiguration_Handler,
		},
		{
			MethodName: "ResetDefaultConfiguration",
			Handler:    _GatewayConfig_ResetDefaultConfiguration_Handler,
		},
		{
			MethodName: "ResetConfiguration",
			Handler:    _GatewayConfig_ResetConfiguration_Handler,
		},
		{
			MethodName: "DryRun",
			Handler:    _GatewayConfig_DryRun_Handler,
		},
		{
			MethodName: "ConfigurationHistory",
			Handler:    _GatewayConfig_ConfigurationHistory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchReactive",
			Handler:       _GatewayConfig_WatchReactive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/rancher/opni/pkg/config/v1/config_server.proto",
}
