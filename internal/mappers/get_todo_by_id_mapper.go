package mappers

import (
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/grpc"
)

func NewGetTodoByIdResponse(todo *responses.GetTodoByIdResponse) *grpc.GetTodoResponse {
	return &grpc.GetTodoResponse{
		TodoId:    todo.TodoId,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt.Unix(),
		UpdatedAt: todo.UpdatedAt.Unix(),
	}
}
