// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-core/ami (interfaces: AMI)

// Package mock is a generated GoMock package.
package mock

import (
	v1 "github.com/baetyl/baetyl-go/spec/v1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAMI is a mock of AMI interface
type MockAMI struct {
	ctrl     *gomock.Controller
	recorder *MockAMIMockRecorder
}

// MockAMIMockRecorder is the mock recorder for MockAMI
type MockAMIMockRecorder struct {
	mock *MockAMI
}

// NewMockAMI creates a new mock instance
func NewMockAMI(ctrl *gomock.Controller) *MockAMI {
	mock := &MockAMI{ctrl: ctrl}
	mock.recorder = &MockAMIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAMI) EXPECT() *MockAMIMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockAMI) Apply(arg0 string, arg1 []v1.AppInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockAMIMockRecorder) Apply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockAMI)(nil).Apply), arg0, arg1)
}

// Collect mocks base method
func (m *MockAMI) Collect(arg0 string) (v1.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", arg0)
	ret0, _ := ret[0].(v1.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect
func (mr *MockAMIMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockAMI)(nil).Collect), arg0)
}
