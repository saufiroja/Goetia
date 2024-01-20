package command

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type IUpdateStatusTodoByIdCommand interface {
	Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error
}

type UpdateStatusTodoByIdCommand struct {
	todoService services.ITodoService
}

func NewUpdateStatusTodoByIdCommand(todoService services.ITodoService) IUpdateStatusTodoByIdCommand {
	return &UpdateStatusTodoByIdCommand{
		todoService: todoService,
	}
}

func (t *UpdateStatusTodoByIdCommand) Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	todoId := params.ByName("todoId")

	input := &requests.UpdateTodoStatusRequest{}

	err := requests.NewRequestMapper(r, input)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parse request, err: %s", err.Error())
		return mappers.NewResponseMapper(w, http.StatusBadRequest, errMsg, nil)
	}

	input.TodoId = todoId

	err = t.todoService.UpdateTodoStatusById(input)
	if err != nil {
		errMsg := fmt.Sprintf("failed to update todo, err: %s", err.Error())
		return mappers.NewResponseMapper(w, http.StatusInternalServerError, errMsg, nil)
	}

	return mappers.NewResponseMapper(w, http.StatusOK, "success update todo status", nil)
}
