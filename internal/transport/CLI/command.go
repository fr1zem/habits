package CLI

type Command struct {
	args []string
}

func NewCommand(args []string) *Command {
	return &Command{args: args}
}

func (c *Command) isAdd() bool {
	if len(c.args) != 3 || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) isList() bool {
	if len(c.args) != 3 || c.args[1] != "list" {
		return false
	}
	return true
}

func (c *Command) isGetHabitByID() bool {
	if len(c.args) != 3 || c.args[1] != "id" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) isDone() bool {
	if len(c.args) != 2 || c.args[1] != "done" {
		return false
	}
	return true
}
