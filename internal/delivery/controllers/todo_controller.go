package controllers

import (
	"context"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/handlers"
)

type Controllers struct {
	handler handlers.TodoHandler
}

func NewControllers(handler handlers.TodoHandler) ITodoController {
	return &Controllers{
		handler: handler,
	}
}

func (c *Controllers) InsertTodo(ctx context.Context, request *grpc.TodoRequest) (*grpc.Empty, error) {
	return c.handler.Command.InsertTodoCommand.Handle(ctx, request)
}

func (c *Controllers) GetAllTodos(ctx context.Context, empty *grpc.Empty) (*grpc.TodoResponse, error) {
	return c.handler.Query.GetAllTodoQuery.Handle(ctx)
}

func (c *Controllers) GetTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error) {
	return c.handler.Query.GetTodoByIdQuery.Handle(ctx, params)
}

func (c *Controllers) UpdateTodoById(ctx context.Context, request *grpc.UpdateTodoRequest) (*grpc.Empty, error) {
	return c.handler.Command.UpdateTodoCommand.Handle(ctx, request)
}

func (c *Controllers) UpdateTodoStatusById(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error) {
	return c.handler.Command.UpdateStatusTodoByIdCommand.Handle(ctx, request)
}

func (c *Controllers) DeleteTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error) {
	return c.handler.Command.DeleteTodoByIdCommand.Handle(ctx, params)
}
