// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/celestiaorg/celestia-node/share (interfaces: Availability)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	header "github.com/celestiaorg/celestia-node/header"
	gomock "github.com/golang/mock/gomock"
)

// MockAvailability is a mock of Availability interface.
type MockAvailability struct {
	ctrl     *gomock.Controller
	recorder *MockAvailabilityMockRecorder
}

// MockAvailabilityMockRecorder is the mock recorder for MockAvailability.
type MockAvailabilityMockRecorder struct {
	mock *MockAvailability
}

// NewMockAvailability creates a new mock instance.
func NewMockAvailability(ctrl *gomock.Controller) *MockAvailability {
	mock := &MockAvailability{ctrl: ctrl}
	mock.recorder = &MockAvailabilityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvailability) EXPECT() *MockAvailabilityMockRecorder {
	return m.recorder
}

// SharesAvailable mocks base method.
func (m *MockAvailability) SharesAvailable(arg0 context.Context, arg1 *header.ExtendedHeader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SharesAvailable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SharesAvailable indicates an expected call of SharesAvailable.
func (mr *MockAvailabilityMockRecorder) SharesAvailable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SharesAvailable", reflect.TypeOf((*MockAvailability)(nil).SharesAvailable), arg0, arg1)
}