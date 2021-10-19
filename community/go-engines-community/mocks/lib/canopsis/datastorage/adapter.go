// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/datastorage (interfaces: Adapter)

// Package mock_datastorage is a generated GoMock package.
package mock_datastorage

import (
	context "context"
	datastorage "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/datastorage"
	types "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAdapter is a mock of Adapter interface
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterMockRecorder
}

// MockAdapterMockRecorder is the mock recorder for MockAdapter
type MockAdapterMockRecorder struct {
	mock *MockAdapter
}

// NewMockAdapter creates a new mock instance
func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &MockAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdapter) EXPECT() *MockAdapterMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockAdapter) Get(arg0 context.Context) (datastorage.DataStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(datastorage.DataStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAdapterMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAdapter)(nil).Get), arg0)
}

// UpdateHistoryAlarm mocks base method
func (m *MockAdapter) UpdateHistoryAlarm(arg0 context.Context, arg1 datastorage.HistoryWithCount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHistoryAlarm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHistoryAlarm indicates an expected call of UpdateHistoryAlarm
func (mr *MockAdapterMockRecorder) UpdateHistoryAlarm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHistoryAlarm", reflect.TypeOf((*MockAdapter)(nil).UpdateHistoryAlarm), arg0, arg1)
}

// UpdateHistoryEntity mocks base method
func (m *MockAdapter) UpdateHistoryEntity(arg0 context.Context, arg1 datastorage.HistoryWithCount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHistoryEntity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHistoryEntity indicates an expected call of UpdateHistoryEntity
func (mr *MockAdapterMockRecorder) UpdateHistoryEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHistoryEntity", reflect.TypeOf((*MockAdapter)(nil).UpdateHistoryEntity), arg0, arg1)
}

// UpdateHistoryHealthCheck mocks base method
func (m *MockAdapter) UpdateHistoryHealthCheck(arg0 context.Context, arg1 types.CpsTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHistoryHealthCheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHistoryHealthCheck indicates an expected call of UpdateHistoryHealthCheck
func (mr *MockAdapterMockRecorder) UpdateHistoryHealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHistoryHealthCheck", reflect.TypeOf((*MockAdapter)(nil).UpdateHistoryHealthCheck), arg0, arg1)
}

// UpdateHistoryJunit mocks base method
func (m *MockAdapter) UpdateHistoryJunit(arg0 context.Context, arg1 types.CpsTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHistoryJunit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHistoryJunit indicates an expected call of UpdateHistoryJunit
func (mr *MockAdapterMockRecorder) UpdateHistoryJunit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHistoryJunit", reflect.TypeOf((*MockAdapter)(nil).UpdateHistoryJunit), arg0, arg1)
}

// UpdateHistoryPbehavior mocks base method
func (m *MockAdapter) UpdateHistoryPbehavior(arg0 context.Context, arg1 types.CpsTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHistoryPbehavior", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHistoryPbehavior indicates an expected call of UpdateHistoryPbehavior
func (mr *MockAdapterMockRecorder) UpdateHistoryPbehavior(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHistoryPbehavior", reflect.TypeOf((*MockAdapter)(nil).UpdateHistoryPbehavior), arg0, arg1)
}

// UpdateHistoryRemediation mocks base method
func (m *MockAdapter) UpdateHistoryRemediation(arg0 context.Context, arg1 types.CpsTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHistoryRemediation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHistoryRemediation indicates an expected call of UpdateHistoryRemediation
func (mr *MockAdapterMockRecorder) UpdateHistoryRemediation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHistoryRemediation", reflect.TypeOf((*MockAdapter)(nil).UpdateHistoryRemediation), arg0, arg1)
}
