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
	GetTodoById(ctx context.Context, db *sql.DB, todoId string) (responses.GetTodoByIdResponse, error)
	UpdateTodoById(ctx context.Context, tx *sql.Tx, todo *requests.UpdateTodoRequest) error
	UpdateTodoStatusById(ctx context.Context, tx *sql.Tx, todo *requests.UpdateTodoStatusRequest) error
	DeleteTodoById(ctx context.Context, tx *sql.Tx, todoId string) error
}
