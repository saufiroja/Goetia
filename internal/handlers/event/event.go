package event

import (
	"github.com/saufiroja/cqrs/internal/handlers/command"
	"github.com/saufiroja/cqrs/internal/handlers/query"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/logger"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"github.com/saufiroja/cqrs/pkg/validator"
)

type TodoHandler struct {
	Query   query.TodoQuery
	Command command.TodoCommand
}

func NewTodoHandler(todoService services.ITodoService, tracing tracing.ITracing, log logger.ILogger) *TodoHandler {
	validation := validator.NewValidation()

	getAllTodoQuery := query.NewGetAllTodoQuery(todoService, tracing)
	getTodoByIdQuery := query.NewGetTodoByIdQuery(todoService, tracing)

	insertTodoCommand := command.NewInsertTodoCommand(todoService, validation, tracing, log)
	updateTodoByIdCommand := command.NewUpdateTodoCommand(todoService, validation, tracing)
	updateTodoStatusByIdCommand := command.NewUpdateStatusTodoByIdCommand(todoService, tracing)
	deleteTodoByIdCommand := command.NewDeleteTodoByIdCommand(todoService, tracing)

	todoQuery := query.NewTodoQuery(getAllTodoQuery, getTodoByIdQuery)
	todoCommands := command.NewTodoCommand(
		insertTodoCommand,
		updateTodoByIdCommand,
		updateTodoStatusByIdCommand,
		deleteTodoByIdCommand,
	)

	return &TodoHandler{
		Query:   *todoQuery,
		Command: *todoCommands,
	}
}
