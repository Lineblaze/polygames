// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "polygames/internal/app/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CheckUserUniqueness mocks base method.
func (m *MockUserRepository) CheckUserUniqueness(ctx context.Context, username, email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserUniqueness", ctx, username, email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserUniqueness indicates an expected call of CheckUserUniqueness.
func (mr *MockUserRepositoryMockRecorder) CheckUserUniqueness(ctx, username, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserUniqueness", reflect.TypeOf((*MockUserRepository)(nil).CheckUserUniqueness), ctx, username, email)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, user)
}

// DisableUser mocks base method.
func (m *MockUserRepository) DisableUser(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUser indicates an expected call of DisableUser.
func (mr *MockUserRepositoryMockRecorder) DisableUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUser", reflect.TypeOf((*MockUserRepository)(nil).DisableUser), ctx, id)
}

// GetActiveUser mocks base method.
func (m *MockUserRepository) GetActiveUser(ctx context.Context, id int32) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveUser", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveUser indicates an expected call of GetActiveUser.
func (mr *MockUserRepositoryMockRecorder) GetActiveUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveUser", reflect.TypeOf((*MockUserRepository)(nil).GetActiveUser), ctx, id)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(ctx context.Context, id int32) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), ctx, id)
}

// GetUserByLogin mocks base method.
func (m *MockUserRepository) GetUserByLogin(ctx context.Context, login string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLogin", ctx, login)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByLogin indicates an expected call of GetUserByLogin.
func (mr *MockUserRepositoryMockRecorder) GetUserByLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLogin", reflect.TypeOf((*MockUserRepository)(nil).GetUserByLogin), ctx, login)
}

// GetUsers mocks base method.
func (m *MockUserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepositoryMockRecorder) GetUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepository)(nil).GetUsers), ctx)
}

// SetUserImage mocks base method.
func (m *MockUserRepository) SetUserImage(ctx context.Context, userID int32, imageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserImage", ctx, userID, imageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserImage indicates an expected call of SetUserImage.
func (mr *MockUserRepositoryMockRecorder) SetUserImage(ctx, userID, imageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserImage", reflect.TypeOf((*MockUserRepository)(nil).SetUserImage), ctx, userID, imageID)
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), ctx, user)
}
