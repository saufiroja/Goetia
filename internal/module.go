package internal

import (
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/internal/delivery/http/controllers"
	"github.com/saufiroja/cqrs/internal/handlers/command"
	"github.com/saufiroja/cqrs/internal/handlers/query"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/sirupsen/logrus"
)

func NewModule(conf *config.AppConfig, db *database.Postgres, log *logrus.Logger) controllers.TodoController {
	// application
	todoRepository := repositories.NewRepository()
	todoService := services.NewService(db, log, todoRepository)

	// command
	todoCommand := command.NewTodoCommand(todoService)
	// query
	todoQuery := query.NewGetAllTodoQuery(todoService)

	// controllers
	todoControllers := controllers.NewControllers(*todoCommand, *todoQuery)

	return todoControllers
}
