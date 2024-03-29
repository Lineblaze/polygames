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

// MockGameService is a mock of GameService interface.
type MockGameService struct {
	ctrl     *gomock.Controller
	recorder *MockGameServiceMockRecorder
}

// MockGameServiceMockRecorder is the mock recorder for MockGameService.
type MockGameServiceMockRecorder struct {
	mock *MockGameService
}

// NewMockGameService creates a new mock instance.
func NewMockGameService(ctrl *gomock.Controller) *MockGameService {
	mock := &MockGameService{ctrl: ctrl}
	mock.recorder = &MockGameServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameService) EXPECT() *MockGameServiceMockRecorder {
	return m.recorder
}

// CreateGame mocks base method.
func (m *MockGameService) CreateGame(ctx context.Context, game *domain.Game) (*domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", ctx, game)
	ret0, _ := ret[0].(*domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGame indicates an expected call of CreateGame.
func (mr *MockGameServiceMockRecorder) CreateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*MockGameService)(nil).CreateGame), ctx, game)
}

// DeleteGame mocks base method.
func (m *MockGameService) DeleteGame(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGame", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGame indicates an expected call of DeleteGame.
func (mr *MockGameServiceMockRecorder) DeleteGame(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGame", reflect.TypeOf((*MockGameService)(nil).DeleteGame), ctx, id)
}

// GetGame mocks base method.
func (m *MockGameService) GetGame(ctx context.Context, id int32) (*domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGame", ctx, id)
	ret0, _ := ret[0].(*domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGame indicates an expected call of GetGame.
func (mr *MockGameServiceMockRecorder) GetGame(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGame", reflect.TypeOf((*MockGameService)(nil).GetGame), ctx, id)
}

// GetGameFile mocks base method.
func (m *MockGameService) GetGameFile(ctx context.Context, gameID int32) (*domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameFile", ctx, gameID)
	ret0, _ := ret[0].(*domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameFile indicates an expected call of GetGameFile.
func (mr *MockGameServiceMockRecorder) GetGameFile(ctx, gameID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameFile", reflect.TypeOf((*MockGameService)(nil).GetGameFile), ctx, gameID)
}

// GetGameImage mocks base method.
func (m *MockGameService) GetGameImage(ctx context.Context, gameID int32) (*domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameImage", ctx, gameID)
	ret0, _ := ret[0].(*domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameImage indicates an expected call of GetGameImage.
func (mr *MockGameServiceMockRecorder) GetGameImage(ctx, gameID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameImage", reflect.TypeOf((*MockGameService)(nil).GetGameImage), ctx, gameID)
}

// GetGames mocks base method.
func (m *MockGameService) GetGames(ctx context.Context) ([]domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGames", ctx)
	ret0, _ := ret[0].([]domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGames indicates an expected call of GetGames.
func (mr *MockGameServiceMockRecorder) GetGames(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGames", reflect.TypeOf((*MockGameService)(nil).GetGames), ctx)
}

// SetGameFile mocks base method.
func (m *MockGameService) SetGameFile(ctx context.Context, gameID int32, file []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGameFile", ctx, gameID, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGameFile indicates an expected call of SetGameFile.
func (mr *MockGameServiceMockRecorder) SetGameFile(ctx, gameID, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGameFile", reflect.TypeOf((*MockGameService)(nil).SetGameFile), ctx, gameID, file)
}

// SetGameImage mocks base method.
func (m *MockGameService) SetGameImage(ctx context.Context, gameID int32, img []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGameImage", ctx, gameID, img)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGameImage indicates an expected call of SetGameImage.
func (mr *MockGameServiceMockRecorder) SetGameImage(ctx, gameID, img interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGameImage", reflect.TypeOf((*MockGameService)(nil).SetGameImage), ctx, gameID, img)
}

// UpdateGame mocks base method.
func (m *MockGameService) UpdateGame(ctx context.Context, game *domain.Game) (*domain.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGame", ctx, game)
	ret0, _ := ret[0].(*domain.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGame indicates an expected call of UpdateGame.
func (mr *MockGameServiceMockRecorder) UpdateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGame", reflect.TypeOf((*MockGameService)(nil).UpdateGame), ctx, game)
}
