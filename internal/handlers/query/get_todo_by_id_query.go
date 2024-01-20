package query

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type IGetTodoByIdQuery interface {
	Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error
}

type GetTodoByIdQuery struct {
	todoService services.ITodoService
}

func NewGetTodoByIdQuery(todoService services.ITodoService) IGetTodoByIdQuery {
	return &GetTodoByIdQuery{
		todoService: todoService,
	}
}

func (t *GetTodoByIdQuery) Handle(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	todoId := params.ByName("todoId")

	todo, err := t.todoService.GetTodoById(todoId)
	if err != nil {
		errMsg := fmt.Sprintf("todo not found")
		return mappers.NewResponseMapper(w, http.StatusNotFound, errMsg, nil)
	}

	result := mappers.NewResponseMapper(w, http.StatusOK, "success get todo by id", todo)
	return result
}
