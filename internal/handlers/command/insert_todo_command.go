package command

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/validator"
	"net/http"
)

type InsertTodoCommand struct {
	todoService services.ITodoService
	validation  *validator.Validation
}

func NewInsertTodoCommand(todoService services.ITodoService, validation *validator.Validation) *InsertTodoCommand {
	return &InsertTodoCommand{
		todoService: todoService,
		validation:  validation,
	}
}

func (t *InsertTodoCommand) Handle(ctx context.Context, request *grpc.TodoRequest) (*grpc.Empty, error) {
	err := t.validation.Validate(request)
	if err != nil {
		errMsg := fmt.Sprintf("failed to validate request, err: %s", t.validation.CustomError(err))
		return nil, mappers.NewResponseMapper(http.StatusBadRequest, errMsg, nil)
	}

	err = t.todoService.InsertTodo(ctx, request)
	if err != nil {
		errMsg := fmt.Sprintf("failed to insert todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusInternalServerError, errMsg, nil)
	}

	return &grpc.Empty{}, nil
}
