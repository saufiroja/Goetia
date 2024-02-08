package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/redis/go-redis/v9"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/logger"
	redisCli "github.com/saufiroja/cqrs/pkg/redis"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"time"
)

type service struct {
	db             *database.Postgres
	log            *logger.Logger
	todoRepository repositories.ITodoRepository
	redisCli       *redisCli.Redis
	tracing        *tracing.Tracing
}

func NewService(
	db *database.Postgres,
	log *logger.Logger,
	todoRepository repositories.ITodoRepository,
	redisCli *redisCli.Redis,
	tracing *tracing.Tracing,
) ITodoService {
	return &service{
		db:             db,
		log:            log,
		todoRepository: todoRepository,
		redisCli:       redisCli,
		tracing:        tracing,
	}
}

func (s *service) InsertTodo(ctx context.Context, request *grpc.TodoRequest) error {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.InsertTodo")
	defer tracer.Finish()

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
		s.log.StartLogger("todo_service.go", "InsertTodo").Error("error starting transaction")
		return err
	}

	err = s.todoRepository.InsertTodo(ctx, tx, input)
	if err != nil {
		s.log.StartLogger("todo_service.go", "InsertTodo").Error("error inserting todos")
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	// delete data from redis
	err = s.redisCli.Del(redisCli.TodosKey)
	if err != nil {
		s.log.StartLogger("todo_service.go", "InsertTodo").Error("error deleting todos")
		return err
	}

	s.log.StartLogger("todo_service.go", "InsertTodo").Info("success inserting todos")

	return nil
}

func (s *service) GetAllTodo(ctx context.Context) ([]responses.GetAllTodoResponse, error) {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.GetAllTodo")
	defer tracer.Finish()

	// get data from redis
	data, err := s.redisCli.Get(redisCli.TodosKey)
	if errors.Is(err, redis.Nil) {
		todos, err := s.todoRepository.GetAllTodos(ctx, s.db.DB)
		if err != nil {
			s.log.StartLogger("todo_service.go", "GetAllTodo").Error("error getting all todos")
			return nil, err
		}

		// marshal todos
		jsonData, err := json.Marshal(todos)
		if err != nil {
			s.log.StartLogger("todo_service.go", "GetAllTodo").Error("error marshal todos")
			return nil, err
		}

		// set data to redis
		err = s.redisCli.Set(redisCli.TodosKey, jsonData, 5*time.Minute)
		if err != nil {
			errMsg := fmt.Sprintf("error setting todos to redis: %v", err)
			s.log.StartLogger("todo_service.go", "GetAllTodo").Error(errMsg)
			return nil, err
		}

		s.log.StartLogger("todo_service.go", "GetAllTodo").Info("success getting all todos")

		return todos, nil
	}

	var todos []responses.GetAllTodoResponse
	err = json.Unmarshal([]byte(data), &todos)
	if err != nil {
		s.log.StartLogger("todo_service.go", "GetAllTodo").Error("error unmarshal todos")
		return nil, err
	}

	s.log.StartLogger("todo_service.go", "GetAllTodo").Info("success getting all todos")

	return todos, nil
}

func (s *service) GetTodoById(ctx context.Context, todoId string) (responses.GetTodoByIdResponse, error) {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.GetTodoById")
	defer tracer.Finish()

	data, err := s.redisCli.Get(redisCli.TodoByIdKey)
	if errors.Is(err, redis.Nil) {
		todo, err := s.todoRepository.GetTodoById(ctx, s.db.DB, todoId)
		if err != nil {
			errMsg := fmt.Sprintf("error getting todos by id: %s", todoId)
			s.log.StartLogger("todo_service.go", "GetTodoById").Error(errMsg)
			return todo, err
		}

		// marshal todos
		jsonData, err := json.Marshal(todo)
		if err != nil {
			s.log.StartLogger("todo_service.go", "GetTodoById").Error("error marshal todos")
			return todo, err
		}

		// set data to redis
		err = s.redisCli.Set(redisCli.TodoByIdKey, jsonData, 5*time.Minute)
		if err != nil {
			errMsg := fmt.Sprintf("error setting todos to redis: %v", err)
			s.log.StartLogger("todo_service.go", "GetTodoById").Error(errMsg)
			return todo, err
		}

		res := fmt.Sprintf("success getting todos by id: %s", todoId)
		s.log.StartLogger("todo_service.go", "GetTodoById").Info(res)

		return todo, nil
	}

	var todo responses.GetTodoByIdResponse
	err = json.Unmarshal([]byte(data), &todo)
	if err != nil {
		s.log.StartLogger("todo_service.go", "GetTodoById").Error("error unmarshal todos")
		return todo, err
	}

	res := fmt.Sprintf("success getting todos by id: %s", todoId)
	s.log.StartLogger("todo_service.go", "GetTodoById").Info(res)

	return todo, nil
}

func (s *service) UpdateTodoById(ctx context.Context, request *grpc.UpdateTodoRequest) error {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.UpdateTodoById")
	defer tracer.Finish()

	input := &requests.UpdateTodoRequest{
		TodoId:      request.TodoId,
		Title:       request.Title,
		Description: request.Description,
		Completed:   request.Completed,
		UpdatedAt:   time.Unix(request.UpdatedAt, 0),
	}

	_, err := s.todoRepository.GetTodoById(ctx, s.db.DB, input.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("error getting todos by id: %s", input.TodoId)
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error(errMsg)
		return err
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error("error starting transaction")
		return err
	}

	err = s.todoRepository.UpdateTodoById(ctx, tx, input)
	if err != nil {
		errMsg := fmt.Sprintf("error updating todos by id: %s", input.TodoId)
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error(errMsg)
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	// delete data from redis
	err = s.redisCli.Del(redisCli.TodosKey)
	if err != nil {
		errMsg := fmt.Sprintf("error deleting key: %s", redisCli.TodosKey)
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error(errMsg)
		return err
	}

	err = s.redisCli.Del(redisCli.TodoByIdKey)
	if err != nil {
		errMsg := fmt.Sprintf("error deleting key: %s", redisCli.TodoByIdKey)
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error(errMsg)
		return err
	}

	res := fmt.Sprintf("success updating todos by id: %s", input.TodoId)
	s.log.StartLogger("todo_service.go", "UpdateTodoById").Info(res)

	return nil
}

func (s *service) UpdateTodoStatusById(ctx context.Context, request *grpc.UpdateTodoStatusRequest) error {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.UpdateTodoStatusById")
	defer tracer.Finish()

	input := &requests.UpdateTodoStatusRequest{
		TodoId:    request.TodoId,
		Completed: request.Completed,
		UpdatedAt: time.Unix(request.UpdatedAt, 0),
	}

	todo, err := s.todoRepository.GetTodoById(ctx, s.db.DB, input.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("error getting todos by id: %s", input.TodoId)
		s.log.StartLogger("todo_service.go", "UpdateTodoStatusById").Error(errMsg)
		return err
	}

	if todo.Completed {
		input.Completed = false
	} else {
		input.Completed = true
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.StartLogger("todo_service.go", "UpdateTodoStatusById").Error("error starting transaction")
		return err
	}

	err = s.todoRepository.UpdateTodoStatusById(ctx, tx, input)
	if err != nil {
		errMsg := fmt.Sprintf("error updating todos status by id: %s", input.TodoId)
		s.log.StartLogger("todo_service.go", "UpdateTodoStatusById").Error(errMsg)
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	// delete data from redis
	err = s.redisCli.Del(redisCli.TodosKey)
	if err != nil {
		errMsg := fmt.Sprintf("error deleting key: %s", redisCli.TodosKey)
		s.log.StartLogger("todo_service.go", "UpdateTodoStatusById").Error(errMsg)
		return err
	}

	err = s.redisCli.Del(redisCli.TodoByIdKey)
	if err != nil {
		errMsg := fmt.Sprintf("error deleting key: %s", redisCli.TodoByIdKey)
		s.log.StartLogger("todo_service.go", "UpdateTodoStatusById").Error(errMsg)
		return err
	}

	res := fmt.Sprintf("success updating todos status by id: %s", input.TodoId)
	s.log.StartLogger("todo_service.go", "UpdateTodoStatusById").Info(res)

	return nil
}

func (s *service) DeleteTodoById(ctx context.Context, todoId string) error {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.DeleteTodoById")
	defer tracer.Finish()

	_, err := s.todoRepository.GetTodoById(ctx, s.db.DB, todoId)
	if err != nil {
		errMsg := fmt.Sprintf("error getting todos by id: %s", todoId)
		s.log.StartLogger("todo_service.go", "DeleteTodoById").Error(errMsg)
		return err
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.StartLogger("todo_service.go", "DeleteTodoById").Error("error starting transaction")
		return err
	}

	err = s.todoRepository.DeleteTodoById(ctx, tx, todoId)
	if err != nil {
		errMsg := fmt.Sprintf("error deleting todos by id: %s", todoId)
		s.log.StartLogger("todo_service.go", "DeleteTodoById").Error(errMsg)
		s.db.RollbackTransaction(tx)
		return err
	}

	s.db.CommitTransaction(tx)

	// delete data from redis
	err = s.redisCli.Del(redisCli.TodosKey)
	if err != nil {
		errMsg := fmt.Sprintf("error deleting key: %s", redisCli.TodosKey)
		s.log.StartLogger("todo_service.go", "DeleteTodoById").Error(errMsg)
		return err
	}

	err = s.redisCli.Del(redisCli.TodoByIdKey)
	if err != nil {
		errMsg := fmt.Sprintf("error deleting key: %s", redisCli.TodoByIdKey)
		s.log.StartLogger("todo_service.go", "DeleteTodoById").Error(errMsg)
		return err
	}

	s.log.StartLogger("todo_service.go", "DeleteTodoById").Info(fmt.Sprintf("success deleting todos by id: %s", todoId))

	return nil
}
