package query

import (
	"context"

	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
)

type IGetTodoByIdQuery interface {
	Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error)
}

type GetTodoByIdQuery struct {
	todoService services.ITodoService
	tracing     tracing.ITracing
}

func NewGetTodoByIdQuery(todoService services.ITodoService, tracing tracing.ITracing) IGetTodoByIdQuery {
	return &GetTodoByIdQuery{
		todoService: todoService,
		tracing:     tracing,
	}
}

func (t *GetTodoByIdQuery) Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error) {
	tracer, ctx := t.tracing.StartSpan(ctx, "GetTodoByIdQuery.Handle")
	defer tracer.Finish()

	todo, err := t.todoService.GetTodoById(ctx, params.TodoId)
	if err != nil {
		return nil, err
	}

	data := mappers.NewGetTodoByIdResponse(todo)

	return data, nil
}
