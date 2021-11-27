// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scipie28/test/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/scipie28/note-service-api/internal/app/model"
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

// Add mocks base method.
func (m *MockRepo) Add(arg0 api.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRepoMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepo)(nil).Add), arg0)
}

// Describe mocks base method.
func (m *MockRepo) Describe(arg0 int64) (api.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Describe", arg0)
	ret0, _ := ret[0].(api.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Describe indicates an expected call of Describe.
func (mr *MockRepoMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockRepo)(nil).Describe), arg0)
}

// MultiAdd mocks base method.
func (m *MockRepo) MultiAdd(arg0 []api.Note) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiAdd", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiAdd indicates an expected call of MultiAdd.
func (mr *MockRepoMockRecorder) MultiAdd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiAdd", reflect.TypeOf((*MockRepo)(nil).MultiAdd), arg0)
}

// Remove mocks base method.
func (m *MockRepo) Remove(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockRepoMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockRepo)(nil).Remove), arg0)
}

// Update mocks base method.
func (m *MockRepo) Update(arg0 int64, arg1 api.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepo)(nil).Update), arg0, arg1)
}
