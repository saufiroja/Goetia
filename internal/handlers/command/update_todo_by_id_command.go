package command

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"github.com/saufiroja/cqrs/pkg/validator"
	"net/http"
)

type IUpdateTodoCommand interface {
	Handle(ctx context.Context, request *grpc.UpdateTodoRequest) (*grpc.Empty, error)
}

type UpdateTodoCommand struct {
	todoService services.ITodoService
	validation  *validator.Validation
	tracing     tracing.ITracing
}

func NewUpdateTodoCommand(todoService services.ITodoService, validation *validator.Validation, tracing tracing.ITracing) IUpdateTodoCommand {
	return &UpdateTodoCommand{
		todoService: todoService,
		validation:  validation,
		tracing:     tracing,
	}
}

func (t *UpdateTodoCommand) Handle(ctx context.Context, request *grpc.UpdateTodoRequest) (*grpc.Empty, error) {
	tracer, ctx := t.tracing.StartSpan(ctx, "UpdateTodoCommand.Handle")
	defer tracer.Finish()

	err := t.validation.Validate(request)
	if err != nil {
		errMsg := fmt.Sprintf("failed to validate request, err: %s", t.validation.CustomError(err))
		return nil, mappers.NewResponseMapper(http.StatusBadRequest, errMsg, nil)
	}

	err = t.todoService.UpdateTodoById(ctx, request)
	if err != nil {
		errMsg := fmt.Sprintf("failed to update todos, err: %s", err.Error())
		return nil, mappers.NewResponseMapper(http.StatusInternalServerError, errMsg, nil)
	}

	return &grpc.Empty{}, nil
}
