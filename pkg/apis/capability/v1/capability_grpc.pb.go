// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/apis/capability/v1/capability.proto

package v1

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendClient interface {
	// Returns info about the backend, including capability name
	Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InfoResponse, error)
	// Deprecated: Do not use.
	// Returns an error if installing the capability would fail.
	CanInstall(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Installs the capability on a cluster.
	Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*InstallResponse, error)
	// Returns common runtime config info for this capability from a specific
	// cluster (node).
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*NodeCapabilityStatus, error)
	// Requests the backend to clean up any resources it owns and prepare
	// for uninstallation. This process is asynchronous. The status of the
	// operation can be queried using the UninstallStatus method, or canceled
	// using the CancelUninstall method.
	Uninstall(ctx context.Context, in *UninstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Gets the status of the uninstall task for the given cluster.
	UninstallStatus(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*v1.TaskStatus, error)
	// Cancels an uninstall task for the given cluster, if it is still pending.
	CancelUninstall(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// Returns a go template string which will generate a shell command used to
	// install the capability. This will be displayed to the user in the UI.
	// See InstallerTemplateSpec above for the available template fields.
	InstallerTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InstallerTemplateResponse, error)
}

type backendClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendClient(cc grpc.ClientConnInterface) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/capability.Backend/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *backendClient) CanInstall(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/capability.Backend/CanInstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*InstallResponse, error) {
	out := new(InstallResponse)
	err := c.cc.Invoke(ctx, "/capability.Backend/Install", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*NodeCapabilityStatus, error) {
	out := new(NodeCapabilityStatus)
	err := c.cc.Invoke(ctx, "/capability.Backend/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Uninstall(ctx context.Context, in *UninstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/capability.Backend/Uninstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UninstallStatus(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*v1.TaskStatus, error) {
	out := new(v1.TaskStatus)
	err := c.cc.Invoke(ctx, "/capability.Backend/UninstallStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) CancelUninstall(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/capability.Backend/CancelUninstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *backendClient) InstallerTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InstallerTemplateResponse, error) {
	out := new(InstallerTemplateResponse)
	err := c.cc.Invoke(ctx, "/capability.Backend/InstallerTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
// All implementations must embed UnimplementedBackendServer
// for forward compatibility
type BackendServer interface {
	// Returns info about the backend, including capability name
	Info(context.Context, *emptypb.Empty) (*InfoResponse, error)
	// Deprecated: Do not use.
	// Returns an error if installing the capability would fail.
	CanInstall(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Installs the capability on a cluster.
	Install(context.Context, *InstallRequest) (*InstallResponse, error)
	// Returns common runtime config info for this capability from a specific
	// cluster (node).
	Status(context.Context, *StatusRequest) (*NodeCapabilityStatus, error)
	// Requests the backend to clean up any resources it owns and prepare
	// for uninstallation. This process is asynchronous. The status of the
	// operation can be queried using the UninstallStatus method, or canceled
	// using the CancelUninstall method.
	Uninstall(context.Context, *UninstallRequest) (*emptypb.Empty, error)
	// Gets the status of the uninstall task for the given cluster.
	UninstallStatus(context.Context, *v1.Reference) (*v1.TaskStatus, error)
	// Cancels an uninstall task for the given cluster, if it is still pending.
	CancelUninstall(context.Context, *v1.Reference) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// Returns a go template string which will generate a shell command used to
	// install the capability. This will be displayed to the user in the UI.
	// See InstallerTemplateSpec above for the available template fields.
	InstallerTemplate(context.Context, *emptypb.Empty) (*InstallerTemplateResponse, error)
	mustEmbedUnimplementedBackendServer()
}

// UnimplementedBackendServer must be embedded to have forward compatible implementations.
type UnimplementedBackendServer struct {
}

func (UnimplementedBackendServer) Info(context.Context, *emptypb.Empty) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedBackendServer) CanInstall(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanInstall not implemented")
}
func (UnimplementedBackendServer) Install(context.Context, *InstallRequest) (*InstallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedBackendServer) Status(context.Context, *StatusRequest) (*NodeCapabilityStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedBackendServer) Uninstall(context.Context, *UninstallRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uninstall not implemented")
}
func (UnimplementedBackendServer) UninstallStatus(context.Context, *v1.Reference) (*v1.TaskStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UninstallStatus not implemented")
}
func (UnimplementedBackendServer) CancelUninstall(context.Context, *v1.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUninstall not implemented")
}
func (UnimplementedBackendServer) InstallerTemplate(context.Context, *emptypb.Empty) (*InstallerTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallerTemplate not implemented")
}
func (UnimplementedBackendServer) mustEmbedUnimplementedBackendServer() {}

// UnsafeBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendServer will
// result in compilation errors.
type UnsafeBackendServer interface {
	mustEmbedUnimplementedBackendServer()
}

func RegisterBackendServer(s grpc.ServiceRegistrar, srv BackendServer) {
	s.RegisterService(&Backend_ServiceDesc, srv)
}

func _Backend_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Info(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CanInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CanInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/CanInstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CanInstall(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/Install",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Install(ctx, req.(*InstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Uninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Uninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/Uninstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Uninstall(ctx, req.(*UninstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UninstallStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UninstallStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/UninstallStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UninstallStatus(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CancelUninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CancelUninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/CancelUninstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CancelUninstall(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_InstallerTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).InstallerTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Backend/InstallerTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).InstallerTemplate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Backend_ServiceDesc is the grpc.ServiceDesc for Backend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "capability.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Backend_Info_Handler,
		},
		{
			MethodName: "CanInstall",
			Handler:    _Backend_CanInstall_Handler,
		},
		{
			MethodName: "Install",
			Handler:    _Backend_Install_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Backend_Status_Handler,
		},
		{
			MethodName: "Uninstall",
			Handler:    _Backend_Uninstall_Handler,
		},
		{
			MethodName: "UninstallStatus",
			Handler:    _Backend_UninstallStatus_Handler,
		},
		{
			MethodName: "CancelUninstall",
			Handler:    _Backend_CancelUninstall_Handler,
		},
		{
			MethodName: "InstallerTemplate",
			Handler:    _Backend_InstallerTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/capability/v1/capability.proto",
}

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	SyncNow(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) SyncNow(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/capability.Node/SyncNow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	SyncNow(context.Context, *Filter) (*emptypb.Empty, error)
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) SyncNow(context.Context, *Filter) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncNow not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_SyncNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SyncNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.Node/SyncNow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SyncNow(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "capability.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncNow",
			Handler:    _Node_SyncNow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/capability/v1/capability.proto",
}

// NodeManagerClient is the client API for NodeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeManagerClient interface {
	RequestSync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nodeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeManagerClient(cc grpc.ClientConnInterface) NodeManagerClient {
	return &nodeManagerClient{cc}
}

func (c *nodeManagerClient) RequestSync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/capability.NodeManager/RequestSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeManagerServer is the server API for NodeManager service.
// All implementations must embed UnimplementedNodeManagerServer
// for forward compatibility
type NodeManagerServer interface {
	RequestSync(context.Context, *SyncRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedNodeManagerServer()
}

// UnimplementedNodeManagerServer must be embedded to have forward compatible implementations.
type UnimplementedNodeManagerServer struct {
}

func (UnimplementedNodeManagerServer) RequestSync(context.Context, *SyncRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSync not implemented")
}
func (UnimplementedNodeManagerServer) mustEmbedUnimplementedNodeManagerServer() {}

// UnsafeNodeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeManagerServer will
// result in compilation errors.
type UnsafeNodeManagerServer interface {
	mustEmbedUnimplementedNodeManagerServer()
}

func RegisterNodeManagerServer(s grpc.ServiceRegistrar, srv NodeManagerServer) {
	s.RegisterService(&NodeManager_ServiceDesc, srv)
}

func _NodeManager_RequestSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeManagerServer).RequestSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/capability.NodeManager/RequestSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeManagerServer).RequestSync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeManager_ServiceDesc is the grpc.ServiceDesc for NodeManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "capability.NodeManager",
	HandlerType: (*NodeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestSync",
			Handler:    _NodeManager_RequestSync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/capability/v1/capability.proto",
}
