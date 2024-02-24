// Code generated by MockGen. DO NOT EDIT.
// Source: game.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "polygames/internal/app/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGameRepository is a mock of GameRepository interface.
type MockGameRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGameRepositoryMockRecorder
}

// MockGameRepositoryMockRecorder is the mock recorder for MockGameRepository.
type MockGameRepositoryMockRecorder struct {
	mock *MockGameRepository
}

// NewMockGameRepository creates a new mock instance.
func NewMockGameRepository(ctrl *gomock.Controller) *MockGameRepository {
	mock := &MockGameRepository{ctrl: ctrl}
	mock.recorder = &MockGameRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameRepository) EXPECT() *MockGameRepositoryMockRecorder {
	return m.recorder
}

// CheckGameUniqueness mocks base method.
func (m *MockGameRepository) CheckGameUniqueness(ctx context.Context, title string) (*domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckGameUniqueness", ctx, title)
	ret0, _ := ret[0].(*domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckGameUniqueness indicates an expected call of CheckGameUniqueness.
func (mr *MockGameRepositoryMockRecorder) CheckGameUniqueness(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckGameUniqueness", reflect.TypeOf((*MockGameRepository)(nil).CheckGameUniqueness), ctx, title)
}

// CreateGame mocks base method.
func (m *MockGameRepository) CreateGame(ctx context.Context, game *domain.Game) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", ctx, game)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGame indicates an expected call of CreateGame.
func (mr *MockGameRepositoryMockRecorder) CreateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*MockGameRepository)(nil).CreateGame), ctx, game)
}

// DeleteGame mocks base method.
func (m *MockGameRepository) DeleteGame(ctx context.Context, gameID int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGame", ctx, gameID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGame indicates an expected call of DeleteGame.
func (mr *MockGameRepositoryMockRecorder) DeleteGame(ctx, gameID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGame", reflect.TypeOf((*MockGameRepository)(nil).DeleteGame), ctx, gameID)
}

// GetGame mocks base method.
func (m *MockGameRepository) GetGame(ctx context.Context, id int32) (*domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGame", ctx, id)
	ret0, _ := ret[0].(*domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGame indicates an expected call of GetGame.
func (mr *MockGameRepositoryMockRecorder) GetGame(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGame", reflect.TypeOf((*MockGameRepository)(nil).GetGame), ctx, id)
}

// GetGames mocks base method.
func (m *MockGameRepository) GetGames(ctx context.Context) ([]domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGames", ctx)
	ret0, _ := ret[0].([]domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGames indicates an expected call of GetGames.
func (mr *MockGameRepositoryMockRecorder) GetGames(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGames", reflect.TypeOf((*MockGameRepository)(nil).GetGames), ctx)
}

// SetGameFile mocks base method.
func (m *MockGameRepository) SetGameFile(ctx context.Context, gameID int32, fileID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGameFile", ctx, gameID, fileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGameFile indicates an expected call of SetGameFile.
func (mr *MockGameRepositoryMockRecorder) SetGameFile(ctx, gameID, fileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGameFile", reflect.TypeOf((*MockGameRepository)(nil).SetGameFile), ctx, gameID, fileID)
}

// SetGameImage mocks base method.
func (m *MockGameRepository) SetGameImage(ctx context.Context, gameID int32, imageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGameImage", ctx, gameID, imageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGameImage indicates an expected call of SetGameImage.
func (mr *MockGameRepositoryMockRecorder) SetGameImage(ctx, gameID, imageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGameImage", reflect.TypeOf((*MockGameRepository)(nil).SetGameImage), ctx, gameID, imageID)
}

// UpdateGame mocks base method.
func (m *MockGameRepository) UpdateGame(ctx context.Context, game *domain.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGame", ctx, game)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGame indicates an expected call of UpdateGame.
func (mr *MockGameRepositoryMockRecorder) UpdateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGame", reflect.TypeOf((*MockGameRepository)(nil).UpdateGame), ctx, game)
}
