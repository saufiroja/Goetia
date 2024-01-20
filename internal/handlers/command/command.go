package command

type TodoCommand struct {
	InsertTodoCommand           IInsertTodoCommand
	UpdateTodoCommand           IUpdateTodoCommand
	UpdateStatusTodoByIdCommand IUpdateStatusTodoByIdCommand
	DeleteTodoByIdCommand       IDeleteTodoByIdCommand
}

func NewTodoCommand(
	insertTodoCommand IInsertTodoCommand,
	updateTodoCommand IUpdateTodoCommand,
	updateStatusTodoByIdCommand IUpdateStatusTodoByIdCommand,
	deleteTodoByIdCommand IDeleteTodoByIdCommand,
) *TodoCommand {
	return &TodoCommand{
		InsertTodoCommand:           insertTodoCommand,
		UpdateTodoCommand:           updateTodoCommand,
		UpdateStatusTodoByIdCommand: updateStatusTodoByIdCommand,
		DeleteTodoByIdCommand:       deleteTodoByIdCommand,
	}
}
