// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-ldap/ldap/v3 (interfaces: Client)

// Package mock_v3 is a generated GoMock package.
package mock_v3

import (
	tls "crypto/tls"
	reflect "reflect"
	time "time"

	ldap "github.com/go-ldap/ldap/v3"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockClient) Add(arg0 *ldap.AddRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockClientMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockClient)(nil).Add), arg0)
}

// Bind mocks base method.
func (m *MockClient) Bind(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockClientMockRecorder) Bind(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockClient)(nil).Bind), arg0, arg1)
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Compare mocks base method.
func (m *MockClient) Compare(arg0, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compare indicates an expected call of Compare.
func (mr *MockClientMockRecorder) Compare(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockClient)(nil).Compare), arg0, arg1, arg2)
}

// Del mocks base method.
func (m *MockClient) Del(arg0 *ldap.DelRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockClientMockRecorder) Del(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockClient)(nil).Del), arg0)
}

// ExternalBind mocks base method.
func (m *MockClient) ExternalBind() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalBind")
	ret0, _ := ret[0].(error)
	return ret0
}

// ExternalBind indicates an expected call of ExternalBind.
func (mr *MockClientMockRecorder) ExternalBind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalBind", reflect.TypeOf((*MockClient)(nil).ExternalBind))
}

// IsClosing mocks base method.
func (m *MockClient) IsClosing() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosing")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosing indicates an expected call of IsClosing.
func (mr *MockClientMockRecorder) IsClosing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosing", reflect.TypeOf((*MockClient)(nil).IsClosing))
}

// Modify mocks base method.
func (m *MockClient) Modify(arg0 *ldap.ModifyRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Modify indicates an expected call of Modify.
func (mr *MockClientMockRecorder) Modify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockClient)(nil).Modify), arg0)
}

// ModifyDN mocks base method.
func (m *MockClient) ModifyDN(arg0 *ldap.ModifyDNRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyDN", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyDN indicates an expected call of ModifyDN.
func (mr *MockClientMockRecorder) ModifyDN(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyDN", reflect.TypeOf((*MockClient)(nil).ModifyDN), arg0)
}

// ModifyWithResult mocks base method.
func (m *MockClient) ModifyWithResult(arg0 *ldap.ModifyRequest) (*ldap.ModifyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyWithResult", arg0)
	ret0, _ := ret[0].(*ldap.ModifyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyWithResult indicates an expected call of ModifyWithResult.
func (mr *MockClientMockRecorder) ModifyWithResult(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyWithResult", reflect.TypeOf((*MockClient)(nil).ModifyWithResult), arg0)
}

// NTLMUnauthenticatedBind mocks base method.
func (m *MockClient) NTLMUnauthenticatedBind(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NTLMUnauthenticatedBind", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NTLMUnauthenticatedBind indicates an expected call of NTLMUnauthenticatedBind.
func (mr *MockClientMockRecorder) NTLMUnauthenticatedBind(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NTLMUnauthenticatedBind", reflect.TypeOf((*MockClient)(nil).NTLMUnauthenticatedBind), arg0, arg1)
}

// PasswordModify mocks base method.
func (m *MockClient) PasswordModify(arg0 *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordModify", arg0)
	ret0, _ := ret[0].(*ldap.PasswordModifyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PasswordModify indicates an expected call of PasswordModify.
func (mr *MockClientMockRecorder) PasswordModify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordModify", reflect.TypeOf((*MockClient)(nil).PasswordModify), arg0)
}

// Search mocks base method.
func (m *MockClient) Search(arg0 *ldap.SearchRequest) (*ldap.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].(*ldap.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockClientMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockClient)(nil).Search), arg0)
}

// SearchWithPaging mocks base method.
func (m *MockClient) SearchWithPaging(arg0 *ldap.SearchRequest, arg1 uint32) (*ldap.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithPaging", arg0, arg1)
	ret0, _ := ret[0].(*ldap.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithPaging indicates an expected call of SearchWithPaging.
func (mr *MockClientMockRecorder) SearchWithPaging(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithPaging", reflect.TypeOf((*MockClient)(nil).SearchWithPaging), arg0, arg1)
}

// SetTimeout mocks base method.
func (m *MockClient) SetTimeout(arg0 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimeout", arg0)
}

// SetTimeout indicates an expected call of SetTimeout.
func (mr *MockClientMockRecorder) SetTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockClient)(nil).SetTimeout), arg0)
}

// SimpleBind mocks base method.
func (m *MockClient) SimpleBind(arg0 *ldap.SimpleBindRequest) (*ldap.SimpleBindResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimpleBind", arg0)
	ret0, _ := ret[0].(*ldap.SimpleBindResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SimpleBind indicates an expected call of SimpleBind.
func (mr *MockClientMockRecorder) SimpleBind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimpleBind", reflect.TypeOf((*MockClient)(nil).SimpleBind), arg0)
}

// Start mocks base method.
func (m *MockClient) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClient)(nil).Start))
}

// StartTLS mocks base method.
func (m *MockClient) StartTLS(arg0 *tls.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTLS", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartTLS indicates an expected call of StartTLS.
func (mr *MockClientMockRecorder) StartTLS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTLS", reflect.TypeOf((*MockClient)(nil).StartTLS), arg0)
}

// TLSConnectionState mocks base method.
func (m *MockClient) TLSConnectionState() (tls.ConnectionState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLSConnectionState")
	ret0, _ := ret[0].(tls.ConnectionState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// TLSConnectionState indicates an expected call of TLSConnectionState.
func (mr *MockClientMockRecorder) TLSConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSConnectionState", reflect.TypeOf((*MockClient)(nil).TLSConnectionState))
}

// UnauthenticatedBind mocks base method.
func (m *MockClient) UnauthenticatedBind(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnauthenticatedBind", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnauthenticatedBind indicates an expected call of UnauthenticatedBind.
func (mr *MockClientMockRecorder) UnauthenticatedBind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnauthenticatedBind", reflect.TypeOf((*MockClient)(nil).UnauthenticatedBind), arg0)
}

// Unbind mocks base method.
func (m *MockClient) Unbind() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unbind")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unbind indicates an expected call of Unbind.
func (mr *MockClientMockRecorder) Unbind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unbind", reflect.TypeOf((*MockClient)(nil).Unbind))
}
