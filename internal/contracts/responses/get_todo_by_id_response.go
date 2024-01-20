package responses

import "time"

type GetTodoByIdResponse struct {
	TodoId      string    `json:"todo_id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
