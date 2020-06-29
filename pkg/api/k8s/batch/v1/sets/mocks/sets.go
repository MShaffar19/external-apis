// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1sets "github.com/solo-io/external-apis/pkg/api/k8s/batch/v1/sets"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1 "k8s.io/api/batch/v1"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockJobSet is a mock of JobSet interface
type MockJobSet struct {
	ctrl     *gomock.Controller
	recorder *MockJobSetMockRecorder
}

// MockJobSetMockRecorder is the mock recorder for MockJobSet
type MockJobSetMockRecorder struct {
	mock *MockJobSet
}

// NewMockJobSet creates a new mock instance
func NewMockJobSet(ctrl *gomock.Controller) *MockJobSet {
	mock := &MockJobSet{ctrl: ctrl}
	mock.recorder = &MockJobSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobSet) EXPECT() *MockJobSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockJobSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockJobSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockJobSet)(nil).Keys))
}

// List mocks base method
func (m *MockJobSet) List() []*v1.Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*v1.Job)
	return ret0
}

// List indicates an expected call of List
func (mr *MockJobSetMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockJobSet)(nil).List))
}

// Map mocks base method
func (m *MockJobSet) Map() map[string]*v1.Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.Job)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockJobSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockJobSet)(nil).Map))
}

// Insert mocks base method
func (m *MockJobSet) Insert(job ...*v1.Job) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range job {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockJobSetMockRecorder) Insert(job ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockJobSet)(nil).Insert), job...)
}

// Equal mocks base method
func (m *MockJobSet) Equal(jobSet v1sets.JobSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", jobSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockJobSetMockRecorder) Equal(jobSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockJobSet)(nil).Equal), jobSet)
}

// Has mocks base method
func (m *MockJobSet) Has(job *v1.Job) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", job)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockJobSetMockRecorder) Has(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockJobSet)(nil).Has), job)
}

// Delete mocks base method
func (m *MockJobSet) Delete(job *v1.Job) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", job)
}

// Delete indicates an expected call of Delete
func (mr *MockJobSetMockRecorder) Delete(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockJobSet)(nil).Delete), job)
}

// Union mocks base method
func (m *MockJobSet) Union(set v1sets.JobSet) v1sets.JobSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.JobSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockJobSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockJobSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockJobSet) Difference(set v1sets.JobSet) v1sets.JobSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.JobSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockJobSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockJobSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockJobSet) Intersection(set v1sets.JobSet) v1sets.JobSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.JobSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockJobSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockJobSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockJobSet) Find(id ezkube.ResourceId) (*v1.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockJobSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockJobSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockJobSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockJobSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockJobSet)(nil).Length))
}
