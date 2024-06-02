package command

import (
	"context"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/logger"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"github.com/saufiroja/cqrs/pkg/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IInsertTodoCommand interface {
	Handle(ctx context.Context, request *grpc.TodoRequest) (*grpc.Empty, error)
}

type InsertTodoCommand struct {
	todoService services.ITodoService
	validation  *validator.Validation
	tracing     tracing.ITracing
	log         logger.ILogger
}

func NewInsertTodoCommand(todoService services.ITodoService, validation *validator.Validation,
	tracing tracing.ITracing,
	log logger.ILogger) IInsertTodoCommand {
	return &InsertTodoCommand{
		todoService: todoService,
		validation:  validation,
		tracing:     tracing,
		log:         log,
	}
}

func (t *InsertTodoCommand) Handle(ctx context.Context, request *grpc.TodoRequest) (*grpc.Empty, error) {
	tracer, ctx := t.tracing.StartSpan(ctx, "InsertTodoCommand.Handle")
	defer tracer.Finish()

	input := &requests.TodoRequest{
		TodoId:      ulid.Make().String(),
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
		CreatedAt:   time.Unix(request.CreatedAt, 0),
		UpdatedAt:   time.Unix(request.UpdatedAt, 0),
	}

	err := t.validation.Validate(input)
	if err != nil {
		t.log.StartLogger("InsertTodoCommand", "Handle").Error(err.Error())
		errMsg := fmt.Sprintf("failed to validate request, err: %s", t.validation.CustomError(err))
		return nil, status.Error(codes.InvalidArgument, errMsg)
	}

	err = t.todoService.InsertTodo(ctx, input)
	if err != nil {
		return nil, err
	}

	return &grpc.Empty{}, nil
}
