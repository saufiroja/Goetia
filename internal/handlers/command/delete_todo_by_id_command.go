package command

import (
	"context"

	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
)

type IDeleteTodoByIdCommand interface {
	Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error)
}

type DeleteTodoByIdCommand struct {
	todoService services.ITodoService
	tracing     tracing.ITracing
}

func NewDeleteTodoByIdCommand(todoService services.ITodoService, tracing tracing.ITracing) IDeleteTodoByIdCommand {
	return &DeleteTodoByIdCommand{
		todoService: todoService,
		tracing:     tracing,
	}
}

func (t *DeleteTodoByIdCommand) Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error) {
	tracer, ctx := t.tracing.StartSpan(ctx, "DeleteTodoByIdCommand.Handle")
	defer tracer.Finish()

	err := t.todoService.DeleteTodoById(ctx, params.TodoId)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
