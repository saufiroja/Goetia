package command

import (
	"context"
	"fmt"
	"time"

	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"github.com/saufiroja/cqrs/pkg/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	input := &requests.UpdateTodoRequest{
		TodoId:      request.TodoId,
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
		UpdatedAt:   time.Unix(request.UpdatedAt, 0),
	}

	err := t.validation.Validate(request)
	if err != nil {
		errMsg := fmt.Sprintf("failed to validate request, err: %s", t.validation.CustomError(err))
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	err = t.todoService.UpdateTodoById(ctx, input)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
