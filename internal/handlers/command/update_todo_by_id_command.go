package command

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/validator"
	"net/http"
)

type IUpdateTodoCommand interface {
	Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error
}

type UpdateTodoCommand struct {
	todoService services.ITodoService
	validation  *validator.Validation
}

func NewUpdateTodoCommand(todoService services.ITodoService, validation *validator.Validation) IUpdateTodoCommand {
	return &UpdateTodoCommand{
		todoService: todoService,
		validation:  validation,
	}
}

func (t *UpdateTodoCommand) Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	todoId := params.ByName("todoId")

	input := &requests.UpdateTodoRequest{}

	err := requests.NewRequestMapper(r, input)
	if err != nil {
		errMsg := fmt.Sprintf("failed to parse request, err: %s", err.Error())
		return mappers.NewResponseMapper(w, http.StatusBadRequest, errMsg, nil)
	}

	err = t.validation.Validate(input)
	if err != nil {
		errMsg := fmt.Sprintf("failed to validate request, err: %s", t.validation.CustomError(err))
		return mappers.NewResponseMapper(w, http.StatusBadRequest, errMsg, nil)
	}

	input.TodoId = todoId

	err = t.todoService.UpdateTodoById(input)
	if err != nil {
		errMsg := fmt.Sprintf("failed to update todo, err: %s", err.Error())
		return mappers.NewResponseMapper(w, http.StatusInternalServerError, errMsg, nil)
	}

	return mappers.NewResponseMapper(w, http.StatusOK, "success todo updated", nil)
}
