// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_handlers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	controller "github.com/solo-io/external-apis/pkg/api/k8s/certificates.k8s.io/v1beta1/controller"
	v1beta1 "k8s.io/api/certificates/v1beta1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockCertificateSigningRequestEventHandler is a mock of CertificateSigningRequestEventHandler interface
type MockCertificateSigningRequestEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestEventHandlerMockRecorder
}

// MockCertificateSigningRequestEventHandlerMockRecorder is the mock recorder for MockCertificateSigningRequestEventHandler
type MockCertificateSigningRequestEventHandlerMockRecorder struct {
	mock *MockCertificateSigningRequestEventHandler
}

// NewMockCertificateSigningRequestEventHandler creates a new mock instance
func NewMockCertificateSigningRequestEventHandler(ctrl *gomock.Controller) *MockCertificateSigningRequestEventHandler {
	mock := &MockCertificateSigningRequestEventHandler{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateSigningRequestEventHandler) EXPECT() *MockCertificateSigningRequestEventHandlerMockRecorder {
	return m.recorder
}

// CreateCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestEventHandler) CreateCertificateSigningRequest(obj *v1beta1.CertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCertificateSigningRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCertificateSigningRequest indicates an expected call of CreateCertificateSigningRequest
func (mr *MockCertificateSigningRequestEventHandlerMockRecorder) CreateCertificateSigningRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestEventHandler)(nil).CreateCertificateSigningRequest), obj)
}

// UpdateCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestEventHandler) UpdateCertificateSigningRequest(old, new *v1beta1.CertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCertificateSigningRequest", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCertificateSigningRequest indicates an expected call of UpdateCertificateSigningRequest
func (mr *MockCertificateSigningRequestEventHandlerMockRecorder) UpdateCertificateSigningRequest(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestEventHandler)(nil).UpdateCertificateSigningRequest), old, new)
}

// DeleteCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestEventHandler) DeleteCertificateSigningRequest(obj *v1beta1.CertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCertificateSigningRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCertificateSigningRequest indicates an expected call of DeleteCertificateSigningRequest
func (mr *MockCertificateSigningRequestEventHandlerMockRecorder) DeleteCertificateSigningRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestEventHandler)(nil).DeleteCertificateSigningRequest), obj)
}

// GenericCertificateSigningRequest mocks base method
func (m *MockCertificateSigningRequestEventHandler) GenericCertificateSigningRequest(obj *v1beta1.CertificateSigningRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericCertificateSigningRequest", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericCertificateSigningRequest indicates an expected call of GenericCertificateSigningRequest
func (mr *MockCertificateSigningRequestEventHandlerMockRecorder) GenericCertificateSigningRequest(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericCertificateSigningRequest", reflect.TypeOf((*MockCertificateSigningRequestEventHandler)(nil).GenericCertificateSigningRequest), obj)
}

// MockCertificateSigningRequestEventWatcher is a mock of CertificateSigningRequestEventWatcher interface
type MockCertificateSigningRequestEventWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestEventWatcherMockRecorder
}

// MockCertificateSigningRequestEventWatcherMockRecorder is the mock recorder for MockCertificateSigningRequestEventWatcher
type MockCertificateSigningRequestEventWatcherMockRecorder struct {
	mock *MockCertificateSigningRequestEventWatcher
}

// NewMockCertificateSigningRequestEventWatcher creates a new mock instance
func NewMockCertificateSigningRequestEventWatcher(ctrl *gomock.Controller) *MockCertificateSigningRequestEventWatcher {
	mock := &MockCertificateSigningRequestEventWatcher{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestEventWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateSigningRequestEventWatcher) EXPECT() *MockCertificateSigningRequestEventWatcherMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method
func (m *MockCertificateSigningRequestEventWatcher) AddEventHandler(ctx context.Context, h controller.CertificateSigningRequestEventHandler, predicates ...predicate.Predicate) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, h}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddEventHandler", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventHandler indicates an expected call of AddEventHandler
func (mr *MockCertificateSigningRequestEventWatcherMockRecorder) AddEventHandler(ctx, h interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, h}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockCertificateSigningRequestEventWatcher)(nil).AddEventHandler), varargs...)
}
