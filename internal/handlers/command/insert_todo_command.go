package command

import (
	"fmt"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/validator"
	"net/http"
)

type IInsertTodoCommand interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

type InsertTodoCommand struct {
	todoService services.ITodoService
	validation  *validator.Validation
}

func NewInsertTodoCommand(todoService services.ITodoService, validation *validator.Validation) IInsertTodoCommand {
	return &InsertTodoCommand{
		todoService: todoService,
		validation:  validation,
	}
}

func (t *InsertTodoCommand) Handle(w http.ResponseWriter, r *http.Request) error {
	input := &requests.TodoRequest{}

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

	err = t.todoService.InsertTodo(input)
	if err != nil {
		errMsg := fmt.Sprintf("failed to insert todo, err: %s", err.Error())
		return mappers.NewResponseMapper(w, http.StatusInternalServerError, errMsg, nil)
	}

	return mappers.NewResponseMapper(w, http.StatusCreated, "success todo created", nil)
}
