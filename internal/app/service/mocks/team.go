// Code generated by MockGen. DO NOT EDIT.
// Source: team.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "polygames/internal/app/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTeamRepository is a mock of TeamRepository interface.
type MockTeamRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTeamRepositoryMockRecorder
}

// MockTeamRepositoryMockRecorder is the mock recorder for MockTeamRepository.
type MockTeamRepositoryMockRecorder struct {
	mock *MockTeamRepository
}

// NewMockTeamRepository creates a new mock instance.
func NewMockTeamRepository(ctrl *gomock.Controller) *MockTeamRepository {
	mock := &MockTeamRepository{ctrl: ctrl}
	mock.recorder = &MockTeamRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamRepository) EXPECT() *MockTeamRepositoryMockRecorder {
	return m.recorder
}

// CheckTeamUniqueness mocks base method.
func (m *MockTeamRepository) CheckTeamUniqueness(ctx context.Context, title string) (*domain.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTeamUniqueness", ctx, title)
	ret0, _ := ret[0].(*domain.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTeamUniqueness indicates an expected call of CheckTeamUniqueness.
func (mr *MockTeamRepositoryMockRecorder) CheckTeamUniqueness(ctx, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTeamUniqueness", reflect.TypeOf((*MockTeamRepository)(nil).CheckTeamUniqueness), ctx, title)
}

// CreateTeam mocks base method.
func (m *MockTeamRepository) CreateTeam(ctx context.Context, team *domain.Team) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", ctx, team)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockTeamRepositoryMockRecorder) CreateTeam(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockTeamRepository)(nil).CreateTeam), ctx, team)
}

// DisableTeam mocks base method.
func (m *MockTeamRepository) DisableTeam(ctx context.Context, teamID int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableTeam", ctx, teamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableTeam indicates an expected call of DisableTeam.
func (mr *MockTeamRepositoryMockRecorder) DisableTeam(ctx, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableTeam", reflect.TypeOf((*MockTeamRepository)(nil).DisableTeam), ctx, teamID)
}

// EnableTeam mocks base method.
func (m *MockTeamRepository) EnableTeam(ctx context.Context, teamID int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableTeam", ctx, teamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableTeam indicates an expected call of EnableTeam.
func (mr *MockTeamRepositoryMockRecorder) EnableTeam(ctx, teamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableTeam", reflect.TypeOf((*MockTeamRepository)(nil).EnableTeam), ctx, teamID)
}

// GetTeam mocks base method.
func (m *MockTeamRepository) GetTeam(ctx context.Context, id int32) (*domain.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeam", ctx, id)
	ret0, _ := ret[0].(*domain.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeam indicates an expected call of GetTeam.
func (mr *MockTeamRepositoryMockRecorder) GetTeam(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeam", reflect.TypeOf((*MockTeamRepository)(nil).GetTeam), ctx, id)
}

// GetTeams mocks base method.
func (m *MockTeamRepository) GetTeams(ctx context.Context) ([]domain.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeams", ctx)
	ret0, _ := ret[0].([]domain.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeams indicates an expected call of GetTeams.
func (mr *MockTeamRepositoryMockRecorder) GetTeams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeams", reflect.TypeOf((*MockTeamRepository)(nil).GetTeams), ctx)
}

// SetTeamImage mocks base method.
func (m *MockTeamRepository) SetTeamImage(ctx context.Context, teamID int32, imageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTeamImage", ctx, teamID, imageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTeamImage indicates an expected call of SetTeamImage.
func (mr *MockTeamRepositoryMockRecorder) SetTeamImage(ctx, teamID, imageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTeamImage", reflect.TypeOf((*MockTeamRepository)(nil).SetTeamImage), ctx, teamID, imageID)
}

// UpdateTeam mocks base method.
func (m *MockTeamRepository) UpdateTeam(ctx context.Context, team *domain.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeam", ctx, team)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTeam indicates an expected call of UpdateTeam.
func (mr *MockTeamRepositoryMockRecorder) UpdateTeam(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeam", reflect.TypeOf((*MockTeamRepository)(nil).UpdateTeam), ctx, team)
}
