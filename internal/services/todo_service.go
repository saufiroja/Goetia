package services

import (
	"context"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/sirupsen/logrus"
	"time"
)

type service struct {
	db             *database.Postgres
	log            *logrus.Logger
	todoRepository repositories.TodoRepository
}

func NewService(db *database.Postgres, log *logrus.Logger, todoRepository repositories.TodoRepository) TodoService {
	return &service{
		db:             db,
		log:            log,
		todoRepository: todoRepository,
	}
}

func (s *service) InsertTodo(input *requests.TodoRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.Error("error starting transaction")
		return err
	}

	err = s.todoRepository.InsertTodo(ctx, tx, input)
	if err != nil {
		s.log.Error("error inserting todo")
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	return nil
}

func (s *service) GetAllTodo() ([]responses.GetAllTodoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todos, err := s.todoRepository.GetAllTodos(ctx, s.db.Open())
	if err != nil {
		s.log.Error("error getting all todos")
		return nil, err
	}

	return todos, nil
}
