package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ITodoController interface {
	InsertTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
