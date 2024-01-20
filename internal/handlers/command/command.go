package command

type TodoCommand struct {
	InsertTodoCommand IInsertTodoCommand
}

func NewTodoCommand(insertTodoCommand IInsertTodoCommand) *TodoCommand {
	return &TodoCommand{
		InsertTodoCommand: insertTodoCommand,
	}
}
