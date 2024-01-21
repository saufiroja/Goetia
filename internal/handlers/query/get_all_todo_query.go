package query

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type GetAllTodoQuery struct {
	todoService services.ITodoService
}

func NewGetAllTodoQuery(todoService services.ITodoService) *GetAllTodoQuery {
	return &GetAllTodoQuery{
		todoService: todoService,
	}
}

func (t *GetAllTodoQuery) Handle(ctx context.Context) (*grpc.TodoResponse, error) {
	todos, err := t.todoService.GetAllTodo(ctx)
	if err != nil {
		errMsg := fmt.Sprintf("failed to get all todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusOK, errMsg, nil)
	}

	data := mappers.NewGetAllTodoResponse(todos)

	return data, nil
}
