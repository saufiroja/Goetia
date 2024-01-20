package repositories

import (
	"context"
	"database/sql"
	"github.com/saufiroja/cqrs/internal/contracts/requests"
	"github.com/saufiroja/cqrs/internal/contracts/responses"
)

type repository struct {
}

func NewRepository() ITodoRepository {
	return &repository{}
}

func (r *repository) InsertTodo(ctx context.Context, tx *sql.Tx, todo *requests.TodoRequest) error {
	query := `INSERT INTO todos (todo_id, title, description, completed) VALUES ($1, $2, $3, $4)`
	_, err := tx.ExecContext(ctx, query, todo.TodoId, todo.Title, todo.Description, todo.Completed)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllTodos(ctx context.Context, db *sql.DB) ([]responses.GetAllTodoResponse, error) {
	query := `SELECT todo_id, title, completed, created_at, updated_at FROM todos`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var todos []responses.GetAllTodoResponse
	for rows.Next() {
		var todo responses.GetAllTodoResponse
		err = rows.Scan(&todo.TodoId, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
