// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo (interfaces: CommandsRegister)

// Package mock_mongo is a generated GoMock package.
package mock_mongo

import (
	context "context"
	reflect "reflect"

	types "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockCommandsRegister is a mock of CommandsRegister interface.
type MockCommandsRegister struct {
	ctrl     *gomock.Controller
	recorder *MockCommandsRegisterMockRecorder
}

// MockCommandsRegisterMockRecorder is the mock recorder for MockCommandsRegister.
type MockCommandsRegisterMockRecorder struct {
	mock *MockCommandsRegister
}

// NewMockCommandsRegister creates a new mock instance.
func NewMockCommandsRegister(ctrl *gomock.Controller) *MockCommandsRegister {
	mock := &MockCommandsRegister{ctrl: ctrl}
	mock.recorder = &MockCommandsRegisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandsRegister) EXPECT() *MockCommandsRegisterMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *MockCommandsRegister) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockCommandsRegisterMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockCommandsRegister)(nil).Clear))
}

// Commit mocks base method.
func (m *MockCommandsRegister) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockCommandsRegisterMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockCommandsRegister)(nil).Commit), arg0)
}

// RegisterInsert mocks base method.
func (m *MockCommandsRegister) RegisterInsert(arg0 *types.Entity) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInsert", arg0)
}

// RegisterInsert indicates an expected call of RegisterInsert.
func (mr *MockCommandsRegisterMockRecorder) RegisterInsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInsert", reflect.TypeOf((*MockCommandsRegister)(nil).RegisterInsert), arg0)
}

// RegisterUpdate mocks base method.
func (m *MockCommandsRegister) RegisterUpdate(arg0 string, arg1 primitive.M) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterUpdate", arg0, arg1)
}

// RegisterUpdate indicates an expected call of RegisterUpdate.
func (mr *MockCommandsRegisterMockRecorder) RegisterUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUpdate", reflect.TypeOf((*MockCommandsRegister)(nil).RegisterUpdate), arg0, arg1)
}
