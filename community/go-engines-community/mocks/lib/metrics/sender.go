// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/metrics (interfaces: Sender)

// Package mock_metrics is a generated GoMock package.
package mock_metrics

import (
	context "context"
	reflect "reflect"
	time "time"

	types "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	gomock "github.com/golang/mock/gomock"
)

// MockSender is a mock of Sender interface.
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender.
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance.
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// SendAck mocks base method.
func (m *MockSender) SendAck(arg0 context.Context, arg1 types.Alarm, arg2 string, arg3 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAck", arg0, arg1, arg2, arg3)
}

// SendAck indicates an expected call of SendAck.
func (mr *MockSenderMockRecorder) SendAck(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAck", reflect.TypeOf((*MockSender)(nil).SendAck), arg0, arg1, arg2, arg3)
}

// SendAutoInstructionStart mocks base method.
func (m *MockSender) SendAutoInstructionStart(arg0 context.Context, arg1 types.Alarm, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAutoInstructionStart", arg0, arg1, arg2)
}

// SendAutoInstructionStart indicates an expected call of SendAutoInstructionStart.
func (mr *MockSenderMockRecorder) SendAutoInstructionStart(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAutoInstructionStart", reflect.TypeOf((*MockSender)(nil).SendAutoInstructionStart), arg0, arg1, arg2)
}

// SendCancelAck mocks base method.
func (m *MockSender) SendCancelAck(arg0 context.Context, arg1 types.Alarm, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendCancelAck", arg0, arg1, arg2)
}

// SendCancelAck indicates an expected call of SendCancelAck.
func (mr *MockSenderMockRecorder) SendCancelAck(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCancelAck", reflect.TypeOf((*MockSender)(nil).SendCancelAck), arg0, arg1, arg2)
}

// SendCorrelation mocks base method.
func (m *MockSender) SendCorrelation(arg0 context.Context, arg1 time.Time, arg2 types.Alarm) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendCorrelation", arg0, arg1, arg2)
}

// SendCorrelation indicates an expected call of SendCorrelation.
func (mr *MockSenderMockRecorder) SendCorrelation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCorrelation", reflect.TypeOf((*MockSender)(nil).SendCorrelation), arg0, arg1, arg2)
}

// SendCreate mocks base method.
func (m *MockSender) SendCreate(arg0 context.Context, arg1 types.Alarm, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendCreate", arg0, arg1, arg2)
}

// SendCreate indicates an expected call of SendCreate.
func (mr *MockSenderMockRecorder) SendCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCreate", reflect.TypeOf((*MockSender)(nil).SendCreate), arg0, arg1, arg2)
}

// SendCreateAndPbhEnter mocks base method.
func (m *MockSender) SendCreateAndPbhEnter(arg0 context.Context, arg1 types.Alarm, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendCreateAndPbhEnter", arg0, arg1, arg2)
}

// SendCreateAndPbhEnter indicates an expected call of SendCreateAndPbhEnter.
func (mr *MockSenderMockRecorder) SendCreateAndPbhEnter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCreateAndPbhEnter", reflect.TypeOf((*MockSender)(nil).SendCreateAndPbhEnter), arg0, arg1, arg2)
}

// SendEventMetrics mocks base method.
func (m *MockSender) SendEventMetrics(arg0 context.Context, arg1 types.Alarm, arg2 types.Entity, arg3 types.AlarmChange, arg4 time.Time, arg5, arg6 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendEventMetrics", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// SendEventMetrics indicates an expected call of SendEventMetrics.
func (mr *MockSenderMockRecorder) SendEventMetrics(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEventMetrics", reflect.TypeOf((*MockSender)(nil).SendEventMetrics), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// SendPbhEnter mocks base method.
func (m *MockSender) SendPbhEnter(arg0 context.Context, arg1 types.Alarm, arg2 types.Entity) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPbhEnter", arg0, arg1, arg2)
}

// SendPbhEnter indicates an expected call of SendPbhEnter.
func (mr *MockSenderMockRecorder) SendPbhEnter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPbhEnter", reflect.TypeOf((*MockSender)(nil).SendPbhEnter), arg0, arg1, arg2)
}

// SendPbhLeave mocks base method.
func (m *MockSender) SendPbhLeave(arg0 context.Context, arg1 types.Entity, arg2 time.Time, arg3 string, arg4 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPbhLeave", arg0, arg1, arg2, arg3, arg4)
}

// SendPbhLeave indicates an expected call of SendPbhLeave.
func (mr *MockSenderMockRecorder) SendPbhLeave(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPbhLeave", reflect.TypeOf((*MockSender)(nil).SendPbhLeave), arg0, arg1, arg2, arg3, arg4)
}

// SendPbhLeaveAndEnter mocks base method.
func (m *MockSender) SendPbhLeaveAndEnter(arg0 context.Context, arg1 types.Alarm, arg2 types.Entity, arg3 string, arg4 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPbhLeaveAndEnter", arg0, arg1, arg2, arg3, arg4)
}

// SendPbhLeaveAndEnter indicates an expected call of SendPbhLeaveAndEnter.
func (mr *MockSenderMockRecorder) SendPbhLeaveAndEnter(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPbhLeaveAndEnter", reflect.TypeOf((*MockSender)(nil).SendPbhLeaveAndEnter), arg0, arg1, arg2, arg3, arg4)
}

// SendResolve mocks base method.
func (m *MockSender) SendResolve(arg0 context.Context, arg1 types.Alarm, arg2 types.Entity, arg3 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendResolve", arg0, arg1, arg2, arg3)
}

// SendResolve indicates an expected call of SendResolve.
func (mr *MockSenderMockRecorder) SendResolve(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResolve", reflect.TypeOf((*MockSender)(nil).SendResolve), arg0, arg1, arg2, arg3)
}

// SendTicket mocks base method.
func (m *MockSender) SendTicket(arg0 context.Context, arg1 types.Alarm, arg2 string, arg3 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendTicket", arg0, arg1, arg2, arg3)
}

// SendTicket indicates an expected call of SendTicket.
func (mr *MockSenderMockRecorder) SendTicket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTicket", reflect.TypeOf((*MockSender)(nil).SendTicket), arg0, arg1, arg2, arg3)
}

// SendUpdateState mocks base method.
func (m *MockSender) SendUpdateState(arg0 context.Context, arg1 types.Alarm, arg2 types.Entity, arg3 types.CpsNumber) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendUpdateState", arg0, arg1, arg2, arg3)
}

// SendUpdateState indicates an expected call of SendUpdateState.
func (mr *MockSenderMockRecorder) SendUpdateState(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUpdateState", reflect.TypeOf((*MockSender)(nil).SendUpdateState), arg0, arg1, arg2, arg3)
}

// SendUserActivity mocks base method.
func (m *MockSender) SendUserActivity(arg0 context.Context, arg1 time.Time, arg2 string, arg3 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendUserActivity", arg0, arg1, arg2, arg3)
}

// SendUserActivity indicates an expected call of SendUserActivity.
func (mr *MockSenderMockRecorder) SendUserActivity(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUserActivity", reflect.TypeOf((*MockSender)(nil).SendUserActivity), arg0, arg1, arg2, arg3)
}
