// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/amqp (interfaces: Connection,Channel,Publisher)

// Package mock_amqp is a generated GoMock package.
package mock_amqp

import (
	context "context"
	reflect "reflect"

	amqp "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/amqp"
	gomock "github.com/golang/mock/gomock"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Channel mocks base method.
func (m *MockConnection) Channel() (amqp.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Channel")
	ret0, _ := ret[0].(amqp.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Channel indicates an expected call of Channel.
func (mr *MockConnectionMockRecorder) Channel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Channel", reflect.TypeOf((*MockConnection)(nil).Channel))
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// IsClosed mocks base method.
func (m *MockConnection) IsClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockConnectionMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockConnection)(nil).IsClosed))
}

// MockChannel is a mock of Channel interface.
type MockChannel struct {
	ctrl     *gomock.Controller
	recorder *MockChannelMockRecorder
}

// MockChannelMockRecorder is the mock recorder for MockChannel.
type MockChannelMockRecorder struct {
	mock *MockChannel
}

// NewMockChannel creates a new mock instance.
func NewMockChannel(ctrl *gomock.Controller) *MockChannel {
	mock := &MockChannel{ctrl: ctrl}
	mock.recorder = &MockChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannel) EXPECT() *MockChannelMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *MockChannel) Ack(arg0 uint64, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *MockChannelMockRecorder) Ack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockChannel)(nil).Ack), arg0, arg1)
}

// Close mocks base method.
func (m *MockChannel) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockChannelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockChannel)(nil).Close))
}

// Consume mocks base method.
func (m *MockChannel) Consume(arg0, arg1 string, arg2, arg3, arg4, arg5 bool, arg6 amqp091.Table) (<-chan amqp091.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(<-chan amqp091.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockChannelMockRecorder) Consume(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockChannel)(nil).Consume), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ExchangeDeclare mocks base method.
func (m *MockChannel) ExchangeDeclare(arg0, arg1 string, arg2, arg3, arg4, arg5 bool, arg6 amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDeclare", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclare indicates an expected call of ExchangeDeclare.
func (mr *MockChannelMockRecorder) ExchangeDeclare(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclare", reflect.TypeOf((*MockChannel)(nil).ExchangeDeclare), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Nack mocks base method.
func (m *MockChannel) Nack(arg0 uint64, arg1, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nack", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Nack indicates an expected call of Nack.
func (mr *MockChannelMockRecorder) Nack(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nack", reflect.TypeOf((*MockChannel)(nil).Nack), arg0, arg1, arg2)
}

// PublishWithContext mocks base method.
func (m *MockChannel) PublishWithContext(arg0 context.Context, arg1, arg2 string, arg3, arg4 bool, arg5 amqp091.Publishing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishWithContext", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishWithContext indicates an expected call of PublishWithContext.
func (mr *MockChannelMockRecorder) PublishWithContext(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithContext", reflect.TypeOf((*MockChannel)(nil).PublishWithContext), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Qos mocks base method.
func (m *MockChannel) Qos(arg0, arg1 int, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Qos", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Qos indicates an expected call of Qos.
func (mr *MockChannelMockRecorder) Qos(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qos", reflect.TypeOf((*MockChannel)(nil).Qos), arg0, arg1, arg2)
}

// QueueBind mocks base method.
func (m *MockChannel) QueueBind(arg0, arg1, arg2 string, arg3 bool, arg4 amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueBind", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueBind indicates an expected call of QueueBind.
func (mr *MockChannelMockRecorder) QueueBind(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueBind", reflect.TypeOf((*MockChannel)(nil).QueueBind), arg0, arg1, arg2, arg3, arg4)
}

// QueueDeclare mocks base method.
func (m *MockChannel) QueueDeclare(arg0 string, arg1, arg2, arg3, arg4 bool, arg5 amqp091.Table) (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDeclare", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclare indicates an expected call of QueueDeclare.
func (mr *MockChannelMockRecorder) QueueDeclare(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclare", reflect.TypeOf((*MockChannel)(nil).QueueDeclare), arg0, arg1, arg2, arg3, arg4, arg5)
}

// QueueDelete mocks base method.
func (m *MockChannel) QueueDelete(arg0 string, arg1, arg2, arg3 bool) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDelete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDelete indicates an expected call of QueueDelete.
func (mr *MockChannelMockRecorder) QueueDelete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDelete", reflect.TypeOf((*MockChannel)(nil).QueueDelete), arg0, arg1, arg2, arg3)
}

// QueueInspect mocks base method.
func (m *MockChannel) QueueInspect(arg0 string) (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueInspect", arg0)
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueInspect indicates an expected call of QueueInspect.
func (mr *MockChannelMockRecorder) QueueInspect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueInspect", reflect.TypeOf((*MockChannel)(nil).QueueInspect), arg0)
}

// QueuePurge mocks base method.
func (m *MockChannel) QueuePurge(arg0 string, arg1 bool) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueuePurge", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueuePurge indicates an expected call of QueuePurge.
func (mr *MockChannelMockRecorder) QueuePurge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueuePurge", reflect.TypeOf((*MockChannel)(nil).QueuePurge), arg0, arg1)
}

// Reject mocks base method.
func (m *MockChannel) Reject(arg0 uint64, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject.
func (mr *MockChannelMockRecorder) Reject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockChannel)(nil).Reject), arg0, arg1)
}

// MockPublisher is a mock of Publisher interface.
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher.
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance.
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// PublishWithContext mocks base method.
func (m *MockPublisher) PublishWithContext(arg0 context.Context, arg1, arg2 string, arg3, arg4 bool, arg5 amqp091.Publishing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishWithContext", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishWithContext indicates an expected call of PublishWithContext.
func (mr *MockPublisherMockRecorder) PublishWithContext(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithContext", reflect.TypeOf((*MockPublisher)(nil).PublishWithContext), arg0, arg1, arg2, arg3, arg4, arg5)
}
