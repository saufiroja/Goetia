package command

import (
	"context"

	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
)

type IUpdateStatusTodoByIdCommand interface {
	Handle(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error)
}

type UpdateStatusTodoByIdCommand struct {
	todoService services.ITodoService
	tracing     tracing.ITracing
}

func NewUpdateStatusTodoByIdCommand(todoService services.ITodoService, tracing tracing.ITracing) IUpdateStatusTodoByIdCommand {
	return &UpdateStatusTodoByIdCommand{
		todoService: todoService,
		tracing:     tracing,
	}
}

func (t *UpdateStatusTodoByIdCommand) Handle(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error) {
	tracer, ctx := t.tracing.StartSpan(ctx, "UpdateStatusTodoByIdCommand.Handle")
	defer tracer.Finish()

	err := t.todoService.UpdateTodoStatusById(ctx, request)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
