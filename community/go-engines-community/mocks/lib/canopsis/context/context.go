// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/context (interfaces: ContextGraphManager,EntityServiceStorage)

// Package mock_context is a generated GoMock package.
package mock_context

import (
	context "context"
	entityservice "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/entityservice"
	types "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContextGraphManager is a mock of ContextGraphManager interface
type MockContextGraphManager struct {
	ctrl     *gomock.Controller
	recorder *MockContextGraphManagerMockRecorder
}

// MockContextGraphManagerMockRecorder is the mock recorder for MockContextGraphManager
type MockContextGraphManagerMockRecorder struct {
	mock *MockContextGraphManager
}

// NewMockContextGraphManager creates a new mock instance
func NewMockContextGraphManager(ctrl *gomock.Controller) *MockContextGraphManager {
	mock := &MockContextGraphManager{ctrl: ctrl}
	mock.recorder = &MockContextGraphManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContextGraphManager) EXPECT() *MockContextGraphManagerMockRecorder {
	return m.recorder
}

// CheckServices mocks base method
func (m *MockContextGraphManager) CheckServices(arg0 context.Context, arg1 []types.Entity) ([]types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckServices", arg0, arg1)
	ret0, _ := ret[0].([]types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckServices indicates an expected call of CheckServices
func (mr *MockContextGraphManagerMockRecorder) CheckServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckServices", reflect.TypeOf((*MockContextGraphManager)(nil).CheckServices), arg0, arg1)
}

// FillResourcesWithInfos mocks base method
func (m *MockContextGraphManager) FillResourcesWithInfos(arg0 context.Context, arg1 types.Entity) ([]types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FillResourcesWithInfos", arg0, arg1)
	ret0, _ := ret[0].([]types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FillResourcesWithInfos indicates an expected call of FillResourcesWithInfos
func (mr *MockContextGraphManagerMockRecorder) FillResourcesWithInfos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FillResourcesWithInfos", reflect.TypeOf((*MockContextGraphManager)(nil).FillResourcesWithInfos), arg0, arg1)
}

// Get mocks base method
func (m *MockContextGraphManager) Get(arg0 context.Context, arg1 types.Event) (types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockContextGraphManagerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContextGraphManager)(nil).Get), arg0, arg1)
}

// Handle mocks base method
func (m *MockContextGraphManager) Handle(arg0 context.Context, arg1 types.Event) (types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle
func (mr *MockContextGraphManagerMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockContextGraphManager)(nil).Handle), arg0, arg1)
}

// HandleEntityServiceUpdate mocks base method
func (m *MockContextGraphManager) HandleEntityServiceUpdate(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleEntityServiceUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleEntityServiceUpdate indicates an expected call of HandleEntityServiceUpdate
func (mr *MockContextGraphManagerMockRecorder) HandleEntityServiceUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEntityServiceUpdate", reflect.TypeOf((*MockContextGraphManager)(nil).HandleEntityServiceUpdate), arg0, arg1)
}

// LoadServices mocks base method
func (m *MockContextGraphManager) LoadServices(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadServices", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadServices indicates an expected call of LoadServices
func (mr *MockContextGraphManagerMockRecorder) LoadServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadServices", reflect.TypeOf((*MockContextGraphManager)(nil).LoadServices), arg0)
}

// ReloadService mocks base method
func (m *MockContextGraphManager) ReloadService(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadService indicates an expected call of ReloadService
func (mr *MockContextGraphManagerMockRecorder) ReloadService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadService", reflect.TypeOf((*MockContextGraphManager)(nil).ReloadService), arg0, arg1)
}

// UpdateEntityInfos mocks base method
func (m *MockContextGraphManager) UpdateEntityInfos(arg0 context.Context, arg1 types.Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntityInfos", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEntityInfos indicates an expected call of UpdateEntityInfos
func (mr *MockContextGraphManagerMockRecorder) UpdateEntityInfos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntityInfos", reflect.TypeOf((*MockContextGraphManager)(nil).UpdateEntityInfos), arg0, arg1)
}

// UpdateImpactedServices mocks base method
func (m *MockContextGraphManager) UpdateImpactedServices(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImpactedServices", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImpactedServices indicates an expected call of UpdateImpactedServices
func (mr *MockContextGraphManagerMockRecorder) UpdateImpactedServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImpactedServices", reflect.TypeOf((*MockContextGraphManager)(nil).UpdateImpactedServices), arg0)
}

// MockEntityServiceStorage is a mock of EntityServiceStorage interface
type MockEntityServiceStorage struct {
	ctrl     *gomock.Controller
	recorder *MockEntityServiceStorageMockRecorder
}

// MockEntityServiceStorageMockRecorder is the mock recorder for MockEntityServiceStorage
type MockEntityServiceStorageMockRecorder struct {
	mock *MockEntityServiceStorage
}

// NewMockEntityServiceStorage creates a new mock instance
func NewMockEntityServiceStorage(ctrl *gomock.Controller) *MockEntityServiceStorage {
	mock := &MockEntityServiceStorage{ctrl: ctrl}
	mock.recorder = &MockEntityServiceStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEntityServiceStorage) EXPECT() *MockEntityServiceStorageMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockEntityServiceStorage) GetAll(arg0 context.Context) ([]entityservice.EntityService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entityservice.EntityService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockEntityServiceStorageMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEntityServiceStorage)(nil).GetAll), arg0)
}
