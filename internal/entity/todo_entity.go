package entity

type Todo struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed" validate:"required"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
