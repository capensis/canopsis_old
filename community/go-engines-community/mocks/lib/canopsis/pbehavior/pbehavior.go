// Code generated by MockGen. DO NOT EDIT.
// Source: git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/pbehavior (interfaces: Service,ModelProvider,EventManager,Store,EntityTypeResolver,ComputedEntityTypeResolver,TypeComputer)

// Package mock_pbehavior is a generated GoMock package.
package mock_pbehavior

import (
	context "context"
	reflect "reflect"
	time "time"

	pbehavior "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/pbehavior"
	types "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	timespan "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/timespan"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Compute mocks base method.
func (m *MockService) Compute(arg0 context.Context, arg1 timespan.Span) (pbehavior.ComputedEntityTypeResolver, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compute", arg0, arg1)
	ret0, _ := ret[0].(pbehavior.ComputedEntityTypeResolver)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Compute indicates an expected call of Compute.
func (mr *MockServiceMockRecorder) Compute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compute", reflect.TypeOf((*MockService)(nil).Compute), arg0, arg1)
}

// Recompute mocks base method.
func (m *MockService) Recompute(arg0 context.Context) (pbehavior.ComputedEntityTypeResolver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recompute", arg0)
	ret0, _ := ret[0].(pbehavior.ComputedEntityTypeResolver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recompute indicates an expected call of Recompute.
func (mr *MockServiceMockRecorder) Recompute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recompute", reflect.TypeOf((*MockService)(nil).Recompute), arg0)
}

// RecomputeByIds mocks base method.
func (m *MockService) RecomputeByIds(arg0 context.Context, arg1 []string) (pbehavior.ComputedEntityTypeResolver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecomputeByIds", arg0, arg1)
	ret0, _ := ret[0].(pbehavior.ComputedEntityTypeResolver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecomputeByIds indicates an expected call of RecomputeByIds.
func (mr *MockServiceMockRecorder) RecomputeByIds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecomputeByIds", reflect.TypeOf((*MockService)(nil).RecomputeByIds), arg0, arg1)
}

// MockModelProvider is a mock of ModelProvider interface.
type MockModelProvider struct {
	ctrl     *gomock.Controller
	recorder *MockModelProviderMockRecorder
}

// MockModelProviderMockRecorder is the mock recorder for MockModelProvider.
type MockModelProviderMockRecorder struct {
	mock *MockModelProvider
}

// NewMockModelProvider creates a new mock instance.
func NewMockModelProvider(ctrl *gomock.Controller) *MockModelProvider {
	mock := &MockModelProvider{ctrl: ctrl}
	mock.recorder = &MockModelProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelProvider) EXPECT() *MockModelProviderMockRecorder {
	return m.recorder
}

// GetEnabledPbehaviors mocks base method.
func (m *MockModelProvider) GetEnabledPbehaviors(arg0 context.Context, arg1 timespan.Span) (map[string]pbehavior.PBehavior, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnabledPbehaviors", arg0, arg1)
	ret0, _ := ret[0].(map[string]pbehavior.PBehavior)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnabledPbehaviors indicates an expected call of GetEnabledPbehaviors.
func (mr *MockModelProviderMockRecorder) GetEnabledPbehaviors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnabledPbehaviors", reflect.TypeOf((*MockModelProvider)(nil).GetEnabledPbehaviors), arg0, arg1)
}

// GetEnabledPbehaviorsByIds mocks base method.
func (m *MockModelProvider) GetEnabledPbehaviorsByIds(arg0 context.Context, arg1 []string, arg2 timespan.Span) (map[string]pbehavior.PBehavior, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnabledPbehaviorsByIds", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]pbehavior.PBehavior)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnabledPbehaviorsByIds indicates an expected call of GetEnabledPbehaviorsByIds.
func (mr *MockModelProviderMockRecorder) GetEnabledPbehaviorsByIds(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnabledPbehaviorsByIds", reflect.TypeOf((*MockModelProvider)(nil).GetEnabledPbehaviorsByIds), arg0, arg1, arg2)
}

// GetExceptions mocks base method.
func (m *MockModelProvider) GetExceptions(arg0 context.Context) (map[string]pbehavior.Exception, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExceptions", arg0)
	ret0, _ := ret[0].(map[string]pbehavior.Exception)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExceptions indicates an expected call of GetExceptions.
func (mr *MockModelProviderMockRecorder) GetExceptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExceptions", reflect.TypeOf((*MockModelProvider)(nil).GetExceptions), arg0)
}

// GetReasons mocks base method.
func (m *MockModelProvider) GetReasons(arg0 context.Context) (map[string]pbehavior.Reason, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReasons", arg0)
	ret0, _ := ret[0].(map[string]pbehavior.Reason)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReasons indicates an expected call of GetReasons.
func (mr *MockModelProviderMockRecorder) GetReasons(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReasons", reflect.TypeOf((*MockModelProvider)(nil).GetReasons), arg0)
}

// GetTypes mocks base method.
func (m *MockModelProvider) GetTypes(arg0 context.Context) (map[string]pbehavior.Type, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypes", arg0)
	ret0, _ := ret[0].(map[string]pbehavior.Type)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypes indicates an expected call of GetTypes.
func (mr *MockModelProviderMockRecorder) GetTypes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypes", reflect.TypeOf((*MockModelProvider)(nil).GetTypes), arg0)
}

// MockEventManager is a mock of EventManager interface.
type MockEventManager struct {
	ctrl     *gomock.Controller
	recorder *MockEventManagerMockRecorder
}

// MockEventManagerMockRecorder is the mock recorder for MockEventManager.
type MockEventManagerMockRecorder struct {
	mock *MockEventManager
}

// NewMockEventManager creates a new mock instance.
func NewMockEventManager(ctrl *gomock.Controller) *MockEventManager {
	mock := &MockEventManager{ctrl: ctrl}
	mock.recorder = &MockEventManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventManager) EXPECT() *MockEventManagerMockRecorder {
	return m.recorder
}

// GetEvent mocks base method.
func (m *MockEventManager) GetEvent(arg0 pbehavior.ResolveResult, arg1 types.Alarm, arg2 types.Entity, arg3 time.Time) types.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(types.Event)
	return ret0
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockEventManagerMockRecorder) GetEvent(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockEventManager)(nil).GetEvent), arg0, arg1, arg2, arg3)
}

// GetEventType mocks base method.
func (m *MockEventManager) GetEventType(arg0 pbehavior.ResolveResult, arg1 types.PbehaviorInfo) (string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventType", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// GetEventType indicates an expected call of GetEventType.
func (mr *MockEventManagerMockRecorder) GetEventType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventType", reflect.TypeOf((*MockEventManager)(nil).GetEventType), arg0, arg1)
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DelComputedPbehavior mocks base method.
func (m *MockStore) DelComputedPbehavior(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelComputedPbehavior", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelComputedPbehavior indicates an expected call of DelComputedPbehavior.
func (mr *MockStoreMockRecorder) DelComputedPbehavior(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelComputedPbehavior", reflect.TypeOf((*MockStore)(nil).DelComputedPbehavior), arg0, arg1)
}

// GetComputed mocks base method.
func (m *MockStore) GetComputed(arg0 context.Context) (pbehavior.ComputeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputed", arg0)
	ret0, _ := ret[0].(pbehavior.ComputeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputed indicates an expected call of GetComputed.
func (mr *MockStoreMockRecorder) GetComputed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputed", reflect.TypeOf((*MockStore)(nil).GetComputed), arg0)
}

// GetComputedByIDs mocks base method.
func (m *MockStore) GetComputedByIDs(arg0 context.Context, arg1 []string) (pbehavior.ComputeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputedByIDs", arg0, arg1)
	ret0, _ := ret[0].(pbehavior.ComputeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputedByIDs indicates an expected call of GetComputedByIDs.
func (mr *MockStoreMockRecorder) GetComputedByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputedByIDs", reflect.TypeOf((*MockStore)(nil).GetComputedByIDs), arg0, arg1)
}

// GetSpan mocks base method.
func (m *MockStore) GetSpan(arg0 context.Context) (timespan.Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpan", arg0)
	ret0, _ := ret[0].(timespan.Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpan indicates an expected call of GetSpan.
func (mr *MockStoreMockRecorder) GetSpan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpan", reflect.TypeOf((*MockStore)(nil).GetSpan), arg0)
}

// SetComputed mocks base method.
func (m *MockStore) SetComputed(arg0 context.Context, arg1 pbehavior.ComputeResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetComputed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetComputed indicates an expected call of SetComputed.
func (mr *MockStoreMockRecorder) SetComputed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetComputed", reflect.TypeOf((*MockStore)(nil).SetComputed), arg0, arg1)
}

// SetComputedPbehavior mocks base method.
func (m *MockStore) SetComputedPbehavior(arg0 context.Context, arg1 string, arg2 pbehavior.ComputedPbehavior) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetComputedPbehavior", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetComputedPbehavior indicates an expected call of SetComputedPbehavior.
func (mr *MockStoreMockRecorder) SetComputedPbehavior(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetComputedPbehavior", reflect.TypeOf((*MockStore)(nil).SetComputedPbehavior), arg0, arg1, arg2)
}

// SetSpan mocks base method.
func (m *MockStore) SetSpan(arg0 context.Context, arg1 timespan.Span) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSpan", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSpan indicates an expected call of SetSpan.
func (mr *MockStoreMockRecorder) SetSpan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSpan", reflect.TypeOf((*MockStore)(nil).SetSpan), arg0, arg1)
}

// MockEntityTypeResolver is a mock of EntityTypeResolver interface.
type MockEntityTypeResolver struct {
	ctrl     *gomock.Controller
	recorder *MockEntityTypeResolverMockRecorder
}

// MockEntityTypeResolverMockRecorder is the mock recorder for MockEntityTypeResolver.
type MockEntityTypeResolverMockRecorder struct {
	mock *MockEntityTypeResolver
}

// NewMockEntityTypeResolver creates a new mock instance.
func NewMockEntityTypeResolver(ctrl *gomock.Controller) *MockEntityTypeResolver {
	mock := &MockEntityTypeResolver{ctrl: ctrl}
	mock.recorder = &MockEntityTypeResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityTypeResolver) EXPECT() *MockEntityTypeResolverMockRecorder {
	return m.recorder
}

// GetPbehaviors mocks base method.
func (m *MockEntityTypeResolver) GetPbehaviors(arg0 context.Context, arg1 []string, arg2 time.Time) (map[string]pbehavior.ResolveResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPbehaviors", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]pbehavior.ResolveResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPbehaviors indicates an expected call of GetPbehaviors.
func (mr *MockEntityTypeResolverMockRecorder) GetPbehaviors(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPbehaviors", reflect.TypeOf((*MockEntityTypeResolver)(nil).GetPbehaviors), arg0, arg1, arg2)
}

// Resolve mocks base method.
func (m *MockEntityTypeResolver) Resolve(arg0 context.Context, arg1 types.Entity, arg2 time.Time) (pbehavior.ResolveResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1, arg2)
	ret0, _ := ret[0].(pbehavior.ResolveResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve.
func (mr *MockEntityTypeResolverMockRecorder) Resolve(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockEntityTypeResolver)(nil).Resolve), arg0, arg1, arg2)
}

// MockComputedEntityTypeResolver is a mock of ComputedEntityTypeResolver interface.
type MockComputedEntityTypeResolver struct {
	ctrl     *gomock.Controller
	recorder *MockComputedEntityTypeResolverMockRecorder
}

// MockComputedEntityTypeResolverMockRecorder is the mock recorder for MockComputedEntityTypeResolver.
type MockComputedEntityTypeResolverMockRecorder struct {
	mock *MockComputedEntityTypeResolver
}

// NewMockComputedEntityTypeResolver creates a new mock instance.
func NewMockComputedEntityTypeResolver(ctrl *gomock.Controller) *MockComputedEntityTypeResolver {
	mock := &MockComputedEntityTypeResolver{ctrl: ctrl}
	mock.recorder = &MockComputedEntityTypeResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputedEntityTypeResolver) EXPECT() *MockComputedEntityTypeResolverMockRecorder {
	return m.recorder
}

// GetComputedEntityIDs mocks base method.
func (m *MockComputedEntityTypeResolver) GetComputedEntityIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputedEntityIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputedEntityIDs indicates an expected call of GetComputedEntityIDs.
func (mr *MockComputedEntityTypeResolverMockRecorder) GetComputedEntityIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputedEntityIDs", reflect.TypeOf((*MockComputedEntityTypeResolver)(nil).GetComputedEntityIDs))
}

// GetPbehaviorsCount mocks base method.
func (m *MockComputedEntityTypeResolver) GetPbehaviorsCount(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPbehaviorsCount", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPbehaviorsCount indicates an expected call of GetPbehaviorsCount.
func (mr *MockComputedEntityTypeResolverMockRecorder) GetPbehaviorsCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPbehaviorsCount", reflect.TypeOf((*MockComputedEntityTypeResolver)(nil).GetPbehaviorsCount), arg0, arg1)
}

// Resolve mocks base method.
func (m *MockComputedEntityTypeResolver) Resolve(arg0 context.Context, arg1 types.Entity, arg2 time.Time) (pbehavior.ResolveResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1, arg2)
	ret0, _ := ret[0].(pbehavior.ResolveResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve.
func (mr *MockComputedEntityTypeResolverMockRecorder) Resolve(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockComputedEntityTypeResolver)(nil).Resolve), arg0, arg1, arg2)
}

// MockTypeComputer is a mock of TypeComputer interface.
type MockTypeComputer struct {
	ctrl     *gomock.Controller
	recorder *MockTypeComputerMockRecorder
}

// MockTypeComputerMockRecorder is the mock recorder for MockTypeComputer.
type MockTypeComputerMockRecorder struct {
	mock *MockTypeComputer
}

// NewMockTypeComputer creates a new mock instance.
func NewMockTypeComputer(ctrl *gomock.Controller) *MockTypeComputer {
	mock := &MockTypeComputer{ctrl: ctrl}
	mock.recorder = &MockTypeComputerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTypeComputer) EXPECT() *MockTypeComputerMockRecorder {
	return m.recorder
}

// Compute mocks base method.
func (m *MockTypeComputer) Compute(arg0 context.Context, arg1 timespan.Span) (pbehavior.ComputeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compute", arg0, arg1)
	ret0, _ := ret[0].(pbehavior.ComputeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compute indicates an expected call of Compute.
func (mr *MockTypeComputerMockRecorder) Compute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compute", reflect.TypeOf((*MockTypeComputer)(nil).Compute), arg0, arg1)
}

// ComputeByIds mocks base method.
func (m *MockTypeComputer) ComputeByIds(arg0 context.Context, arg1 timespan.Span, arg2 []string) (pbehavior.ComputeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeByIds", arg0, arg1, arg2)
	ret0, _ := ret[0].(pbehavior.ComputeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeByIds indicates an expected call of ComputeByIds.
func (mr *MockTypeComputerMockRecorder) ComputeByIds(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeByIds", reflect.TypeOf((*MockTypeComputer)(nil).ComputeByIds), arg0, arg1, arg2)
}
