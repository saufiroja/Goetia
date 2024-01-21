package mappers

import (
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/grpc"
)

func NewGetAllTodoResponse(todos []responses.GetAllTodoResponse) *grpc.TodoResponse {
	var todoResponse grpc.TodoResponse
	for _, todo := range todos {
		todoResponse.Todos = append(todoResponse.Todos, &grpc.GetTodoResponse{
			TodoId:    todo.TodoId,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt.Unix(),
			UpdatedAt: todo.UpdatedAt.Unix(),
		})
	}

	return &todoResponse
}
