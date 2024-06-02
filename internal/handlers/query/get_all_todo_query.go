package query

import (
	"context"

	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
)

type IGetAllTodoQuery interface {
	Handle(ctx context.Context) (*grpc.TodoResponse, error)
}

type GetAllTodoQuery struct {
	todoService services.ITodoService
	tracing     tracing.ITracing
}

func NewGetAllTodoQuery(todoService services.ITodoService, tracing tracing.ITracing) IGetAllTodoQuery {
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
		return nil, err
	}

	data := mappers.NewGetAllTodoResponse(todos)

	return data, nil
}
