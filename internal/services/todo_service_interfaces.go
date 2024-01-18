package services

import (
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
)

type TodoService interface {
	InsertTodo(input *requests.TodoRequest) error
	GetAllTodo() ([]responses.GetAllTodoResponse, error)
}
