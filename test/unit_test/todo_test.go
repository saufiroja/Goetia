package services

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/mocks"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

type ServiceTest struct {
	mockCtrl     *gomock.Controller
	mockDB       *mocks.MockIPostgres
	mockLogger   *mocks.MockILogger
	mockTodoRepo *mocks.MockITodoRepository
	mockRedisCli *mocks.MockIRedis
	mockTracing  *mocks.MockITracing
	service      services.ITodoService
	ctx          context.Context
}

func setupTest(t *testing.T) *ServiceTest {
	mockCtrl := gomock.NewController(t)

	mockDB := mocks.NewMockIPostgres(mockCtrl)
	mockLogger := mocks.NewMockILogger(mockCtrl)
	mockTodoRepo := mocks.NewMockITodoRepository(mockCtrl)
	mockRedisCli := mocks.NewMockIRedis(mockCtrl)
	mockTracing := mocks.NewMockITracing(mockCtrl)

	ctx := context.Background()
	service := services.NewService(mockDB, mockLogger, mockTodoRepo, mockRedisCli, mockTracing)

	return &ServiceTest{
		mockCtrl:     mockCtrl,
		mockDB:       mockDB,
		mockLogger:   mockLogger,
		mockTodoRepo: mockTodoRepo,
		mockRedisCli: mockRedisCli,
		mockTracing:  mockTracing,
		service:      service,
		ctx:          ctx,
	}
}

func TestInsertTodo(t *testing.T) {
	t.Run("[Positive] success insert todo", func(t *testing.T) {
		request := &grpc.TodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTx := &sql.Tx{} // Create a mock transaction
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "InsertTodo").Return(testLogger.WithField("todo_service.go", "InsertTodo"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("test"), context.Background())

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)

		init.mockTodoRepo.EXPECT().InsertTodo(gomock.Any(), mockTx, gomock.Any()).Return(nil)

		err := init.service.InsertTodo(init.ctx, request)
		assert.NoError(t, err)
	})

	t.Run("[Negative] failed start transaction", func(t *testing.T) {
		request := &grpc.TodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTracer := mocktracer.New()
		testLogger, _ := test.NewNullLogger()

		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("test"), context.Background())
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "InsertTodo").Return(testLogger.WithField("todo_service.go", "InsertTodo"))
		init.mockDB.EXPECT().StartTransaction().Return(nil, assert.AnError)

		err := init.service.InsertTodo(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed insert todo", func(t *testing.T) {
		request := &grpc.TodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTx := &sql.Tx{} // Create a mock transaction
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().RollbackTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "InsertTodo").Return(testLogger.WithField("todo_service.go", "InsertTodo"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("test"), context.Background())

		init.mockTodoRepo.EXPECT().InsertTodo(gomock.Any(), mockTx, gomock.Any()).Return(assert.AnError)

		err := init.service.InsertTodo(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete cache", func(t *testing.T) {
		request := &grpc.TodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTx := &sql.Tx{} // Create a mock transaction
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "InsertTodo").Return(testLogger.WithField("todo_service.go", "InsertTodo"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("test"), context.Background())

		init.mockRedisCli.EXPECT().Del("todos").Return(assert.AnError)

		init.mockTodoRepo.EXPECT().InsertTodo(gomock.Any(), mockTx, gomock.Any()).Return(nil)

		err := init.service.InsertTodo(init.ctx, request)
		assert.Error(t, err)
	})
}
