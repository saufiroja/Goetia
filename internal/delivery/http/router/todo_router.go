package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/saufiroja/cqrs/internal/delivery/http/controllers"
	"github.com/saufiroja/cqrs/internal/middlewares"
)

func NewRouter(todoController controllers.ITodoController, router *httprouter.Router) *httprouter.Router {
	group := "/api/v1"

	router.GET(fmt.Sprintf("%s/todos", group), middlewares.LoggerMiddleware(todoController.GetAllTodo))
	router.POST(fmt.Sprintf("%s/todos", group), middlewares.LoggerMiddleware(todoController.InsertTodo))
	router.GET(fmt.Sprintf("%s/todos/:todoId", group), middlewares.LoggerMiddleware(todoController.GetTodoById))
	router.PUT(fmt.Sprintf("%s/todos/:todoId", group), middlewares.LoggerMiddleware(todoController.UpdateTodoById))
	router.PUT(fmt.Sprintf("%s/todos/:todoId/status", group), middlewares.LoggerMiddleware(todoController.UpdateStatusTodoById))
	router.DELETE(fmt.Sprintf("%s/todos/:todoId", group), middlewares.LoggerMiddleware(todoController.DeleteTodoById))

	return router
}
