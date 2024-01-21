package services

import (
	"context"
	"github.com/oklog/ulid/v2"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/sirupsen/logrus"
	"time"
)

type service struct {
	db             *database.Postgres
	log            *logrus.Logger
	todoRepository repositories.ITodoRepository
}

func NewService(db *database.Postgres, log *logrus.Logger, todoRepository repositories.ITodoRepository) ITodoService {
	return &service{
		db:             db,
		log:            log,
		todoRepository: todoRepository,
	}
}

func (s *service) InsertTodo(ctx context.Context, request *grpc.TodoRequest) error {
	input := &requests.TodoRequest{
		TodoId:      ulid.Make().String(),
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
		CreatedAt:   time.Unix(request.CreatedAt, 0),
		UpdatedAt:   time.Unix(request.UpdatedAt, 0),
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.Error("error starting transaction")
		return err
	}

	err = s.todoRepository.InsertTodo(ctx, tx, input)
	if err != nil {
		s.log.Error("error inserting todos")
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	return nil
}

func (s *service) GetAllTodo(ctx context.Context) ([]responses.GetAllTodoResponse, error) {
	todos, err := s.todoRepository.GetAllTodos(ctx, s.db.Open())
	if err != nil {
		s.log.Error("error getting all todos")
		return nil, err
	}

	return todos, nil
}

func (s *service) GetTodoById(ctx context.Context, todoId string) (responses.GetTodoByIdResponse, error) {
	todo, err := s.todoRepository.GetTodoById(ctx, s.db.Open(), todoId)
	if err != nil {
		s.log.Error("error getting todos by id")
		return todo, err
	}

	return todo, nil
}

func (s *service) UpdateTodoById(ctx context.Context, request *grpc.UpdateTodoRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &requests.UpdateTodoRequest{
		TodoId:      request.TodoId,
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
		UpdatedAt:   time.Unix(request.UpdatedAt, 0),
	}

	_, err := s.todoRepository.GetTodoById(ctx, s.db.Open(), input.TodoId)
	if err != nil {
		s.log.Error("error getting todos by id")
		return err
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.Error("error starting transaction")
		return err
	}

	err = s.todoRepository.UpdateTodoById(ctx, tx, input)
	if err != nil {
		s.log.Error("error updating todos by id")
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	return nil
}

func (s *service) UpdateTodoStatusById(ctx context.Context, request *grpc.UpdateTodoStatusRequest) error {
	input := &requests.UpdateTodoStatusRequest{
		TodoId:    request.TodoId,
		Completed: request.Completed,
		UpdatedAt: time.Unix(request.UpdatedAt, 0),
	}

	todo, err := s.todoRepository.GetTodoById(ctx, s.db.Open(), input.TodoId)
	if err != nil {
		s.log.Error("error getting todos by id")
		return err
	}

	if todo.Completed {
		input.Completed = false
	} else {
		input.Completed = true
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.Error("error starting transaction")
		return err
	}

	err = s.todoRepository.UpdateTodoStatusById(ctx, tx, input)
	if err != nil {
		s.log.Error("error updating todos status by id")
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	return nil
}

func (s *service) DeleteTodoById(ctx context.Context, todoId string) error {
	_, err := s.todoRepository.GetTodoById(ctx, s.db.Open(), todoId)
	if err != nil {
		s.log.Error("error getting todos by id")
		return err
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.Error("error starting transaction")
		return err
	}

	err = s.todoRepository.DeleteTodoById(ctx, tx, todoId)
	if err != nil {
		s.log.Error("error deleting todos by id")
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	return nil
}
