package handlers

import (
	"github.com/saufiroja/cqrs/internal/handlers/command"
	"github.com/saufiroja/cqrs/internal/handlers/query"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/validator"
)

type TodoHandler struct {
	Query   query.TodoQuery
	Command command.TodoCommand
}

func NewTodoHandler(todoService services.ITodoService) *TodoHandler {
	validation := validator.NewValidation()

	getAllTodoQuery := query.NewGetAllTodoQuery(todoService)
	getTodoByIdQeury := query.NewGetTodoByIdQuery(todoService)

	insertTodoCommand := command.NewInsertTodoCommand(todoService, validation)
	updateTodoByIdCommand := command.NewUpdateTodoCommand(todoService, validation)
	updateTodoStatusByIdCommand := command.NewUpdateStatusTodoByIdCommand(todoService)
	deleteTodoByIdCommand := command.NewDeleteTodoByIdCommand(todoService)

	todoQuery := query.NewTodoQuery(getAllTodoQuery, getTodoByIdQeury)
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
