// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/rbac/rbac.go

// Package mock_rbac is a generated GoMock package.
package mock_rbac

import (
	context "context"
	reflect "reflect"

	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	gomock "go.uber.org/mock/gomock"
	rbac "github.com/rancher/opni/pkg/rbac"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// AccessHeader mocks base method.
func (m *MockProvider) AccessHeader(arg0 context.Context, arg1 *v1.ReferenceList) (rbac.RBACHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessHeader", arg0, arg1)
	ret0, _ := ret[0].(rbac.RBACHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessHeader indicates an expected call of AccessHeader.
func (mr *MockProviderMockRecorder) AccessHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessHeader", reflect.TypeOf((*MockProvider)(nil).AccessHeader), arg0, arg1)
}
