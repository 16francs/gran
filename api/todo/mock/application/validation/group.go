// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/validation/group.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	request "github.com/calmato/gran/api/todo/internal/application/request"
	domain "github.com/calmato/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGroupRequestValidation is a mock of GroupRequestValidation interface
type MockGroupRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRequestValidationMockRecorder
}

// MockGroupRequestValidationMockRecorder is the mock recorder for MockGroupRequestValidation
type MockGroupRequestValidationMockRecorder struct {
	mock *MockGroupRequestValidation
}

// NewMockGroupRequestValidation creates a new mock instance
func NewMockGroupRequestValidation(ctrl *gomock.Controller) *MockGroupRequestValidation {
	mock := &MockGroupRequestValidation{ctrl: ctrl}
	mock.recorder = &MockGroupRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGroupRequestValidation) EXPECT() *MockGroupRequestValidationMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method
func (m *MockGroupRequestValidation) CreateGroup(req *request.CreateGroup) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", req)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// CreateGroup indicates an expected call of CreateGroup
func (mr *MockGroupRequestValidationMockRecorder) CreateGroup(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockGroupRequestValidation)(nil).CreateGroup), req)
}

// UpdateGroup mocks base method
func (m *MockGroupRequestValidation) UpdateGroup(req *request.UpdateGroup) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", req)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// UpdateGroup indicates an expected call of UpdateGroup
func (mr *MockGroupRequestValidationMockRecorder) UpdateGroup(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*MockGroupRequestValidation)(nil).UpdateGroup), req)
}

// InviteUsers mocks base method
func (m *MockGroupRequestValidation) InviteUsers(req *request.InviteUsers) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteUsers", req)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// InviteUsers indicates an expected call of InviteUsers
func (mr *MockGroupRequestValidationMockRecorder) InviteUsers(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUsers", reflect.TypeOf((*MockGroupRequestValidation)(nil).InviteUsers), req)
}
