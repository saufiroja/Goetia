package responses

import "time"

type GetAllTodoResponse struct {
	TodoId    string    `json:"todo_id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
