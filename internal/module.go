package internal

import (
	"github.com/saufiroja/cqrs/internal/delivery/controllers"
	"github.com/saufiroja/cqrs/internal/handlers"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/redis"
	"github.com/sirupsen/logrus"
)

func NewModule(db *database.Postgres, log *logrus.Logger, redis *redis.Redis) controllers.ITodoController {
	// application
	todoRepository := repositories.NewRepository()
	todoService := services.NewService(db, log, todoRepository, redis)

	// handlers
	todoHandler := handlers.NewTodoHandler(todoService)

	// controllers
	todoControllers := controllers.NewControllers(*todoHandler)

	return todoControllers
}
