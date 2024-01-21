package command

type TodoCommand struct {
	InsertTodoCommand           *InsertTodoCommand
	UpdateTodoCommand           *UpdateTodoCommand
	UpdateStatusTodoByIdCommand *UpdateStatusTodoByIdCommand
	DeleteTodoByIdCommand       *DeleteTodoByIdCommand
}

func NewTodoCommand(
	insertTodoCommand *InsertTodoCommand,
	updateTodoCommand *UpdateTodoCommand,
	updateStatusTodoByIdCommand *UpdateStatusTodoByIdCommand,
	deleteTodoByIdCommand *DeleteTodoByIdCommand,
) *TodoCommand {
	return &TodoCommand{
		InsertTodoCommand:           insertTodoCommand,
		UpdateTodoCommand:           updateTodoCommand,
		UpdateStatusTodoByIdCommand: updateStatusTodoByIdCommand,
		DeleteTodoByIdCommand:       deleteTodoByIdCommand,
	}
}
