// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetByShort mocks base method.
func (m *MockRepository) GetByShort(short string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByShort", short)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByShort indicates an expected call of GetByShort.
func (mr *MockRepositoryMockRecorder) GetByShort(short interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByShort", reflect.TypeOf((*MockRepository)(nil).GetByShort), short)
}

// GetByURL mocks base method.
func (m *MockRepository) GetByURL(url string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByURL", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByURL indicates an expected call of GetByURL.
func (mr *MockRepositoryMockRecorder) GetByURL(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByURL", reflect.TypeOf((*MockRepository)(nil).GetByURL), url)
}

// Insert mocks base method.
func (m *MockRepository) Insert(url, short string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", url, short)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockRepositoryMockRecorder) Insert(url, short interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRepository)(nil).Insert), url, short)
}
