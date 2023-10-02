// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/contextgraph (interfaces: Manager,EntityServiceStorage)

// Package mock_contextgraph is a generated GoMock package.
package mock_contextgraph

import (
	context "context"
	reflect "reflect"

	entityservice "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/entityservice"
	types "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// CheckServices mocks base method.
func (m *MockManager) CheckServices(arg0 context.Context, arg1 []types.Entity) ([]types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckServices", arg0, arg1)
	ret0, _ := ret[0].([]types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckServices indicates an expected call of CheckServices.
func (mr *MockManagerMockRecorder) CheckServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckServices", reflect.TypeOf((*MockManager)(nil).CheckServices), arg0, arg1)
}

// FillResourcesWithInfos mocks base method.
func (m *MockManager) FillResourcesWithInfos(arg0 context.Context, arg1 types.Entity) ([]types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FillResourcesWithInfos", arg0, arg1)
	ret0, _ := ret[0].([]types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FillResourcesWithInfos indicates an expected call of FillResourcesWithInfos.
func (mr *MockManagerMockRecorder) FillResourcesWithInfos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FillResourcesWithInfos", reflect.TypeOf((*MockManager)(nil).FillResourcesWithInfos), arg0, arg1)
}

// HandleEvent mocks base method.
func (m *MockManager) HandleEvent(arg0 context.Context, arg1 types.Event) (types.Entity, []types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleEvent", arg0, arg1)
	ret0, _ := ret[0].(types.Entity)
	ret1, _ := ret[1].([]types.Entity)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// HandleEvent indicates an expected call of HandleEvent.
func (mr *MockManagerMockRecorder) HandleEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockManager)(nil).HandleEvent), arg0, arg1)
}

// RecomputeService mocks base method.
func (m *MockManager) RecomputeService(arg0 context.Context, arg1 string) (types.Entity, []types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecomputeService", arg0, arg1)
	ret0, _ := ret[0].(types.Entity)
	ret1, _ := ret[1].([]types.Entity)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecomputeService indicates an expected call of RecomputeService.
func (mr *MockManagerMockRecorder) RecomputeService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecomputeService", reflect.TypeOf((*MockManager)(nil).RecomputeService), arg0, arg1)
}

// UpdateEntities mocks base method.
func (m *MockManager) UpdateEntities(arg0 context.Context, arg1 string, arg2 []types.Entity, arg3 bool) (types.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntities", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(types.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEntities indicates an expected call of UpdateEntities.
func (mr *MockManagerMockRecorder) UpdateEntities(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntities", reflect.TypeOf((*MockManager)(nil).UpdateEntities), arg0, arg1, arg2, arg3)
}

// UpdateImpactedServicesFromDependencies mocks base method.
func (m *MockManager) UpdateImpactedServicesFromDependencies(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImpactedServicesFromDependencies", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImpactedServicesFromDependencies indicates an expected call of UpdateImpactedServicesFromDependencies.
func (mr *MockManagerMockRecorder) UpdateImpactedServicesFromDependencies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImpactedServicesFromDependencies", reflect.TypeOf((*MockManager)(nil).UpdateImpactedServicesFromDependencies), arg0)
}

// UpdateLastEventDate mocks base method.
func (m *MockManager) UpdateLastEventDate(arg0 context.Context, arg1, arg2 string, arg3 types.CpsTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastEventDate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastEventDate indicates an expected call of UpdateLastEventDate.
func (mr *MockManagerMockRecorder) UpdateLastEventDate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastEventDate", reflect.TypeOf((*MockManager)(nil).UpdateLastEventDate), arg0, arg1, arg2, arg3)
}

// MockEntityServiceStorage is a mock of EntityServiceStorage interface.
type MockEntityServiceStorage struct {
	ctrl     *gomock.Controller
	recorder *MockEntityServiceStorageMockRecorder
}

// MockEntityServiceStorageMockRecorder is the mock recorder for MockEntityServiceStorage.
type MockEntityServiceStorageMockRecorder struct {
	mock *MockEntityServiceStorage
}

// NewMockEntityServiceStorage creates a new mock instance.
func NewMockEntityServiceStorage(ctrl *gomock.Controller) *MockEntityServiceStorage {
	mock := &MockEntityServiceStorage{ctrl: ctrl}
	mock.recorder = &MockEntityServiceStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityServiceStorage) EXPECT() *MockEntityServiceStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockEntityServiceStorage) Get(arg0 context.Context, arg1 string) (entityservice.EntityService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(entityservice.EntityService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEntityServiceStorageMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEntityServiceStorage)(nil).Get), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockEntityServiceStorage) GetAll(arg0 context.Context) ([]entityservice.EntityService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entityservice.EntityService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEntityServiceStorageMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEntityServiceStorage)(nil).GetAll), arg0)
}
