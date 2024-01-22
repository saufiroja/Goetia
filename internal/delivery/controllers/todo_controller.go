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
	ctxs, span := c.tracing.StartGlobalTracerSpan(ctx, "Controllers.InsertTodo")
	defer span.End()
	return c.handler.Command.InsertTodoCommand.Handle(ctxs, request)
}

func (c *Controllers) GetAllTodos(ctx context.Context, empty *grpc.Empty) (*grpc.TodoResponse, error) {
	ctxs, span := c.tracing.StartGlobalTracerSpan(ctx, "Controllers.GetAllTodos")
	defer span.End()
	return c.handler.Query.GetAllTodoQuery.Handle(ctxs)
}

func (c *Controllers) GetTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error) {
	ctxs, span := c.tracing.StartGlobalTracerSpan(ctx, "Controllers.GetTodoById")
	defer span.End()
	return c.handler.Query.GetTodoByIdQuery.Handle(ctxs, params)
}

func (c *Controllers) UpdateTodoById(ctx context.Context, request *grpc.UpdateTodoRequest) (*grpc.Empty, error) {
	ctxs, span := c.tracing.StartGlobalTracerSpan(ctx, "Controllers.UpdateTodoById")
	defer span.End()
	return c.handler.Command.UpdateTodoCommand.Handle(ctxs, request)
}

func (c *Controllers) UpdateTodoStatusById(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error) {
	ctxs, span := c.tracing.StartGlobalTracerSpan(ctx, "Controllers.UpdateTodoStatusById")
	defer span.End()
	return c.handler.Command.UpdateStatusTodoByIdCommand.Handle(ctxs, request)
}

func (c *Controllers) DeleteTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error) {
	ctxs, span := c.tracing.StartGlobalTracerSpan(ctx, "Controllers.DeleteTodoById")
	defer span.End()
	return c.handler.Command.DeleteTodoByIdCommand.Handle(ctxs, params)
}
