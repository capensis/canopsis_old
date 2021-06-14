// Code generated by MockGen. DO NOT EDIT.
// Source: lib/redis/redislock.go

// Package mock_redis is a generated GoMock package.
package mock_redis

import (
	redis "git.canopsis.net/canopsis/go-engines/lib/redis"
	redislock "github.com/bsm/redislock"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockLockClient is a mock of LockClient interface
type MockLockClient struct {
	ctrl     *gomock.Controller
	recorder *MockLockClientMockRecorder
}

// MockLockClientMockRecorder is the mock recorder for MockLockClient
type MockLockClientMockRecorder struct {
	mock *MockLockClient
}

// NewMockLockClient creates a new mock instance
func NewMockLockClient(ctrl *gomock.Controller) *MockLockClient {
	mock := &MockLockClient{ctrl: ctrl}
	mock.recorder = &MockLockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLockClient) EXPECT() *MockLockClientMockRecorder {
	return m.recorder
}

// Obtain mocks base method
func (m *MockLockClient) Obtain(key string, ttl time.Duration, opt *redislock.Options) (redis.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Obtain", key, ttl, opt)
	ret0, _ := ret[0].(redis.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Obtain indicates an expected call of Obtain
func (mr *MockLockClientMockRecorder) Obtain(key, ttl, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Obtain", reflect.TypeOf((*MockLockClient)(nil).Obtain), key, ttl, opt)
}

// MockLock is a mock of Lock interface
type MockLock struct {
	ctrl     *gomock.Controller
	recorder *MockLockMockRecorder
}

// MockLockMockRecorder is the mock recorder for MockLock
type MockLockMockRecorder struct {
	mock *MockLock
}

// NewMockLock creates a new mock instance
func NewMockLock(ctrl *gomock.Controller) *MockLock {
	mock := &MockLock{ctrl: ctrl}
	mock.recorder = &MockLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLock) EXPECT() *MockLockMockRecorder {
	return m.recorder
}

// Key mocks base method
func (m *MockLock) Key() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

// Key indicates an expected call of Key
func (mr *MockLockMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockLock)(nil).Key))
}

// Token mocks base method
func (m *MockLock) Token() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(string)
	return ret0
}

// Token indicates an expected call of Token
func (mr *MockLockMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockLock)(nil).Token))
}

// TTL mocks base method
func (m *MockLock) TTL() (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL")
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TTL indicates an expected call of TTL
func (mr *MockLockMockRecorder) TTL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockLock)(nil).TTL))
}

// Refresh mocks base method
func (m *MockLock) Refresh(ttl time.Duration, opt *redislock.Options) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", ttl, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockLockMockRecorder) Refresh(ttl, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockLock)(nil).Refresh), ttl, opt)
}

// Release mocks base method
func (m *MockLock) Release() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release")
	ret0, _ := ret[0].(error)
	return ret0
}

// Release indicates an expected call of Release
func (mr *MockLockMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockLock)(nil).Release))
}
