package requests

import "time"

type TodoRequest struct {
	TodoId      string    `json:"todo_id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Completed   bool      `json:"completed" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
