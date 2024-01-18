package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/handlers/command"
	"github.com/saufiroja/cqrs/internal/handlers/query"
	"net/http"
)

type Controllers struct {
	command command.InsertTodoCommand
	Query   query.GetAllTodoQuery
}

func NewControllers(command command.InsertTodoCommand, query query.GetAllTodoQuery) TodoController {
	return &Controllers{
		command: command,
		Query:   query,
	}
}

func (c *Controllers) InsertTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.command.Handle(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (c *Controllers) GetAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := c.Query.Handle(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
