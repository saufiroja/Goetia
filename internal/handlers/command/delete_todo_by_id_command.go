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

type DeleteTodoByIdCommand struct {
	todoService services.ITodoService
	tracing     *tracing.Tracing
}

func NewDeleteTodoByIdCommand(todoService services.ITodoService, tracing *tracing.Tracing) *DeleteTodoByIdCommand {
	return &DeleteTodoByIdCommand{
		todoService: todoService,
		tracing:     tracing,
	}
}

func (t *DeleteTodoByIdCommand) Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error) {
	ctxs, span := t.tracing.StartGlobalTracerSpan(ctx, "DeleteTodoByIdCommand.Handle")
	defer span.End()

	err := t.todoService.DeleteTodoById(ctxs, params.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("failed to delete todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusInternalServerError, errMsg, nil)
	}

	return &grpc.Empty{}, nil
}
