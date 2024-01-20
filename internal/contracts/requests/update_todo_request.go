package requests

import "time"

type UpdateTodoRequest struct {
	TodoId      string    `json:"todo_id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Completed   bool      `json:"completed"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTodoStatusRequest struct {
	TodoId    string    `json:"todo_id"`
	Completed bool      `json:"completed"`
	UpdatedAt time.Time `json:"updated_at"`
}
