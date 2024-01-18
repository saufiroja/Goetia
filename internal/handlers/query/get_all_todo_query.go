package query

import (
	"fmt"
	"github.com/saufiroja/cqrs/internal/mappers"
	"github.com/saufiroja/cqrs/internal/services"
	"net/http"
)

type GetAllTodoQuery struct {
	todoService services.TodoService
}

func NewGetAllTodoQuery(todoService services.TodoService) *GetAllTodoQuery {
	return &GetAllTodoQuery{
		todoService: todoService,
	}
}

func (t *GetAllTodoQuery) Handle(w http.ResponseWriter, r *http.Request) error {
	todos, err := t.todoService.GetAllTodo()
	if err != nil {
		errMsg := fmt.Sprintf("failed to get all todo, err: %s", err.Error())
		return mappers.NewResponseMapper(w, http.StatusOK, errMsg, nil)
	}

	result := mappers.NewResponseMapper(w, http.StatusOK, "success get all todo", todos)
	return result
}
