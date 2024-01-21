package command

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type UpdateStatusTodoByIdCommand struct {
	todoService services.ITodoService
}

func NewUpdateStatusTodoByIdCommand(todoService services.ITodoService) *UpdateStatusTodoByIdCommand {
	return &UpdateStatusTodoByIdCommand{
		todoService: todoService,
	}
}

func (t *UpdateStatusTodoByIdCommand) Handle(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error) {
	err := t.todoService.UpdateTodoStatusById(ctx, request)
	if err != nil {
		errMsg := fmt.Sprintf("failed to update todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusInternalServerError, errMsg, nil)
	}

	return &grpc.Empty{}, nil
}
