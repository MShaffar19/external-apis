// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1beta1sets is a generated GoMock package.
package mock_v1beta1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1sets "github.com/solo-io/external-apis/pkg/api/k8s/certificates.k8s.io/v1beta1/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1beta1 "k8s.io/api/certificates/v1beta1"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockCertificateSigningRequestSet is a mock of CertificateSigningRequestSet interface
type MockCertificateSigningRequestSet struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateSigningRequestSetMockRecorder
}

// MockCertificateSigningRequestSetMockRecorder is the mock recorder for MockCertificateSigningRequestSet
type MockCertificateSigningRequestSetMockRecorder struct {
	mock *MockCertificateSigningRequestSet
}

// NewMockCertificateSigningRequestSet creates a new mock instance
func NewMockCertificateSigningRequestSet(ctrl *gomock.Controller) *MockCertificateSigningRequestSet {
	mock := &MockCertificateSigningRequestSet{ctrl: ctrl}
	mock.recorder = &MockCertificateSigningRequestSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCertificateSigningRequestSet) EXPECT() *MockCertificateSigningRequestSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockCertificateSigningRequestSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockCertificateSigningRequestSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Keys))
}

// List mocks base method
func (m *MockCertificateSigningRequestSet) List() []*v1beta1.CertificateSigningRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1beta1.CertificateSigningRequest)
	return ret0
}

// List indicates an expected call of List
func (mr *MockCertificateSigningRequestSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).List))
}

// Map mocks base method
func (m *MockCertificateSigningRequestSet) Map() map[string]*v1beta1.CertificateSigningRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1beta1.CertificateSigningRequest)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockCertificateSigningRequestSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Map))
}

// Insert mocks base method
func (m *MockCertificateSigningRequestSet) Insert(certificateSigningRequest ...*v1beta1.CertificateSigningRequest) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range certificateSigningRequest {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockCertificateSigningRequestSetMockRecorder) Insert(certificateSigningRequest ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Insert), certificateSigningRequest...)
}

// Equal mocks base method
func (m *MockCertificateSigningRequestSet) Equal(certificateSigningRequestSet v1beta1sets.CertificateSigningRequestSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", certificateSigningRequestSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockCertificateSigningRequestSetMockRecorder) Equal(certificateSigningRequestSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Equal), certificateSigningRequestSet)
}

// Has mocks base method
func (m *MockCertificateSigningRequestSet) Has(certificateSigningRequest *v1beta1.CertificateSigningRequest) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", certificateSigningRequest)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockCertificateSigningRequestSetMockRecorder) Has(certificateSigningRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Has), certificateSigningRequest)
}

// Delete mocks base method
func (m *MockCertificateSigningRequestSet) Delete(certificateSigningRequest *v1beta1.CertificateSigningRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", certificateSigningRequest)
}

// Delete indicates an expected call of Delete
func (mr *MockCertificateSigningRequestSetMockRecorder) Delete(certificateSigningRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Delete), certificateSigningRequest)
}

// Union mocks base method
func (m *MockCertificateSigningRequestSet) Union(set v1beta1sets.CertificateSigningRequestSet) v1beta1sets.CertificateSigningRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1beta1sets.CertificateSigningRequestSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockCertificateSigningRequestSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockCertificateSigningRequestSet) Difference(set v1beta1sets.CertificateSigningRequestSet) v1beta1sets.CertificateSigningRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1beta1sets.CertificateSigningRequestSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockCertificateSigningRequestSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockCertificateSigningRequestSet) Intersection(set v1beta1sets.CertificateSigningRequestSet) v1beta1sets.CertificateSigningRequestSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1beta1sets.CertificateSigningRequestSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockCertificateSigningRequestSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockCertificateSigningRequestSet) Find(id ezkube.ResourceId) (*v1beta1.CertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1beta1.CertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockCertificateSigningRequestSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockCertificateSigningRequestSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockCertificateSigningRequestSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockCertificateSigningRequestSet)(nil).Length))
}
