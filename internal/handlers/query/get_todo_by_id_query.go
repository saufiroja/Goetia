package query

import (
	"context"
	"fmt"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"net/http"
)

type GetTodoByIdQuery struct {
	todoService services.ITodoService
	tracing     *tracing.Tracing
}

func NewGetTodoByIdQuery(todoService services.ITodoService, tracing *tracing.Tracing) *GetTodoByIdQuery {
	return &GetTodoByIdQuery{
		todoService: todoService,
		tracing:     tracing,
	}
}

func (t *GetTodoByIdQuery) Handle(ctx context.Context, params *grpc.TodoParams) (*grpc.GetTodoResponse, error) {
	ctxs, span := t.tracing.StartGlobalTracerSpan(ctx, "GetTodoByIdQuery.Handle")
	defer span.End()

	todo, err := t.todoService.GetTodoById(ctxs, params.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("todos not found")
		return nil, mappers.NewResponseMapper(http.StatusNotFound, errMsg, nil)
	}

	data := mappers.NewGetTodoByIdResponse(todo)

	return data, nil
}
