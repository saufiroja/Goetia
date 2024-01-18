package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/delivery/http/controllers"
	"github.com/saufiroja/cqrs/internal/middlewares"
)

func NewRouter(todoController controllers.TodoController, router *httprouter.Router) *httprouter.Router {
	group := "/api/v1"

	router.GET(fmt.Sprintf("%s/todos", group), middlewares.LoggerMiddleware(todoController.GetAllTodo))
	router.POST(fmt.Sprintf("%s/todos", group), middlewares.LoggerMiddleware(todoController.InsertTodo))

	return router
}
