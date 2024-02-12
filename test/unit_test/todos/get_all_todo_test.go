package todos_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/redis/go-redis/v9"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestGetAllTodo(t *testing.T) {
	dataByte := []byte(`[{"todo_id":"1","title":"test","completed":false,"created_at":"2021-08-01T00:00:00Z","updated_at":"2021-08-01T00:00:00Z"}]`)

	t.Run("[Positive] success get all todo data from cache", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetAllTodo"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetAllTodo").Return(testLogger.WithField("todo_service.go", "GetAllTodo"))

		init.mockRedisCli.EXPECT().Get("todos").Return(string(dataByte), nil)

		var data []responses.GetAllTodoResponse
		err := json.Unmarshal(dataByte, &data)
		if err != nil {
			t.Error(err)
		}

		res, err := init.service.GetAllTodo(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, data, res)
	})

	t.Run("[Positive] success get all todo data from database", func(t *testing.T) {
		mockData := []responses.GetAllTodoResponse{
			{
				TodoId:    "1",
				Title:     "test",
				Completed: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetAllTodo"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetAllTodo").Return(testLogger.WithField("todo_service.go", "GetAllTodo"))

		var db sql.DB
		init.mockDB.EXPECT().Db().Return(&db)
		jsonData, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
		}

		init.mockRedisCli.EXPECT().Get("todos").Return(string(jsonData), redis.Nil)
		init.mockTodoRepo.EXPECT().GetAllTodos(gomock.Any(), &db).Return(mockData, nil)
		init.mockRedisCli.EXPECT().Set("todos", jsonData, 5*time.Minute).Return(nil)

		todo, err := init.service.GetAllTodo(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, mockData, todo)
	})

	t.Run("[Negative] failed get all todo data from database", func(t *testing.T) {
		mockData := []responses.GetAllTodoResponse{
			{
				TodoId:    "1",
				Title:     "test",
				Completed: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetAllTodo"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetAllTodo").Return(testLogger.WithField("todo_service.go", "GetAllTodo"))

		var db sql.DB
		init.mockDB.EXPECT().Db().Return(&db)
		jsonData, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
		}

		init.mockRedisCli.EXPECT().Get("todos").Return(string(jsonData), redis.Nil)
		init.mockTodoRepo.EXPECT().GetAllTodos(gomock.Any(), &db).Return(nil, assert.AnError)

		todo, err := init.service.GetAllTodo(context.Background())
		assert.Error(t, err)
		assert.Nil(t, todo)
	})

	t.Run("[Negative] failed set data todo to cache", func(t *testing.T) {
		mockData := []responses.GetAllTodoResponse{
			{
				TodoId:    "1",
				Title:     "test",
				Completed: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetAllTodo"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetAllTodo").Return(testLogger.WithField("todo_service.go", "GetAllTodo"))

		var db sql.DB
		init.mockDB.EXPECT().Db().Return(&db)
		jsonData, err := json.Marshal(mockData)
		if err != nil {
			t.Error(err)
		}

		init.mockRedisCli.EXPECT().Get("todos").Return(string(jsonData), redis.Nil)
		init.mockTodoRepo.EXPECT().GetAllTodos(gomock.Any(), &db).Return(mockData, nil)
		init.mockRedisCli.EXPECT().Set("todos", jsonData, 5*time.Minute).Return(assert.AnError)

		todo, err := init.service.GetAllTodo(context.Background())
		assert.Error(t, err)
		assert.Nil(t, todo)
	})
}
