package query

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"net/http"
)

type GetAllTodoQuery struct {
	todoService services.ITodoService
	tracing     *tracing.Tracing
}

func NewGetAllTodoQuery(todoService services.ITodoService, tracing *tracing.Tracing) *GetAllTodoQuery {
	return &GetAllTodoQuery{
		todoService: todoService,
		tracing:     tracing,
	}
}

func (t *GetAllTodoQuery) Handle(ctx context.Context) (*grpc.TodoResponse, error) {
	tracer, ctx := t.tracing.StartSpan(ctx, "GetAllTodoQuery.Handle")
	defer tracer.Finish()

	todos, err := t.todoService.GetAllTodo(ctx)
	if err != nil {
		errMsg := fmt.Sprintf("failed to get all todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusOK, errMsg, nil)
	}

	data := mappers.NewGetAllTodoResponse(todos)

	return data, nil
}
