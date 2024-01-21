package query

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type GetTodoByIdQuery struct {
	todoService services.ITodoService
}

func NewGetTodoByIdQuery(todoService services.ITodoService) *GetTodoByIdQuery {
	return &GetTodoByIdQuery{
		todoService: todoService,
	}
}

func (t *GetTodoByIdQuery) Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error) {
	todo, err := t.todoService.GetTodoById(ctx, params.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("todos not found")
		return nil, mappers.NewResponseMapper(http.StatusNotFound, errMsg, nil)
	}

	data := mappers.NewGetTodoByIdResponse(todo)

	return data, nil
}
