package CLI

import (
	"fmt"
	"os"
)

type Command struct {
	args      []string
	service   Service
	presenter Presenter
}

func NewCommand(args []string, service Service) *Command {
	return &Command{
		args:    args,
		service: service,
	}
}

func (c *Command) Run() {

	switch {
	case c.isAdd():
		c.Add()
	case c.isList():
		c.List()
	case c.isGetHabit():
		c.GetHabit()
	case c.isDone():
		c.Done()
	case c.isDelete():
		c.Delete()
	default:
		fmt.Println("Неизвестная команда")
	}

}

func (c *Command) isAdd() bool {
	if len(c.args) != 3 || c.args[2] == "" || c.args[1] != "add" {
		return false
	}
	return true
}

func (c *Command) Add() {

}

func (c *Command) isList() bool {
	if len(c.args) != 2 || c.args[1] != "list" {
		return false
	}
	return true
}

func (c *Command) List() {

}

func (c *Command) isGetHabit() bool {
	if len(c.args) != 3 || c.args[1] != "id" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) GetHabit() {

}

func (c *Command) isDone() bool {
	if len(c.args) != 3 || c.args[1] != "done" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) Done() {

}

func (c *Command) isDelete() bool {
	if len(c.args) != 3 || c.args[1] != "delete" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) Delete() {

}

func (c *Command) isHelp() bool {
	if len(os.Args) != 2 || os.Args[1] != "help" {
		return false
	}
	return true
}

func (c *Command) Help() {

}
