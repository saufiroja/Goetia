package todos_test

import (
	"database/sql"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestUpdateTodoByID(t *testing.T) {
	t.Run("[Positive] success update todo by ID", func(t *testing.T) {
		now := time.Now()
		mockData := &requests.UpdateTodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTx := &sql.Tx{} // Create a mock transaction
		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, mockData.TodoId).Return(nil, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoById(init.ctx, mockTx, mockData).Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)
		init.mockRedisCli.EXPECT().Del("todoById").Return(nil)

		request := &grpc.UpdateTodoRequest{
			TodoId:      mockData.TodoId,
			Title:       mockData.Title,
			Description: mockData.Description,
			Completed:   mockData.Completed,
			UpdatedAt:   mockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoById(init.ctx, request)
		assert.NoError(t, err)
	})

	t.Run("[Negative] failed start transaction", func(t *testing.T) {
		now := time.Now()
		mockData := &requests.UpdateTodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(nil, assert.AnError)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, mockData.TodoId).Return(nil, nil)

		request := &grpc.UpdateTodoRequest{
			TodoId:      mockData.TodoId,
			Title:       mockData.Title,
			Description: mockData.Description,
			Completed:   mockData.Completed,
			UpdatedAt:   mockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed get todo by ID", func(t *testing.T) {
		now := time.Now()
		mockData := &requests.UpdateTodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, mockData.TodoId).Return(nil, assert.AnError)

		request := &grpc.UpdateTodoRequest{
			TodoId:      mockData.TodoId,
			Title:       mockData.Title,
			Description: mockData.Description,
			Completed:   mockData.Completed,
			UpdatedAt:   mockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed update todo by ID", func(t *testing.T) {
		now := time.Now()
		mockData := &requests.UpdateTodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTx := &sql.Tx{} // Create a mock transaction
		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().RollbackTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, mockData.TodoId).Return(nil, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoById(init.ctx, mockTx, mockData).Return(assert.AnError)

		request := &grpc.UpdateTodoRequest{
			TodoId:      mockData.TodoId,
			Title:       mockData.Title,
			Description: mockData.Description,
			Completed:   mockData.Completed,
			UpdatedAt:   mockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete cache todos", func(t *testing.T) {
		now := time.Now()
		mockData := &requests.UpdateTodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTx := &sql.Tx{} // Create a mock transaction
		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, mockData.TodoId).Return(nil, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoById(init.ctx, mockTx, mockData).Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(assert.AnError)

		request := &grpc.UpdateTodoRequest{
			TodoId:      mockData.TodoId,
			Title:       mockData.Title,
			Description: mockData.Description,
			Completed:   mockData.Completed,
			UpdatedAt:   mockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete cache todoById", func(t *testing.T) {
		now := time.Now()
		mockData := &requests.UpdateTodoRequest{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init := setupTest(t)
		defer init.mockCtrl.Finish()

		mockTx := &sql.Tx{} // Create a mock transaction
		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(gomock.Any(), db, mockData.TodoId).Return(nil, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoById(init.ctx, mockTx, mockData).Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)
		init.mockRedisCli.EXPECT().Del("todoById").Return(assert.AnError)

		request := &grpc.UpdateTodoRequest{
			TodoId:      mockData.TodoId,
			Title:       mockData.Title,
			Description: mockData.Description,
			Completed:   mockData.Completed,
			UpdatedAt:   mockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoById(init.ctx, request)
		assert.Error(t, err)
	})
}
