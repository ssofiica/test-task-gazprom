// Code generated by MockGen. DO NOT EDIT.
// Source: ./repo.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/ssofiica/test-task-gazprom/internal/entity"
	dto "github.com/ssofiica/test-task-gazprom/internal/entity/dto"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepo) CreateUser(ctx context.Context, user *dto.SignUp) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepoMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepo)(nil).CreateUser), ctx, user)
}

// DeleteSessionValue mocks base method.
func (m *MockRepo) DeleteSessionValue(ctx context.Context, sessionId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSessionValue", ctx, sessionId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSessionValue indicates an expected call of DeleteSessionValue.
func (mr *MockRepoMockRecorder) DeleteSessionValue(ctx, sessionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSessionValue", reflect.TypeOf((*MockRepo)(nil).DeleteSessionValue), ctx, sessionId)
}

// GetSessionValue mocks base method.
func (m *MockRepo) GetSessionValue(ctx context.Context, sessionId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionValue", ctx, sessionId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionValue indicates an expected call of GetSessionValue.
func (mr *MockRepoMockRecorder) GetSessionValue(ctx, sessionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionValue", reflect.TypeOf((*MockRepo)(nil).GetSessionValue), ctx, sessionId)
}

// SetSessionValue mocks base method.
func (m *MockRepo) SetSessionValue(ctx context.Context, session *entity.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSessionValue", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSessionValue indicates an expected call of SetSessionValue.
func (mr *MockRepoMockRecorder) SetSessionValue(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSessionValue", reflect.TypeOf((*MockRepo)(nil).SetSessionValue), ctx, session)
}