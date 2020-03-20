// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository/board.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	domain "github.com/16francs/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBoardRepository is a mock of BoardRepository interface
type MockBoardRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBoardRepositoryMockRecorder
}

// MockBoardRepositoryMockRecorder is the mock recorder for MockBoardRepository
type MockBoardRepositoryMockRecorder struct {
	mock *MockBoardRepository
}

// NewMockBoardRepository creates a new mock instance
func NewMockBoardRepository(ctrl *gomock.Controller) *MockBoardRepository {
	mock := &MockBoardRepository{ctrl: ctrl}
	mock.recorder = &MockBoardRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBoardRepository) EXPECT() *MockBoardRepositoryMockRecorder {
	return m.recorder
}

// Index mocks base method
func (m *MockBoardRepository) Index(ctx context.Context, groupID string) ([]*domain.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", ctx, groupID)
	ret0, _ := ret[0].([]*domain.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Index indicates an expected call of Index
func (mr *MockBoardRepositoryMockRecorder) Index(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockBoardRepository)(nil).Index), ctx, groupID)
}

// Show mocks base method
func (m *MockBoardRepository) Show(ctx context.Context, groupID, boardID string) (*domain.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, groupID, boardID)
	ret0, _ := ret[0].(*domain.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show
func (mr *MockBoardRepositoryMockRecorder) Show(ctx, groupID, boardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockBoardRepository)(nil).Show), ctx, groupID, boardID)
}

// Create mocks base method
func (m *MockBoardRepository) Create(ctx context.Context, b *domain.Board) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockBoardRepositoryMockRecorder) Create(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBoardRepository)(nil).Create), ctx, b)
}

// IndexBoardList mocks base method
func (m *MockBoardRepository) IndexBoardList(ctx context.Context, groupID, boardID string) ([]*domain.BoardList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexBoardList", ctx, groupID, boardID)
	ret0, _ := ret[0].([]*domain.BoardList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IndexBoardList indicates an expected call of IndexBoardList
func (mr *MockBoardRepositoryMockRecorder) IndexBoardList(ctx, groupID, boardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexBoardList", reflect.TypeOf((*MockBoardRepository)(nil).IndexBoardList), ctx, groupID, boardID)
}