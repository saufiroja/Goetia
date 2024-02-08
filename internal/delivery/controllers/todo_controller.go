package controllers

import (
	"context"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/handlers"
	metric "github.com/saufiroja/cqrs/pkg/metrics"
	"github.com/saufiroja/cqrs/pkg/tracing"
)

type Controllers struct {
	handler *handlers.TodoHandler
	tracing *tracing.Tracing
	metrics *metric.Metrics
}

func NewControllers(handler *handlers.TodoHandler, tracing *tracing.Tracing, metrics *metric.Metrics) ITodoController {
	return &Controllers{
		handler: handler,
		tracing: tracing,
		metrics: metrics,
	}
}

func (c *Controllers) InsertTodo(ctx context.Context, request *grpc.TodoRequest) (*grpc.Empty, error) {
	c.metrics.CreateTodoRequests.WithLabelValues("POST").Inc()

	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.InsertTodo")
	defer tracer.Finish()

	c.metrics.SuccessRequests.WithLabelValues("POST").Inc()
	return c.handler.Command.InsertTodoCommand.Handle(ctx, request)
}

func (c *Controllers) GetAllTodos(ctx context.Context, empty *grpc.Empty) (*grpc.TodoResponse, error) {
	c.metrics.GetAllTodoRequests.WithLabelValues("GET").Inc()

	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.GetAllTodos")
	defer tracer.Finish()

	c.metrics.SuccessRequests.WithLabelValues("GET").Inc()
	return c.handler.Query.GetAllTodoQuery.Handle(ctx)
}

func (c *Controllers) GetTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error) {
	c.metrics.GetTodoRequests.WithLabelValues("GET").Inc()

	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.GetTodoById")
	defer tracer.Finish()

	c.metrics.SuccessRequests.WithLabelValues("GET").Inc()
	return c.handler.Query.GetTodoByIdQuery.Handle(ctx, params)
}

func (c *Controllers) UpdateTodoById(ctx context.Context, request *grpc.UpdateTodoRequest) (*grpc.Empty, error) {
	c.metrics.UpdateTodoRequests.WithLabelValues("PUT").Inc()

	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.UpdateTodoById")
	defer tracer.Finish()

	c.metrics.SuccessRequests.WithLabelValues("PUT").Inc()
	return c.handler.Command.UpdateTodoCommand.Handle(ctx, request)
}

func (c *Controllers) UpdateTodoStatusById(ctx context.Context, request *grpc.UpdateTodoStatusRequest) (*grpc.Empty, error) {
	c.metrics.UpdateStatusTodoRequests.WithLabelValues("PATCH").Inc()

	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.UpdateTodoStatusById")
	defer tracer.Finish()

	c.metrics.SuccessRequests.WithLabelValues("PATCH").Inc()
	return c.handler.Command.UpdateStatusTodoByIdCommand.Handle(ctx, request)
}

func (c *Controllers) DeleteTodoById(ctx context.Context, params *grpc.TodoParams) (*grpc.Empty, error) {
	c.metrics.DeleteTodoRequests.WithLabelValues("DELETE").Inc()

	tracer, ctx := c.tracing.StartSpan(ctx, "Controllers.DeleteTodoById")
	defer tracer.Finish()

	c.metrics.SuccessRequests.WithLabelValues("DELETE").Inc()
	return c.handler.Command.DeleteTodoByIdCommand.Handle(ctx, params)
}
