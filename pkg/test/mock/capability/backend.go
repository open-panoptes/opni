// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/apis/capability/v1/capability_grpc.pb.go

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	v1 "github.com/rancher/opni/pkg/apis/capability/v1"
	v10 "github.com/rancher/opni/pkg/apis/core/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockBackendClient is a mock of BackendClient interface.
type MockBackendClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackendClientMockRecorder
}

// MockBackendClientMockRecorder is the mock recorder for MockBackendClient.
type MockBackendClientMockRecorder struct {
	mock *MockBackendClient
}

// NewMockBackendClient creates a new mock instance.
func NewMockBackendClient(ctrl *gomock.Controller) *MockBackendClient {
	mock := &MockBackendClient{ctrl: ctrl}
	mock.recorder = &MockBackendClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendClient) EXPECT() *MockBackendClientMockRecorder {
	return m.recorder
}

// CanInstall mocks base method.
func (m *MockBackendClient) CanInstall(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CanInstall", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanInstall indicates an expected call of CanInstall.
func (mr *MockBackendClientMockRecorder) CanInstall(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanInstall", reflect.TypeOf((*MockBackendClient)(nil).CanInstall), varargs...)
}

// CancelUninstall mocks base method.
func (m *MockBackendClient) CancelUninstall(ctx context.Context, in *v10.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelUninstall", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelUninstall indicates an expected call of CancelUninstall.
func (mr *MockBackendClientMockRecorder) CancelUninstall(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelUninstall", reflect.TypeOf((*MockBackendClient)(nil).CancelUninstall), varargs...)
}

// Info mocks base method.
func (m *MockBackendClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Details, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*v1.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockBackendClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBackendClient)(nil).Info), varargs...)
}

// Install mocks base method.
func (m *MockBackendClient) Install(ctx context.Context, in *v1.InstallRequest, opts ...grpc.CallOption) (*v1.InstallResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Install", varargs...)
	ret0, _ := ret[0].(*v1.InstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockBackendClientMockRecorder) Install(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockBackendClient)(nil).Install), varargs...)
}

// InstallerTemplate mocks base method.
func (m *MockBackendClient) InstallerTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.InstallerTemplateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InstallerTemplate", varargs...)
	ret0, _ := ret[0].(*v1.InstallerTemplateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallerTemplate indicates an expected call of InstallerTemplate.
func (mr *MockBackendClientMockRecorder) InstallerTemplate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallerTemplate", reflect.TypeOf((*MockBackendClient)(nil).InstallerTemplate), varargs...)
}

// Status mocks base method.
func (m *MockBackendClient) Status(ctx context.Context, in *v10.Reference, opts ...grpc.CallOption) (*v1.NodeCapabilityStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*v1.NodeCapabilityStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockBackendClientMockRecorder) Status(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockBackendClient)(nil).Status), varargs...)
}

// Uninstall mocks base method.
func (m *MockBackendClient) Uninstall(ctx context.Context, in *v1.UninstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Uninstall", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockBackendClientMockRecorder) Uninstall(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockBackendClient)(nil).Uninstall), varargs...)
}

// UninstallStatus mocks base method.
func (m *MockBackendClient) UninstallStatus(ctx context.Context, in *v10.Reference, opts ...grpc.CallOption) (*v10.TaskStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UninstallStatus", varargs...)
	ret0, _ := ret[0].(*v10.TaskStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UninstallStatus indicates an expected call of UninstallStatus.
func (mr *MockBackendClientMockRecorder) UninstallStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallStatus", reflect.TypeOf((*MockBackendClient)(nil).UninstallStatus), varargs...)
}

// MockBackendServer is a mock of BackendServer interface.
type MockBackendServer struct {
	ctrl     *gomock.Controller
	recorder *MockBackendServerMockRecorder
}

// MockBackendServerMockRecorder is the mock recorder for MockBackendServer.
type MockBackendServerMockRecorder struct {
	mock *MockBackendServer
}

// NewMockBackendServer creates a new mock instance.
func NewMockBackendServer(ctrl *gomock.Controller) *MockBackendServer {
	mock := &MockBackendServer{ctrl: ctrl}
	mock.recorder = &MockBackendServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendServer) EXPECT() *MockBackendServerMockRecorder {
	return m.recorder
}

// CanInstall mocks base method.
func (m *MockBackendServer) CanInstall(arg0 context.Context, arg1 *emptypb.Empty) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanInstall", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanInstall indicates an expected call of CanInstall.
func (mr *MockBackendServerMockRecorder) CanInstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanInstall", reflect.TypeOf((*MockBackendServer)(nil).CanInstall), arg0, arg1)
}

// CancelUninstall mocks base method.
func (m *MockBackendServer) CancelUninstall(arg0 context.Context, arg1 *v10.Reference) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelUninstall", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelUninstall indicates an expected call of CancelUninstall.
func (mr *MockBackendServerMockRecorder) CancelUninstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelUninstall", reflect.TypeOf((*MockBackendServer)(nil).CancelUninstall), arg0, arg1)
}

// Info mocks base method.
func (m *MockBackendServer) Info(arg0 context.Context, arg1 *emptypb.Empty) (*v1.Details, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*v1.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockBackendServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBackendServer)(nil).Info), arg0, arg1)
}

// Install mocks base method.
func (m *MockBackendServer) Install(arg0 context.Context, arg1 *v1.InstallRequest) (*v1.InstallResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1)
	ret0, _ := ret[0].(*v1.InstallResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockBackendServerMockRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockBackendServer)(nil).Install), arg0, arg1)
}

// InstallerTemplate mocks base method.
func (m *MockBackendServer) InstallerTemplate(arg0 context.Context, arg1 *emptypb.Empty) (*v1.InstallerTemplateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallerTemplate", arg0, arg1)
	ret0, _ := ret[0].(*v1.InstallerTemplateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallerTemplate indicates an expected call of InstallerTemplate.
func (mr *MockBackendServerMockRecorder) InstallerTemplate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallerTemplate", reflect.TypeOf((*MockBackendServer)(nil).InstallerTemplate), arg0, arg1)
}

// Status mocks base method.
func (m *MockBackendServer) Status(arg0 context.Context, arg1 *v10.Reference) (*v1.NodeCapabilityStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*v1.NodeCapabilityStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockBackendServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockBackendServer)(nil).Status), arg0, arg1)
}

// Uninstall mocks base method.
func (m *MockBackendServer) Uninstall(arg0 context.Context, arg1 *v1.UninstallRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockBackendServerMockRecorder) Uninstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockBackendServer)(nil).Uninstall), arg0, arg1)
}

// UninstallStatus mocks base method.
func (m *MockBackendServer) UninstallStatus(arg0 context.Context, arg1 *v10.Reference) (*v10.TaskStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallStatus", arg0, arg1)
	ret0, _ := ret[0].(*v10.TaskStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UninstallStatus indicates an expected call of UninstallStatus.
func (mr *MockBackendServerMockRecorder) UninstallStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallStatus", reflect.TypeOf((*MockBackendServer)(nil).UninstallStatus), arg0, arg1)
}

// mustEmbedUnimplementedBackendServer mocks base method.
func (m *MockBackendServer) mustEmbedUnimplementedBackendServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBackendServer")
}

// mustEmbedUnimplementedBackendServer indicates an expected call of mustEmbedUnimplementedBackendServer.
func (mr *MockBackendServerMockRecorder) mustEmbedUnimplementedBackendServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBackendServer", reflect.TypeOf((*MockBackendServer)(nil).mustEmbedUnimplementedBackendServer))
}

// MockUnsafeBackendServer is a mock of UnsafeBackendServer interface.
type MockUnsafeBackendServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBackendServerMockRecorder
}

// MockUnsafeBackendServerMockRecorder is the mock recorder for MockUnsafeBackendServer.
type MockUnsafeBackendServerMockRecorder struct {
	mock *MockUnsafeBackendServer
}

// NewMockUnsafeBackendServer creates a new mock instance.
func NewMockUnsafeBackendServer(ctrl *gomock.Controller) *MockUnsafeBackendServer {
	mock := &MockUnsafeBackendServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBackendServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBackendServer) EXPECT() *MockUnsafeBackendServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBackendServer mocks base method.
func (m *MockUnsafeBackendServer) mustEmbedUnimplementedBackendServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBackendServer")
}

// mustEmbedUnimplementedBackendServer indicates an expected call of mustEmbedUnimplementedBackendServer.
func (mr *MockUnsafeBackendServerMockRecorder) mustEmbedUnimplementedBackendServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBackendServer", reflect.TypeOf((*MockUnsafeBackendServer)(nil).mustEmbedUnimplementedBackendServer))
}

// MockNodeClient is a mock of NodeClient interface.
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient.
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance.
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// SyncNow mocks base method.
func (m *MockNodeClient) SyncNow(ctx context.Context, in *v1.Filter, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncNow", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncNow indicates an expected call of SyncNow.
func (mr *MockNodeClientMockRecorder) SyncNow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncNow", reflect.TypeOf((*MockNodeClient)(nil).SyncNow), varargs...)
}

// MockNodeServer is a mock of NodeServer interface.
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServerMockRecorder
}

// MockNodeServerMockRecorder is the mock recorder for MockNodeServer.
type MockNodeServerMockRecorder struct {
	mock *MockNodeServer
}

// NewMockNodeServer creates a new mock instance.
func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &MockNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeServer) EXPECT() *MockNodeServerMockRecorder {
	return m.recorder
}

// SyncNow mocks base method.
func (m *MockNodeServer) SyncNow(arg0 context.Context, arg1 *v1.Filter) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncNow", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncNow indicates an expected call of SyncNow.
func (mr *MockNodeServerMockRecorder) SyncNow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncNow", reflect.TypeOf((*MockNodeServer)(nil).SyncNow), arg0, arg1)
}

// mustEmbedUnimplementedNodeServer mocks base method.
func (m *MockNodeServer) mustEmbedUnimplementedNodeServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedNodeServer")
}

// mustEmbedUnimplementedNodeServer indicates an expected call of mustEmbedUnimplementedNodeServer.
func (mr *MockNodeServerMockRecorder) mustEmbedUnimplementedNodeServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedNodeServer", reflect.TypeOf((*MockNodeServer)(nil).mustEmbedUnimplementedNodeServer))
}

// MockUnsafeNodeServer is a mock of UnsafeNodeServer interface.
type MockUnsafeNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeNodeServerMockRecorder
}

// MockUnsafeNodeServerMockRecorder is the mock recorder for MockUnsafeNodeServer.
type MockUnsafeNodeServerMockRecorder struct {
	mock *MockUnsafeNodeServer
}

// NewMockUnsafeNodeServer creates a new mock instance.
func NewMockUnsafeNodeServer(ctrl *gomock.Controller) *MockUnsafeNodeServer {
	mock := &MockUnsafeNodeServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeNodeServer) EXPECT() *MockUnsafeNodeServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedNodeServer mocks base method.
func (m *MockUnsafeNodeServer) mustEmbedUnimplementedNodeServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedNodeServer")
}

// mustEmbedUnimplementedNodeServer indicates an expected call of mustEmbedUnimplementedNodeServer.
func (mr *MockUnsafeNodeServerMockRecorder) mustEmbedUnimplementedNodeServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedNodeServer", reflect.TypeOf((*MockUnsafeNodeServer)(nil).mustEmbedUnimplementedNodeServer))
}

// MockRBACManagerClient is a mock of RBACManagerClient interface.
type MockRBACManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockRBACManagerClientMockRecorder
}

// MockRBACManagerClientMockRecorder is the mock recorder for MockRBACManagerClient.
type MockRBACManagerClientMockRecorder struct {
	mock *MockRBACManagerClient
}

// NewMockRBACManagerClient creates a new mock instance.
func NewMockRBACManagerClient(ctrl *gomock.Controller) *MockRBACManagerClient {
	mock := &MockRBACManagerClient{ctrl: ctrl}
	mock.recorder = &MockRBACManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRBACManagerClient) EXPECT() *MockRBACManagerClientMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockRBACManagerClient) CreateRole(ctx context.Context, in *v10.Role, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRole", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRBACManagerClientMockRecorder) CreateRole(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRBACManagerClient)(nil).CreateRole), varargs...)
}

// DeleteRole mocks base method.
func (m *MockRBACManagerClient) DeleteRole(ctx context.Context, in *v10.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRole", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRBACManagerClientMockRecorder) DeleteRole(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRBACManagerClient)(nil).DeleteRole), varargs...)
}

// GetAvailablePermissions mocks base method.
func (m *MockRBACManagerClient) GetAvailablePermissions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v10.AvailablePermissions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAvailablePermissions", varargs...)
	ret0, _ := ret[0].(*v10.AvailablePermissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailablePermissions indicates an expected call of GetAvailablePermissions.
func (mr *MockRBACManagerClientMockRecorder) GetAvailablePermissions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailablePermissions", reflect.TypeOf((*MockRBACManagerClient)(nil).GetAvailablePermissions), varargs...)
}

// GetRole mocks base method.
func (m *MockRBACManagerClient) GetRole(ctx context.Context, in *v10.Reference, opts ...grpc.CallOption) (*v10.Role, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRole", varargs...)
	ret0, _ := ret[0].(*v10.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockRBACManagerClientMockRecorder) GetRole(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockRBACManagerClient)(nil).GetRole), varargs...)
}

// Info mocks base method.
func (m *MockRBACManagerClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.Details, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*v1.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockRBACManagerClientMockRecorder) Info(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockRBACManagerClient)(nil).Info), varargs...)
}

// ListRoles mocks base method.
func (m *MockRBACManagerClient) ListRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v10.RoleList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoles", varargs...)
	ret0, _ := ret[0].(*v10.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockRBACManagerClientMockRecorder) ListRoles(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockRBACManagerClient)(nil).ListRoles), varargs...)
}

// UpdateRole mocks base method.
func (m *MockRBACManagerClient) UpdateRole(ctx context.Context, in *v10.Role, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRole", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockRBACManagerClientMockRecorder) UpdateRole(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockRBACManagerClient)(nil).UpdateRole), varargs...)
}

// MockRBACManagerServer is a mock of RBACManagerServer interface.
type MockRBACManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockRBACManagerServerMockRecorder
}

// MockRBACManagerServerMockRecorder is the mock recorder for MockRBACManagerServer.
type MockRBACManagerServerMockRecorder struct {
	mock *MockRBACManagerServer
}

// NewMockRBACManagerServer creates a new mock instance.
func NewMockRBACManagerServer(ctrl *gomock.Controller) *MockRBACManagerServer {
	mock := &MockRBACManagerServer{ctrl: ctrl}
	mock.recorder = &MockRBACManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRBACManagerServer) EXPECT() *MockRBACManagerServerMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockRBACManagerServer) CreateRole(arg0 context.Context, arg1 *v10.Role) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockRBACManagerServerMockRecorder) CreateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockRBACManagerServer)(nil).CreateRole), arg0, arg1)
}

// DeleteRole mocks base method.
func (m *MockRBACManagerServer) DeleteRole(arg0 context.Context, arg1 *v10.Reference) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockRBACManagerServerMockRecorder) DeleteRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockRBACManagerServer)(nil).DeleteRole), arg0, arg1)
}

// GetAvailablePermissions mocks base method.
func (m *MockRBACManagerServer) GetAvailablePermissions(arg0 context.Context, arg1 *emptypb.Empty) (*v10.AvailablePermissions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailablePermissions", arg0, arg1)
	ret0, _ := ret[0].(*v10.AvailablePermissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailablePermissions indicates an expected call of GetAvailablePermissions.
func (mr *MockRBACManagerServerMockRecorder) GetAvailablePermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailablePermissions", reflect.TypeOf((*MockRBACManagerServer)(nil).GetAvailablePermissions), arg0, arg1)
}

// GetRole mocks base method.
func (m *MockRBACManagerServer) GetRole(arg0 context.Context, arg1 *v10.Reference) (*v10.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0, arg1)
	ret0, _ := ret[0].(*v10.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockRBACManagerServerMockRecorder) GetRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockRBACManagerServer)(nil).GetRole), arg0, arg1)
}

// Info mocks base method.
func (m *MockRBACManagerServer) Info(arg0 context.Context, arg1 *emptypb.Empty) (*v1.Details, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0, arg1)
	ret0, _ := ret[0].(*v1.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockRBACManagerServerMockRecorder) Info(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockRBACManagerServer)(nil).Info), arg0, arg1)
}

// ListRoles mocks base method.
func (m *MockRBACManagerServer) ListRoles(arg0 context.Context, arg1 *emptypb.Empty) (*v10.RoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoles", arg0, arg1)
	ret0, _ := ret[0].(*v10.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoles indicates an expected call of ListRoles.
func (mr *MockRBACManagerServerMockRecorder) ListRoles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockRBACManagerServer)(nil).ListRoles), arg0, arg1)
}

// UpdateRole mocks base method.
func (m *MockRBACManagerServer) UpdateRole(arg0 context.Context, arg1 *v10.Role) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockRBACManagerServerMockRecorder) UpdateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockRBACManagerServer)(nil).UpdateRole), arg0, arg1)
}

// mustEmbedUnimplementedRBACManagerServer mocks base method.
func (m *MockRBACManagerServer) mustEmbedUnimplementedRBACManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedRBACManagerServer")
}

// mustEmbedUnimplementedRBACManagerServer indicates an expected call of mustEmbedUnimplementedRBACManagerServer.
func (mr *MockRBACManagerServerMockRecorder) mustEmbedUnimplementedRBACManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedRBACManagerServer", reflect.TypeOf((*MockRBACManagerServer)(nil).mustEmbedUnimplementedRBACManagerServer))
}

// MockUnsafeRBACManagerServer is a mock of UnsafeRBACManagerServer interface.
type MockUnsafeRBACManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeRBACManagerServerMockRecorder
}

// MockUnsafeRBACManagerServerMockRecorder is the mock recorder for MockUnsafeRBACManagerServer.
type MockUnsafeRBACManagerServerMockRecorder struct {
	mock *MockUnsafeRBACManagerServer
}

// NewMockUnsafeRBACManagerServer creates a new mock instance.
func NewMockUnsafeRBACManagerServer(ctrl *gomock.Controller) *MockUnsafeRBACManagerServer {
	mock := &MockUnsafeRBACManagerServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeRBACManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeRBACManagerServer) EXPECT() *MockUnsafeRBACManagerServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedRBACManagerServer mocks base method.
func (m *MockUnsafeRBACManagerServer) mustEmbedUnimplementedRBACManagerServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedRBACManagerServer")
}

// mustEmbedUnimplementedRBACManagerServer indicates an expected call of mustEmbedUnimplementedRBACManagerServer.
func (mr *MockUnsafeRBACManagerServerMockRecorder) mustEmbedUnimplementedRBACManagerServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedRBACManagerServer", reflect.TypeOf((*MockUnsafeRBACManagerServer)(nil).mustEmbedUnimplementedRBACManagerServer))
}
