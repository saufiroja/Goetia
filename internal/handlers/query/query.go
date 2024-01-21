package query

type TodoQuery struct {
	GetAllTodoQuery  *GetAllTodoQuery
	GetTodoByIdQuery *GetTodoByIdQuery
}

func NewTodoQuery(
	getAllTodoQuery *GetAllTodoQuery,
	GetTodoByIdQuery *GetTodoByIdQuery,
) *TodoQuery {
	return &TodoQuery{
		GetAllTodoQuery:  getAllTodoQuery,
		GetTodoByIdQuery: GetTodoByIdQuery,
	}
}
