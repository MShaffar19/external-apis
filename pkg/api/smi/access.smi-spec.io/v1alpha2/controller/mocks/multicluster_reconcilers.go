// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	controller "github.com/solo-io/external-apis/pkg/api/smi/access.smi-spec.io/v1alpha2/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterTrafficTargetReconciler is a mock of MulticlusterTrafficTargetReconciler interface
type MockMulticlusterTrafficTargetReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTrafficTargetReconcilerMockRecorder
}

// MockMulticlusterTrafficTargetReconcilerMockRecorder is the mock recorder for MockMulticlusterTrafficTargetReconciler
type MockMulticlusterTrafficTargetReconcilerMockRecorder struct {
	mock *MockMulticlusterTrafficTargetReconciler
}

// NewMockMulticlusterTrafficTargetReconciler creates a new mock instance
func NewMockMulticlusterTrafficTargetReconciler(ctrl *gomock.Controller) *MockMulticlusterTrafficTargetReconciler {
	mock := &MockMulticlusterTrafficTargetReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTrafficTargetReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterTrafficTargetReconciler) EXPECT() *MockMulticlusterTrafficTargetReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficTarget mocks base method
func (m *MockMulticlusterTrafficTargetReconciler) ReconcileTrafficTarget(clusterName string, obj *v1alpha2.TrafficTarget) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficTarget", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileTrafficTarget indicates an expected call of ReconcileTrafficTarget
func (mr *MockMulticlusterTrafficTargetReconcilerMockRecorder) ReconcileTrafficTarget(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficTarget", reflect.TypeOf((*MockMulticlusterTrafficTargetReconciler)(nil).ReconcileTrafficTarget), clusterName, obj)
}

// MockMulticlusterTrafficTargetDeletionReconciler is a mock of MulticlusterTrafficTargetDeletionReconciler interface
type MockMulticlusterTrafficTargetDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTrafficTargetDeletionReconcilerMockRecorder
}

// MockMulticlusterTrafficTargetDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterTrafficTargetDeletionReconciler
type MockMulticlusterTrafficTargetDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterTrafficTargetDeletionReconciler
}

// NewMockMulticlusterTrafficTargetDeletionReconciler creates a new mock instance
func NewMockMulticlusterTrafficTargetDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterTrafficTargetDeletionReconciler {
	mock := &MockMulticlusterTrafficTargetDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTrafficTargetDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterTrafficTargetDeletionReconciler) EXPECT() *MockMulticlusterTrafficTargetDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileTrafficTargetDeletion mocks base method
func (m *MockMulticlusterTrafficTargetDeletionReconciler) ReconcileTrafficTargetDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileTrafficTargetDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileTrafficTargetDeletion indicates an expected call of ReconcileTrafficTargetDeletion
func (mr *MockMulticlusterTrafficTargetDeletionReconcilerMockRecorder) ReconcileTrafficTargetDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileTrafficTargetDeletion", reflect.TypeOf((*MockMulticlusterTrafficTargetDeletionReconciler)(nil).ReconcileTrafficTargetDeletion), clusterName, req)
}

// MockMulticlusterTrafficTargetReconcileLoop is a mock of MulticlusterTrafficTargetReconcileLoop interface
type MockMulticlusterTrafficTargetReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterTrafficTargetReconcileLoopMockRecorder
}

// MockMulticlusterTrafficTargetReconcileLoopMockRecorder is the mock recorder for MockMulticlusterTrafficTargetReconcileLoop
type MockMulticlusterTrafficTargetReconcileLoopMockRecorder struct {
	mock *MockMulticlusterTrafficTargetReconcileLoop
}

// NewMockMulticlusterTrafficTargetReconcileLoop creates a new mock instance
func NewMockMulticlusterTrafficTargetReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterTrafficTargetReconcileLoop {
	mock := &MockMulticlusterTrafficTargetReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterTrafficTargetReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterTrafficTargetReconcileLoop) EXPECT() *MockMulticlusterTrafficTargetReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterTrafficTargetReconciler mocks base method
func (m *MockMulticlusterTrafficTargetReconcileLoop) AddMulticlusterTrafficTargetReconciler(ctx context.Context, rec controller.MulticlusterTrafficTargetReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterTrafficTargetReconciler", varargs...)
}

// AddMulticlusterTrafficTargetReconciler indicates an expected call of AddMulticlusterTrafficTargetReconciler
func (mr *MockMulticlusterTrafficTargetReconcileLoopMockRecorder) AddMulticlusterTrafficTargetReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterTrafficTargetReconciler", reflect.TypeOf((*MockMulticlusterTrafficTargetReconcileLoop)(nil).AddMulticlusterTrafficTargetReconciler), varargs...)
}
