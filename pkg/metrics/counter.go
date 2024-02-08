package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	SuccessRequests          *prometheus.CounterVec
	FailRequests             *prometheus.CounterVec
	GetAllTodoRequests       *prometheus.CounterVec
	GetTodoRequests          *prometheus.CounterVec
	CreateTodoRequests       *prometheus.CounterVec
	UpdateTodoRequests       *prometheus.CounterVec
	UpdateStatusTodoRequests *prometheus.CounterVec
	DeleteTodoRequests       *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer, service string) *Metrics {
	m := &Metrics{
		SuccessRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_success_requests",
			Help:      "The total number of successful requests",
		}, []string{"method"}),
		FailRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_fail_requests",
			Help:      "The total number of failed requests",
		}, []string{"method"}),
		GetAllTodoRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_get_all_todo_requests",
			Help:      "The total number of get all todo requests",
		}, []string{"method"}),
		GetTodoRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_get_todo_requests",
			Help:      "The total number of get todo requests",
		}, []string{"method"}),
		CreateTodoRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_create_todo_requests",
			Help:      "The total number of create todo requests",
		}, []string{"method"}),
		UpdateTodoRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_update_todo_requests",
			Help:      "The total number of update todo requests",
		}, []string{"method"}),
		UpdateStatusTodoRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_update_status_todo_requests",
			Help:      "The total number of update status todo requests",
		}, []string{"method"}),
		DeleteTodoRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "todo_delete_todo_requests",
			Help:      "The total number of delete todo requests",
		}, []string{"method"}),
	}

	reg.MustRegister(
		m.SuccessRequests,
		m.FailRequests,
		m.GetAllTodoRequests,
		m.GetTodoRequests,
		m.CreateTodoRequests,
		m.UpdateTodoRequests,
		m.UpdateStatusTodoRequests,
		m.DeleteTodoRequests,
	)

	return m
}
