package command

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type IDeleteTodoByIdCommand interface {
	Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error
}

type DeleteTodoByIdCommand struct {
	todoService services.ITodoService
}

func NewDeleteTodoByIdCommand(todoService services.ITodoService) IDeleteTodoByIdCommand {
	return &DeleteTodoByIdCommand{
		todoService: todoService,
	}
}

func (t *DeleteTodoByIdCommand) Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	todoId := params.ByName("todoId")

	err := t.todoService.DeleteTodoById(todoId)
	if err != nil {
		errMsg := fmt.Sprintf("failed to delete todo, err: %s", err.Error())
		return mappers.NewResponseMapper(w, http.StatusInternalServerError, errMsg, nil)
	}

	return mappers.NewResponseMapper(w, http.StatusOK, "success todo deleted", nil)
}
