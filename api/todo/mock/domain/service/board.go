// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/service/board.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	domain "github.com/calmato/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBoardService is a mock of BoardService interface
type MockBoardService struct {
	ctrl     *gomock.Controller
	recorder *MockBoardServiceMockRecorder
}

// MockBoardServiceMockRecorder is the mock recorder for MockBoardService
type MockBoardServiceMockRecorder struct {
	mock *MockBoardService
}

// NewMockBoardService creates a new mock instance
func NewMockBoardService(ctrl *gomock.Controller) *MockBoardService {
	mock := &MockBoardService{ctrl: ctrl}
	mock.recorder = &MockBoardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBoardService) EXPECT() *MockBoardServiceMockRecorder {
	return m.recorder
}

// Index mocks base method
func (m *MockBoardService) Index(ctx context.Context, groupID string) ([]*domain.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", ctx, groupID)
	ret0, _ := ret[0].([]*domain.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Index indicates an expected call of Index
func (mr *MockBoardServiceMockRecorder) Index(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockBoardService)(nil).Index), ctx, groupID)
}

// Show mocks base method
func (m *MockBoardService) Show(ctx context.Context, groupID, boardID string) (*domain.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, groupID, boardID)
	ret0, _ := ret[0].(*domain.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show
func (mr *MockBoardServiceMockRecorder) Show(ctx, groupID, boardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockBoardService)(nil).Show), ctx, groupID, boardID)
}

// Create mocks base method
func (m *MockBoardService) Create(ctx context.Context, b *domain.Board) (*domain.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, b)
	ret0, _ := ret[0].(*domain.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBoardServiceMockRecorder) Create(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBoardService)(nil).Create), ctx, b)
}

// UploadThumbnail mocks base method
func (m *MockBoardService) UploadThumbnail(ctx context.Context, data []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadThumbnail", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadThumbnail indicates an expected call of UploadThumbnail
func (mr *MockBoardServiceMockRecorder) UploadThumbnail(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadThumbnail", reflect.TypeOf((*MockBoardService)(nil).UploadThumbnail), ctx, data)
}

// Exists mocks base method
func (m *MockBoardService) Exists(ctx context.Context, groupID, boardID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, groupID, boardID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockBoardServiceMockRecorder) Exists(ctx, groupID, boardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockBoardService)(nil).Exists), ctx, groupID, boardID)
}

// ShowBoardList mocks base method
func (m *MockBoardService) ShowBoardList(ctx context.Context, groupID, boardID, boardListID string) (*domain.BoardList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowBoardList", ctx, groupID, boardID, boardListID)
	ret0, _ := ret[0].(*domain.BoardList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowBoardList indicates an expected call of ShowBoardList
func (mr *MockBoardServiceMockRecorder) ShowBoardList(ctx, groupID, boardID, boardListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowBoardList", reflect.TypeOf((*MockBoardService)(nil).ShowBoardList), ctx, groupID, boardID, boardListID)
}

// CreateBoardList mocks base method
func (m *MockBoardService) CreateBoardList(ctx context.Context, groupID, boardID string, bl *domain.BoardList) (*domain.BoardList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBoardList", ctx, groupID, boardID, bl)
	ret0, _ := ret[0].(*domain.BoardList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBoardList indicates an expected call of CreateBoardList
func (mr *MockBoardServiceMockRecorder) CreateBoardList(ctx, groupID, boardID, bl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBoardList", reflect.TypeOf((*MockBoardService)(nil).CreateBoardList), ctx, groupID, boardID, bl)
}

// UpdateBoardList mocks base method
func (m *MockBoardService) UpdateBoardList(ctx context.Context, groupID, boardID string, bl *domain.BoardList) (*domain.BoardList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBoardList", ctx, groupID, boardID, bl)
	ret0, _ := ret[0].(*domain.BoardList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBoardList indicates an expected call of UpdateBoardList
func (mr *MockBoardServiceMockRecorder) UpdateBoardList(ctx, groupID, boardID, bl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBoardList", reflect.TypeOf((*MockBoardService)(nil).UpdateBoardList), ctx, groupID, boardID, bl)
}

// ExistsBoardList mocks base method
func (m *MockBoardService) ExistsBoardList(ctx context.Context, groupID, boardID, boardListID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsBoardList", ctx, groupID, boardID, boardListID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsBoardList indicates an expected call of ExistsBoardList
func (mr *MockBoardServiceMockRecorder) ExistsBoardList(ctx, groupID, boardID, boardListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsBoardList", reflect.TypeOf((*MockBoardService)(nil).ExistsBoardList), ctx, groupID, boardID, boardListID)
}

// UpdateKanban mocks base method
func (m *MockBoardService) UpdateKanban(ctx context.Context, groupID, boardID string, b *domain.Board) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKanban", ctx, groupID, boardID, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateKanban indicates an expected call of UpdateKanban
func (mr *MockBoardServiceMockRecorder) UpdateKanban(ctx, groupID, boardID, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKanban", reflect.TypeOf((*MockBoardService)(nil).UpdateKanban), ctx, groupID, boardID, b)
}
