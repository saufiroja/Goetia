// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/saufiroja/cqrs/internal/repositories (interfaces: ITodoRepository)
//
// Generated by this command:
//
//	mockgen -destination ../../mocks/mock_todo.go -package mocks github.com/saufiroja/cqrs/internal/repositories ITodoRepository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	requests "github.com/saufiroja/cqrs/internal/contracts/requests"
	responses "github.com/saufiroja/cqrs/internal/contracts/responses"
	gomock "go.uber.org/mock/gomock"
)

// MockITodoRepository is a mock of ITodoRepository interface.
type MockITodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITodoRepositoryMockRecorder
}

// MockITodoRepositoryMockRecorder is the mock recorder for MockITodoRepository.
type MockITodoRepositoryMockRecorder struct {
	mock *MockITodoRepository
}

// NewMockITodoRepository creates a new mock instance.
func NewMockITodoRepository(ctrl *gomock.Controller) *MockITodoRepository {
	mock := &MockITodoRepository{ctrl: ctrl}
	mock.recorder = &MockITodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITodoRepository) EXPECT() *MockITodoRepositoryMockRecorder {
	return m.recorder
}

// DeleteTodoById mocks base method.
func (m *MockITodoRepository) DeleteTodoById(arg0 context.Context, arg1 *sql.Tx, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodoById", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodoById indicates an expected call of DeleteTodoById.
func (mr *MockITodoRepositoryMockRecorder) DeleteTodoById(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodoById", reflect.TypeOf((*MockITodoRepository)(nil).DeleteTodoById), arg0, arg1, arg2)
}

// GetAllTodos mocks base method.
func (m *MockITodoRepository) GetAllTodos(arg0 context.Context, arg1 *sql.DB) ([]responses.GetAllTodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTodos", arg0, arg1)
	ret0, _ := ret[0].([]responses.GetAllTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTodos indicates an expected call of GetAllTodos.
func (mr *MockITodoRepositoryMockRecorder) GetAllTodos(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodos", reflect.TypeOf((*MockITodoRepository)(nil).GetAllTodos), arg0, arg1)
}

// GetTodoById mocks base method.
func (m *MockITodoRepository) GetTodoById(arg0 context.Context, arg1 *sql.DB, arg2 string) (responses.GetTodoByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoById", arg0, arg1, arg2)
	ret0, _ := ret[0].(responses.GetTodoByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoById indicates an expected call of GetTodoById.
func (mr *MockITodoRepositoryMockRecorder) GetTodoById(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoById", reflect.TypeOf((*MockITodoRepository)(nil).GetTodoById), arg0, arg1, arg2)
}

// InsertTodo mocks base method.
func (m *MockITodoRepository) InsertTodo(arg0 context.Context, arg1 *sql.Tx, arg2 *requests.TodoRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTodo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTodo indicates an expected call of InsertTodo.
func (mr *MockITodoRepositoryMockRecorder) InsertTodo(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTodo", reflect.TypeOf((*MockITodoRepository)(nil).InsertTodo), arg0, arg1, arg2)
}

// UpdateTodoById mocks base method.
func (m *MockITodoRepository) UpdateTodoById(arg0 context.Context, arg1 *sql.Tx, arg2 *requests.UpdateTodoRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodoById", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTodoById indicates an expected call of UpdateTodoById.
func (mr *MockITodoRepositoryMockRecorder) UpdateTodoById(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodoById", reflect.TypeOf((*MockITodoRepository)(nil).UpdateTodoById), arg0, arg1, arg2)
}

// UpdateTodoStatusById mocks base method.
func (m *MockITodoRepository) UpdateTodoStatusById(arg0 context.Context, arg1 *sql.Tx, arg2 *requests.UpdateTodoStatusRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodoStatusById", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTodoStatusById indicates an expected call of UpdateTodoStatusById.
func (mr *MockITodoRepositoryMockRecorder) UpdateTodoStatusById(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodoStatusById", reflect.TypeOf((*MockITodoRepository)(nil).UpdateTodoStatusById), arg0, arg1, arg2)
}