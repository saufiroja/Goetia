package controllers

import (
	"context"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/handlers"
	"github.com/saufiroja/cqrs/pkg/tracing"
)

type Controllers struct {
	handler *handlers.TodoHandler
	tracing *tracing.Tracing
}

func NewControllers(handler *handlers.TodoHandler, tracing *tracing.Tracing) ITodoController {
	return &Controllers{
		handler: handler,
		tracing: tracing,
	}
}

func (c *Controllers) InsertTodo(ctx context.Context, request *grpc.TodoRequest) (*grpc.Empty, error) {
	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.InsertTodo")
	defer tracer.Finish()
	return c.handler.Command.InsertTodoCommand.Handle(ctx, request)
}

func (c *Controllers) GetAllTodos(ctx context.Context, empty *grpc.Empty) (*grpc.TodoResponse, error) {
	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.GetAllTodos")
	defer tracer.Finish()
	return c.handler.Query.GetAllTodoQuery.Handle(ctx)
}

func (c *Controllers) GetTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error) {
	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.GetTodoById")
	defer tracer.Finish()
	return c.handler.Query.GetTodoByIdQuery.Handle(ctx, params)
}

func (c *Controllers) UpdateTodoById(ctx context.Context, request *grpc.UpdateTodoRequest) (*grpc.Empty, error) {
	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.UpdateTodoById")
	defer tracer.Finish()
	return c.handler.Command.UpdateTodoCommand.Handle(ctx, request)
}

func (c *Controllers) UpdateTodoStatusById(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error) {
	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.UpdateTodoStatusById")
	defer tracer.Finish()
	return c.handler.Command.UpdateStatusTodoByIdCommand.Handle(ctx, request)
}

func (c *Controllers) DeleteTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error) {
	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.DeleteTodoById")
	defer tracer.Finish()
	return c.handler.Command.DeleteTodoByIdCommand.Handle(ctx, params)
}
