package services

import (
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
)

type ITodoService interface {
	InsertTodo(input *requests.TodoRequest) error
	GetAllTodo() ([]responses.GetAllTodoResponse, error)
}
