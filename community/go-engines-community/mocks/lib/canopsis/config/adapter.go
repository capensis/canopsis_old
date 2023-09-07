// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config (interfaces: MaintenanceAdapter)

// Package mock_config is a generated GoMock package.
package mock_config

import (
	context "context"
	reflect "reflect"

	config "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	gomock "github.com/golang/mock/gomock"
)

// MockMaintenanceAdapter is a mock of MaintenanceAdapter interface.
type MockMaintenanceAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceAdapterMockRecorder
}

// MockMaintenanceAdapterMockRecorder is the mock recorder for MockMaintenanceAdapter.
type MockMaintenanceAdapterMockRecorder struct {
	mock *MockMaintenanceAdapter
}

// NewMockMaintenanceAdapter creates a new mock instance.
func NewMockMaintenanceAdapter(ctrl *gomock.Controller) *MockMaintenanceAdapter {
	mock := &MockMaintenanceAdapter{ctrl: ctrl}
	mock.recorder = &MockMaintenanceAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceAdapter) EXPECT() *MockMaintenanceAdapterMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockMaintenanceAdapter) GetConfig(arg0 context.Context) (config.MaintenanceConf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0)
	ret0, _ := ret[0].(config.MaintenanceConf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockMaintenanceAdapterMockRecorder) GetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockMaintenanceAdapter)(nil).GetConfig), arg0)
}
