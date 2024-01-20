package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ITodoController interface {
	InsertTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateStatusTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteTodoById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
