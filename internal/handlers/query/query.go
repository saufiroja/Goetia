package query

type TodoQuery struct {
	GetAllTodoQuery  IGetAllTodoQuery
	GetTodoByIdQuery IGetTodoByIdQuery
}

func NewTodoQuery(
	getAllTodoQuery IGetAllTodoQuery,
	GetTodoByIdQuery IGetTodoByIdQuery,
) *TodoQuery {
	return &TodoQuery{
		GetAllTodoQuery:  getAllTodoQuery,
		GetTodoByIdQuery: GetTodoByIdQuery,
	}
}
