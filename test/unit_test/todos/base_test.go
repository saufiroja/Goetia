package todos_test

import (
	"context"
	"github.com/saufiroja/cqrs/internal/services"
	"github.com/saufiroja/cqrs/mocks"
	"go.uber.org/mock/gomock"
	"testing"
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
