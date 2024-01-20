package query

type TodoQuery struct {
	GetAllTodoQuery IGetAllTodoQuery
}

func NewTodoQuery(getAllTodoQuery IGetAllTodoQuery) *TodoQuery {
	return &TodoQuery{
		GetAllTodoQuery: getAllTodoQuery,
	}
}
