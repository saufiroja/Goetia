package services

import (
	"context"

	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/grpc"
)

type ITodoService interface {
	InsertTodo(ctx context.Context, request *requests.TodoRequest) error
	GetAllTodo(ctx context.Context) ([]responses.GetAllTodoResponse, error)
	GetTodoById(ctx context.Context, todoId string) (*responses.GetTodoByIdResponse, error)
	UpdateTodoById(ctx context.Context, request *requests.UpdateTodoRequest) error
	UpdateTodoStatusById(ctx context.Context, request *grpc.UpdateTodoStatusRequest) error
	DeleteTodoById(ctx context.Context, todoId string) error
}
