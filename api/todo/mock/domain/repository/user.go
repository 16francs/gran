// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository/user.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	domain "github.com/calmato/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Authentication mocks base method
func (m *MockUserRepository) Authentication(ctx context.Context) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authentication", ctx)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authentication indicates an expected call of Authentication
func (mr *MockUserRepositoryMockRecorder) Authentication(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authentication", reflect.TypeOf((*MockUserRepository)(nil).Authentication), ctx)
}

// Update mocks base method
func (m *MockUserRepository) Update(ctx context.Context, u *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockUserRepositoryMockRecorder) Update(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), ctx, u)
}
