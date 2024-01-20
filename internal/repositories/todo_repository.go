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

func (r *repository) GetTodoById(ctx context.Context, db *sql.DB, todoId string) (responses.GetTodoByIdResponse, error) {
	query := `SELECT todo_id, title, description, completed, created_at, updated_at FROM todos WHERE todo_id = $1`
	row := db.QueryRowContext(ctx, query, todoId)

	var todo responses.GetTodoByIdResponse
	err := row.Scan(&todo.TodoId, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (r *repository) UpdateTodoById(ctx context.Context, tx *sql.Tx, todo *requests.UpdateTodoRequest) error {
	query := `UPDATE todos SET title = $1, description = $2, completed = $3, updated_at = $4 WHERE todo_id = $5`
	_, err := tx.ExecContext(ctx, query, todo.Title, todo.Description, todo.Completed, todo.UpdatedAt, todo.TodoId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateTodoStatusById(ctx context.Context, tx *sql.Tx, todo *requests.UpdateTodoStatusRequest) error {
	query := `UPDATE todos SET completed = $1, updated_at = $2 WHERE todo_id = $3`
	_, err := tx.ExecContext(ctx, query, todo.Completed, todo.UpdatedAt, todo.TodoId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteTodoById(ctx context.Context, tx *sql.Tx, todoId string) error {
	query := `DELETE FROM todos WHERE todo_id = $1`
	_, err := tx.ExecContext(ctx, query, todoId)
	if err != nil {
		return err
	}

	return nil
}
