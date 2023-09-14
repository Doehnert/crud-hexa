// Code generated by MockGen. DO NOT EDIT.
// Source: src/application/port/output/user_port.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/Doehnert/crud-hexa/src/application/domain"
	rest_errors "github.com/Doehnert/crud-hexa/src/configuration/rest_errors"
	gomock "go.uber.org/mock/gomock"
)

// MockUserPort is a mock of UserPort interface.
type MockUserPort struct {
	ctrl     *gomock.Controller
	recorder *MockUserPortMockRecorder
}

// MockUserPortMockRecorder is the mock recorder for MockUserPort.
type MockUserPortMockRecorder struct {
	mock *MockUserPort
}

// NewMockUserPort creates a new mock instance.
func NewMockUserPort(ctrl *gomock.Controller) *MockUserPort {
	mock := &MockUserPort{ctrl: ctrl}
	mock.recorder = &MockUserPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserPort) EXPECT() *MockUserPortMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserPort) CreateUser(userDomain domain.UserDomain) (*domain.UserDomain, *rest_errors.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", userDomain)
	ret0, _ := ret[0].(*domain.UserDomain)
	ret1, _ := ret[1].(*rest_errors.RestErr)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserPortMockRecorder) CreateUser(userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserPort)(nil).CreateUser), userDomain)
}

// DeleteUser mocks base method.
func (m *MockUserPort) DeleteUser(userId string) *rest_errors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userId)
	ret0, _ := ret[0].(*rest_errors.RestErr)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserPortMockRecorder) DeleteUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserPort)(nil).DeleteUser), userId)
}

// FindUserByEmail mocks base method.
func (m *MockUserPort) FindUserByEmail(email string) (*domain.UserDomain, *rest_errors.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(*domain.UserDomain)
	ret1, _ := ret[1].(*rest_errors.RestErr)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserPortMockRecorder) FindUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserPort)(nil).FindUserByEmail), email)
}

// FindUserByEmailAndPassword mocks base method.
func (m *MockUserPort) FindUserByEmailAndPassword(email, password string) (*domain.UserDomain, *rest_errors.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailAndPassword", email, password)
	ret0, _ := ret[0].(*domain.UserDomain)
	ret1, _ := ret[1].(*rest_errors.RestErr)
	return ret0, ret1
}

// FindUserByEmailAndPassword indicates an expected call of FindUserByEmailAndPassword.
func (mr *MockUserPortMockRecorder) FindUserByEmailAndPassword(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailAndPassword", reflect.TypeOf((*MockUserPort)(nil).FindUserByEmailAndPassword), email, password)
}

// FindUserByID mocks base method.
func (m *MockUserPort) FindUserByID(id string) (*domain.UserDomain, *rest_errors.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", id)
	ret0, _ := ret[0].(*domain.UserDomain)
	ret1, _ := ret[1].(*rest_errors.RestErr)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockUserPortMockRecorder) FindUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockUserPort)(nil).FindUserByID), id)
}

// UpdateUser mocks base method.
func (m *MockUserPort) UpdateUser(userId string, userDomain domain.UserDomain) *rest_errors.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, userDomain)
	ret0, _ := ret[0].(*rest_errors.RestErr)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserPortMockRecorder) UpdateUser(userId, userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserPort)(nil).UpdateUser), userId, userDomain)
}