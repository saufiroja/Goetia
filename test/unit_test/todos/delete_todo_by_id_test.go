package todos_test

import (
	"database/sql"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestDeleteTodoById(t *testing.T) {
	t.Run("[Positive] success delete todo by id", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}
		mockTx := &sql.Tx{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "DeleteTodoById").Return(testLogger.WithField("todo_service.go", "DeleteTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.DeleteTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, "1").Return(nil, nil)
		init.mockTodoRepo.EXPECT().DeleteTodoById(init.ctx, mockTx, "1").Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)
		init.mockRedisCli.EXPECT().Del("todoById").Return(nil)

		err := init.service.DeleteTodoById(init.ctx, "1")
		assert.NoError(t, err)
	})

	t.Run("[Negative] failed get todo by id", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "DeleteTodoById").Return(testLogger.WithField("todo_service.go", "DeleteTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.DeleteTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, "1").Return(nil, assert.AnError)

		err := init.service.DeleteTodoById(init.ctx, "1")
		assert.Error(t, err)
	})

	t.Run("[Negative] failed start transaction", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(nil, assert.AnError)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "DeleteTodoById").Return(testLogger.WithField("todo_service.go", "DeleteTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.DeleteTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, "1").Return(nil, nil)

		err := init.service.DeleteTodoById(init.ctx, "1")
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete todo by id", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}
		mockTx := &sql.Tx{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().RollbackTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "DeleteTodoById").Return(testLogger.WithField("todo_service.go", "DeleteTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.DeleteTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, "1").Return(nil, nil)
		init.mockTodoRepo.EXPECT().DeleteTodoById(init.ctx, mockTx, "1").Return(assert.AnError)

		err := init.service.DeleteTodoById(init.ctx, "1")
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete cache todos", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}
		mockTx := &sql.Tx{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "DeleteTodoById").Return(testLogger.WithField("todo_service.go", "DeleteTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.DeleteTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, "1").Return(nil, nil)
		init.mockTodoRepo.EXPECT().DeleteTodoById(init.ctx, mockTx, "1").Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(assert.AnError)

		err := init.service.DeleteTodoById(init.ctx, "1")
		assert.Error(t, err)
	})

	t.Run("[Negative] failed delete cache todo by id", func(t *testing.T) {
		init := setupTest(t)
		defer init.mockCtrl.Finish()

		db := &sql.DB{}
		mockTx := &sql.Tx{}

		init.mockDB.EXPECT().Db().Return(db)
		init.mockDB.EXPECT().StartTransaction().Return(mockTx, nil)
		init.mockDB.EXPECT().CommitTransaction(mockTx)

		testLogger, _ := test.NewNullLogger()
		init.mockLogger.EXPECT().StartLogger("todo_service.go", "DeleteTodoById").Return(testLogger.WithField("todo_service.go", "DeleteTodoById"))

		mockTracer := mocktracer.New()
		init.mockTracing.EXPECT().StartSpan(gomock.Any(), gomock.Any()).Return(mockTracer.StartSpan("Service.DeleteTodoById"), init.ctx)

		init.mockTodoRepo.EXPECT().GetTodoById(init.ctx, db, "1").Return(nil, nil)
		init.mockTodoRepo.EXPECT().DeleteTodoById(init.ctx, mockTx, "1").Return(nil)

		init.mockRedisCli.EXPECT().Del("todos").Return(nil)
		init.mockRedisCli.EXPECT().Del("todoById").Return(assert.AnError)

		err := init.service.DeleteTodoById(init.ctx, "1")
		assert.Error(t, err)
	})
}
