package command

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"net/http"
)

type UpdateStatusTodoByIdCommand struct {
	todoService services.ITodoService
	tracing     *tracing.Tracing
}

func NewUpdateStatusTodoByIdCommand(todoService services.ITodoService, tracing *tracing.Tracing) *UpdateStatusTodoByIdCommand {
	return &UpdateStatusTodoByIdCommand{
		todoService: todoService,
		tracing:     tracing,
	}
}

func (t *UpdateStatusTodoByIdCommand) Handle(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error) {
	ctxs, span := t.tracing.StartGlobalTracerSpan(ctx, "UpdateStatusTodoByIdCommand.Handle")
	defer span.End()

	err := t.todoService.UpdateTodoStatusById(ctxs, request)
	if err != nil {
		errMsg := fmt.Sprintf("failed to update todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusInternalServerError, errMsg, nil)
	}

	return &grpc.Empty{}, nil
}
