package todos_test

import (
	"database/sql"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
	"github.com/saufiroja/cqrs/internal/grpc"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestUpdateTodoStatus(t *testing.T) {
	t.Run("[Positive] success update todo status", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		now := time.Now()
		updateMockData := &requests.UpdateTodoStatusRequest{
			TodoId:    "1",
			Completed: true,
			UpdatedAt: time.Unix(now.Unix(), 0),
		}

		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Unix(now.Unix(), 0),
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		tx := &sql.Tx{} // Create a mock transaction
		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(tx, nil)
		init.mockDB.EXPECT().CommitTransaction(tx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoStatusById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, updateMockData.TodoId).Return(mockData, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoStatusById(init.ctx, tx, updateMockData).Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)
		init.mockRedisCli.EXPECT().Del("todoById").Return(nil)

		request := &grpc.UpdateTodoStatusRequest{
			TodoId:    updateMockData.TodoId,
			Completed: updateMockData.Completed,
			UpdatedAt: updateMockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoStatusById(init.ctx, request)
		assert.NoError(t, err)
	})

	t.Run("[Positive] success update todo status from true to false", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		now := time.Now()
		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   true,
			CreatedAt:   time.Unix(now.Unix(), 0),
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		updateMockData := &requests.UpdateTodoStatusRequest{
			TodoId:    "1",
			Completed: false,
			UpdatedAt: time.Unix(now.Unix(), 0),
		}

		tx := &sql.Tx{} // Create a mock transaction
		db := &sql.DB{}
		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(tx, nil)
		init.mockDB.EXPECT().CommitTransaction(tx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoStatusById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, updateMockData.TodoId).Return(mockData, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoStatusById(init.ctx, tx, updateMockData).Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)
		init.mockRedisCli.EXPECT().Del("todoById").Return(nil)

		request := &grpc.UpdateTodoStatusRequest{
			TodoId:    updateMockData.TodoId,
			Completed: updateMockData.Completed,
			UpdatedAt: updateMockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoStatusById(init.ctx, request)
		assert.NoError(t, err)
	})

	t.Run("[Negative] failed get todo by ID", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		now := time.Now()
		updateMockData := &requests.UpdateTodoStatusRequest{
			TodoId:    "1",
			Completed: true,
			UpdatedAt: time.Unix(now.Unix(), 0),
		}

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoStatusById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		db := &sql.DB{}
		init.mockDB.EXPECT().Db().Return(db)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, updateMockData.TodoId).Return(nil, assert.AnError)

		request := &grpc.UpdateTodoStatusRequest{
			TodoId:    updateMockData.TodoId,
			Completed: updateMockData.Completed,
			UpdatedAt: updateMockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoStatusById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed start transaction", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		now := time.Now()
		updateMockData := &requests.UpdateTodoStatusRequest{
			TodoId:    "1",
			Completed: true,
			UpdatedAt: time.Unix(now.Unix(), 0),
		}

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoStatusById").Return(testLogger.WithField("todo_service.go", "UpdateTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoById"), init.ctx)

		db := &sql.DB{}
		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(nil, assert.AnError)

		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Unix(now.Unix(), 0),
		}

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, updateMockData.TodoId).Return(mockData, nil)

		request := &grpc.UpdateTodoStatusRequest{
			TodoId:    updateMockData.TodoId,
			Completed: updateMockData.Completed,
			UpdatedAt: updateMockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoStatusById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed update todo status by ID", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		now := time.Now()
		updateMockData := &requests.UpdateTodoStatusRequest{
			TodoId:    "1",
			Completed: true,
			UpdatedAt: time.Unix(now.Unix(), 0),
		}

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoStatusById").Return(testLogger.WithField("todo_service.go", "UpdateTodoStatusById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoStatusById"), init.ctx)

		db := &sql.DB{}
		tx := &sql.Tx{}
		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(tx, nil)

		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Unix(now.Unix(), 0),
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, updateMockData.TodoId).Return(mockData, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoStatusById(init.ctx, tx, updateMockData).Return(assert.AnError)

		init.mockDB.EXPECT().RollbackTransaction(tx)

		request := &grpc.UpdateTodoStatusRequest{
			TodoId:    updateMockData.TodoId,
			Completed: updateMockData.Completed,
			UpdatedAt: updateMockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoStatusById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete cache todos", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		now := time.Now()
		updateMockData := &requests.UpdateTodoStatusRequest{
			TodoId:    "1",
			Completed: true,
			UpdatedAt: time.Unix(now.Unix(), 0),
		}

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoStatusById").Return(testLogger.WithField("todo_service.go", "UpdateTodoStatusById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoStatusById"), init.ctx)

		db := &sql.DB{}
		tx := &sql.Tx{}
		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(tx, nil)
		init.mockDB.EXPECT().CommitTransaction(tx)

		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Unix(now.Unix(), 0),
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, updateMockData.TodoId).Return(mockData, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoStatusById(init.ctx, tx, updateMockData).Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(assert.AnError)

		request := &grpc.UpdateTodoStatusRequest{
			TodoId:    updateMockData.TodoId,
			Completed: updateMockData.Completed,
			UpdatedAt: updateMockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoStatusById(init.ctx, request)
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete cache todo by ID", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		now := time.Now()
		updateMockData := &requests.UpdateTodoStatusRequest{
			TodoId:    "1",
			Completed: true,
			UpdatedAt: time.Unix(now.Unix(), 0),
		}

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "UpdateTodoStatusById").Return(testLogger.WithField("todo_service.go", "UpdateTodoStatusById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.UpdateTodoStatusById"), init.ctx)

		db := &sql.DB{}
		tx := &sql.Tx{}
		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(tx, nil)
		init.mockDB.EXPECT().CommitTransaction(tx)

		mockData := &responses.GetTodoByIdResponse{
			TodoId:      "1",
			Title:       "test",
			Description: "test",
			Completed:   false,
			CreatedAt:   time.Unix(now.Unix(), 0),
			UpdatedAt:   time.Unix(now.Unix(), 0),
		}

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, updateMockData.TodoId).Return(mockData, nil)
		init.mockTodoRepo.EXPECT().UpdateTodoStatusById(init.ctx, tx, updateMockData).Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)
		init.mockRedisCli.EXPECT().Del("todoById").Return(assert.AnError)

		request := &grpc.UpdateTodoStatusRequest{
			TodoId:    updateMockData.TodoId,
			Completed: updateMockData.Completed,
			UpdatedAt: updateMockData.UpdatedAt.Unix(),
		}

		err := init.service.UpdateTodoStatusById(init.ctx, request)
		assert.Error(t, err)
	})
}
