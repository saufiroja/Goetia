package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/repositories"
	"github.com/saufiroja/cqrs/pkg/database"
	"github.com/saufiroja/cqrs/pkg/logger"
	redisCli "github.com/saufiroja/cqrs/pkg/redis"
	"github.com/saufiroja/cqrs/pkg/tracing"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	db             database.IPostgres
	log            logger.ILogger
	todoRepository repositories.ITodoRepository
	redisCli       redisCli.IRedis
	tracing        tracing.ITracing
}

func NewService(
	db database.IPostgres,
	log logger.ILogger,
	todoRepository repositories.ITodoRepository,
	redisCli redisCli.IRedis,
	tracing tracing.ITracing,
) ITodoService {
	return &service{
		db:             db,
		log:            log,
		todoRepository: todoRepository,
		redisCli:       redisCli,
		tracing:        tracing,
	}
}

func (s *service) InsertTodo(ctx context.Context, request *requests.TodoRequest) error {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.InsertTodo")
	defer tracer.Finish()

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.StartLogger("todo_service.go", "InsertTodo").Error("error starting transaction")
		return err
	}

	err = s.todoRepository.InsertTodo(ctx, tx, request)
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
		todos, err := s.todoRepository.GetAllTodos(ctx, s.db.Db())
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

func (s *service) GetTodoById(ctx context.Context, todoId string) (*responses.GetTodoByIdResponse, error) {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.GetTodoById")
	defer tracer.Finish()

	data, err := s.redisCli.Get(redisCli.TodoByIdKey)
	if errors.Is(err, redis.Nil) {
		todo, err := s.todoRepository.GetTodoById(ctx, s.db.Db(), todoId)
		if err != nil {
			errMsg := fmt.Sprintf("error getting todos by id: %s", todoId)
			s.log.StartLogger("todo_service.go", "GetTodoById").Error(errMsg)
			return nil, status.Error(codes.NotFound, err.Error())
		}

		// marshal todos
		jsonData, err := json.Marshal(todo)
		if err != nil {
			s.log.StartLogger("todo_service.go", "GetTodoById").Error("error marshal todos")
			return nil, err
		}

		// set data to redis
		err = s.redisCli.Set(redisCli.TodoByIdKey, jsonData, 5*time.Minute)
		if err != nil {
			errMsg := fmt.Sprintf("error setting todos to redis: %v", err)
			s.log.StartLogger("todo_service.go", "GetTodoById").Error(errMsg)
			return nil, err
		}

		res := fmt.Sprintf("success getting todos by id: %s", todoId)
		s.log.StartLogger("todo_service.go", "GetTodoById").Info(res)

		return todo, nil
	}

	var todo responses.GetTodoByIdResponse
	err = json.Unmarshal([]byte(data), &todo)
	if err != nil {
		s.log.StartLogger("todo_service.go", "GetTodoById").Error("error unmarshal todos")
		return nil, err
	}

	res := fmt.Sprintf("success getting todos by id: %s", todoId)
	s.log.StartLogger("todo_service.go", "GetTodoById").Info(res)

	return &todo, nil
}

func (s *service) UpdateTodoById(ctx context.Context, request *requests.UpdateTodoRequest) error {
	tracer, ctx := s.tracing.StartSpan(ctx, "Service.UpdateTodoById")
	defer tracer.Finish()

	_, err := s.todoRepository.GetTodoById(ctx, s.db.Db(), request.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("error getting todos by id: %s", request.TodoId)
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error(errMsg)
		return status.Error(codes.NotFound, err.Error())
	}

	tx, err := s.db.StartTransaction()
	if err != nil {
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error("error starting transaction")
		return err
	}

	err = s.todoRepository.UpdateTodoById(ctx, tx, request)
	if err != nil {
		errMsg := fmt.Sprintf("error updating todos by id: %s", request.TodoId)
		s.log.StartLogger("todo_service.go", "UpdateTodoById").Error(errMsg)
		s.db.RollbackTransaction(tx)
		return status.Error(codes.Internal, err.Error())
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

	res := fmt.Sprintf("success updating todos by id: %s", request.TodoId)
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

	todo, err := s.todoRepository.GetTodoById(ctx, s.db.Db(), input.TodoId)
	if err != nil {
		errMsg := fmt.Sprintf("error getting todos by id: %s", input.TodoId)
		s.log.StartLogger("todo_service.go", "UpdateTodoStatusById").Error(errMsg)
		return status.Error(codes.NotFound, err.Error())
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
		return status.Error(codes.Internal, err.Error())
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

	_, err := s.todoRepository.GetTodoById(ctx, s.db.Db(), todoId)
	if err != nil {
		errMsg := fmt.Sprintf("error getting todos by id: %s", todoId)
		s.log.StartLogger("todo_service.go", "DeleteTodoById").Error(errMsg)
		return status.Error(codes.NotFound, err.Error())
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
		return status.Error(codes.Internal, err.Error())
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
