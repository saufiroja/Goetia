package app

import (
	"github.com/saufiroja/cqrs/internal/delivery/controllers"
	"github.com/saufiroja/cqrs/internal/handlers"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/redis"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"github.com/sirupsen/logrus"
)

func NewModule(db *database.Postgres, log *logrus.Logger, redis *redis.Redis, tracing *tracing.Tracing) controllers.ITodoController {
	// application
	todoRepository := repositories.NewRepository(tracing)
	todoService := services.NewService(db, log, todoRepository, redis, tracing)

	// handlers
	todoHandler := handlers.NewTodoHandler(todoService, tracing)

	// controllers
	todoControllers := controllers.NewControllers(todoHandler, tracing)

	return todoControllers
}
