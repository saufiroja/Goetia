package app

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/saufiroja/cqrs/config"
	"github.com/saufiroja/cqrs/internal/delivery/controllers"
	"github.com/saufiroja/cqrs/internal/handlers/event"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/logger"
	metric "github.com/saufiroja/cqrs/pkg/metrics"
	"github.com/saufiroja/cqrs/pkg/redis"
	"github.com/saufiroja/cqrs/pkg/tracing"
)

type Module struct {
	controllers.ITodoController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) StartModule(conf *config.AppConfig, reg *prometheus.Registry) {
	// configuration
	log := logger.NewLogger()

	// metrics
	metrics := metric.NewMetrics(reg, conf.App.ServiceName)
	trace := tracing.NewTracing(conf, log)
	m.StartMetrics(reg, conf)

	// database
	redisCli := redis.NewRedis(conf, log)
	db := database.NewPostgres(conf, log)
	db.StartDatabase()

	// application
	todoRepository := repositories.NewRepository(trace)
	todoService := services.NewService(db, log, todoRepository, redisCli, trace)

	// handlers
	todoHandler := event.NewTodoHandler(todoService, trace)

	// controllers
	todoControllers := controllers.NewControllers(todoHandler, trace, metrics)

	m.ITodoController = todoControllers
}

func (m *Module) StartMetrics(reg *prometheus.Registry, conf *config.AppConfig) {
	metricsGauge := metric.NewMetricsGauge(reg, conf.App.ServiceName)
	metricsGauge.SetTotalCPU()
	metricsGauge.SetTotalMemory()
}
