// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/validation/user.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	context "context"
	domain "github.com/calmato/gran/api/user/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserDomainValidation is a mock of UserDomainValidation interface
type MockUserDomainValidation struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainValidationMockRecorder
}

// MockUserDomainValidationMockRecorder is the mock recorder for MockUserDomainValidation
type MockUserDomainValidationMockRecorder struct {
	mock *MockUserDomainValidation
}

// NewMockUserDomainValidation creates a new mock instance
func NewMockUserDomainValidation(ctrl *gomock.Controller) *MockUserDomainValidation {
	mock := &MockUserDomainValidation{ctrl: ctrl}
	mock.recorder = &MockUserDomainValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDomainValidation) EXPECT() *MockUserDomainValidationMockRecorder {
	return m.recorder
}

// User mocks base method
func (m *MockUserDomainValidation) User(ctx context.Context, u *domain.User) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", ctx, u)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// User indicates an expected call of User
func (mr *MockUserDomainValidationMockRecorder) User(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockUserDomainValidation)(nil).User), ctx, u)
}
