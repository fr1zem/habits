package CLI

import (
	"fmt"
	"os"
	"strconv"
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
	ID, err := c.service.CreateHabit(os.Args[2])
	if err != nil {
		c.presenter.FormatError(err)
		return
	}
	fmt.Printf("Была записана новая привычка: %s c айди %d\n", os.Args[2], ID)

	habit, err := c.service.GetHabit(ID)
	if err != nil {
		c.presenter.FormatError(err)
		return
	}

	c.presenter.FormatAdd(habit)
}

func (c *Command) isList() bool {
	if len(c.args) != 2 || c.args[1] != "list" {
		return false
	}
	return true
}

func (c *Command) List() {
	hs, err := c.service.GetHabits()
	if err != nil {
		c.presenter.FormatError(err)
		return
	}
	c.presenter.FormatList(hs)
}

func (c *Command) isGetHabit() bool {
	if len(c.args) != 3 || c.args[1] != "id" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) GetHabit() {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		c.presenter.FormatError(err)
	}
	habit, err := c.service.GetHabit(int64(id))
	if err != nil {
		c.presenter.FormatError(err)
		return
	}
	c.presenter.FormatGetHabit(habit)
}

func (c *Command) isDone() bool {
	if len(c.args) != 3 || c.args[1] != "done" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) Done() {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		c.presenter.FormatError(err)
		return
	}
	err = c.service.MarkHabitDone(int64(id))
	if err != nil {
		c.presenter.FormatError(err)
		return
	}
	h, err := c.service.GetHabit(int64(id))
	if err != nil {
		c.presenter.FormatError(err)
		return
	}
	c.presenter.FormatDone(h)
}

func (c *Command) isDelete() bool {
	if len(c.args) != 3 || c.args[1] != "delete" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) Delete() {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		c.presenter.FormatError(err)
		return
	}

	habit, err := c.service.GetHabit(int64(id))
	if err != nil {
		c.presenter.FormatError(err)
		return
	}

	err = c.service.DeleteHabit(int64(id))
	if err != nil {
		c.presenter.FormatError(err)
		return
	}
	c.presenter.FormatDelete(habit)
}

func (c *Command) isHelp() bool {
	if len(os.Args) != 2 || os.Args[1] != "help" {
		return false
	}
	return true
}

func (c *Command) Help() {

}
