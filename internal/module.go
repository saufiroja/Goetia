package internal

import (
	"github.com/saufiroja/cqrs/internal/delivery/http/controllers"
	"github.com/saufiroja/cqrs/internal/handlers"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/sirupsen/logrus"
)

func NewModule(db *database.Postgres, log *logrus.Logger) controllers.ITodoController {
	// application
	todoRepository := repositories.NewRepository()
	todoService := services.NewService(db, log, todoRepository)

	// handlers
	todoHandler := handlers.NewTodoHandler(todoService)

	// controllers
	todoControllers := controllers.NewControllers(*todoHandler)

	return todoControllers
}
