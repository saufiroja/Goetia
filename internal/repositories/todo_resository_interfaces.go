package repositories

import (
	"context"
	"database/sql"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
)

type ITodoRepository interface {
	InsertTodo(ctx context.Context, tx *sql.Tx, todo *requests.TodoRequest) error
	GetAllTodos(ctx context.Context, db *sql.DB) ([]responses.GetAllTodoResponse, error)
}
