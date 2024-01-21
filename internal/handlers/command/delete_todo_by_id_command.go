package command

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type DeleteTodoByIdCommand struct {
	todoService services.ITodoService
}

func NewDeleteTodoByIdCommand(todoService services.ITodoService) *DeleteTodoByIdCommand {
	return &DeleteTodoByIdCommand{
		todoService: todoService,
	}
}

func (t *DeleteTodoByIdCommand) Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error) {
	err := t.todoService.DeleteTodoById(ctx, params.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("failed to delete todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusInternalServerError, errMsg, nil)
	}

	return &grpc.Empty{}, nil
}
