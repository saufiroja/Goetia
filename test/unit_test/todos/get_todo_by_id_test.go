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

func TestGetTodoById(t *testing.T) {
	dataByte := []byte(`{"todo_id":"1","title":"test","completed":false,"created_at":"2021-08-01T00:00:00Z","updated_at":"2021-08-01T00:00:00Z"}`)
	t.Run("[Positive] success get todo by id from cache", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetTodoById"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetTodoById").Return(testLogger.WithField("todo_service.go", "GetTodoById"))

		init.mockRedisCli.EXPECT().Get("todoById").Return(string(dataByte), nil)

		todo, err := init.service.GetTodoById(context.Background(), "1")
		assert.NoError(t, err)
		assert.NotNil(t, todo)
	})

	t.Run("[Positive] success get todo by id from database", func(t *testing.T) {
		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Completed:   false,
			Description: "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetTodoById"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetTodoById").Return(testLogger.WithField("todo_service.go", "GetTodoById"))

		var db *sql.DB
		init.mockDB.EXPECT().Db().Return(db)

		dataBytes, _ := json.Marshal(mockData)

		init.mockRedisCli.EXPECT().Get("todoById").Return(string(dataByte), redis.Nil)
		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, "1").Return(mockData, nil)
		init.mockRedisCli.EXPECT().Set("todoById", dataBytes, 5*time.Minute).Return(nil)

		todo, err := init.service.GetTodoById(context.Background(), "1")
		assert.NoError(t, err)
		assert.NotNil(t, todo)
	})

	t.Run("[Negative] failed get todo by id from database", func(t *testing.T) {
		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Completed:   false,
			Description: "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetTodoById"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetTodoById").Return(testLogger.WithField("todo_service.go", "GetTodoById"))

		var db *sql.DB
		init.mockDB.EXPECT().Db().Return(db)

		dataBytes, _ := json.Marshal(mockData)

		init.mockRedisCli.EXPECT().Get("todoById").Return(string(dataBytes), redis.Nil)
		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, "0").Return(nil, assert.AnError)

		todo, err := init.service.GetTodoById(context.Background(), "0")
		assert.Error(t, err)
		assert.Nil(t, todo)
	})

	t.Run("[Negative] failed set cache data todo by id", func(t *testing.T) {
		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Completed:   false,
			Description: "test",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.GetTodoById"), context.Background())

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "GetTodoById").Return(testLogger.WithField("todo_service.go", "GetTodoById"))

		var db *sql.DB
		init.mockDB.EXPECT().Db().Return(db)

		dataBytes, _ := json.Marshal(mockData)

		init.mockRedisCli.EXPECT().Get("todoById").Return(string(dataBytes), redis.Nil)
		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, "1").Return(mockData, nil)
		init.mockRedisCli.EXPECT().Set("todoById", dataBytes, 5*time.Minute).Return(assert.AnError)

		todo, err := init.service.GetTodoById(context.Background(), "1")
		assert.Error(t, err)
		assert.Nil(t, todo)
	})
}
