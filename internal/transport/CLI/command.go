package CLI

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Command struct {
	args []string
}

func NewCommand(args []string) *Command {
	return &Command{args: args}
}

func (c *Command) isAdd() bool {
	if len(c.args) != 3 || c.args[2] == "" || c.args[1] != "add" {
		return false
	}
	return true
}

func (c *Command) Add(h *Handler) {
	ID, err := h.service.CreateHabit(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("create habit: [%w]", err))
	}
	fmt.Printf("Была записана новая привычка: %s c айди %d\n", os.Args[2], ID)
	habit, err := h.service.GetHabit(ID)
	if err != nil {
		log.Fatal(fmt.Errorf("get habit in create process: [%w]", err))
	}
	fmt.Printf("Нынешняя запись:\n")
	fmt.Printf("Идентификатор: %d\n", habit.HabitID)
	fmt.Printf("Название: %s\n", habit.Name)
	fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
	fmt.Printf("Последнее повторение: %s\n", habit.LastRepetition)
}

func (c *Command) isList() bool {
	if len(c.args) != 2 || c.args[1] != "list" {
		return false
	}
	return true
}

func (c *Command) List(h *Handler) {
	hs, err := h.service.GetHabits()
	if err != nil {
		log.Fatal(fmt.Errorf("get habbits: [%w]", err))
	}
	for _, habit := range hs {
		fmt.Printf("Идентификатор: %d\n", habit.HabitID)
		fmt.Printf("Название: %s\n", habit.Name)
		fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
		fmt.Printf("Последнее повторение: %s\n\n", habit.LastRepetition)
	}
}

func (c *Command) isGetHabit() bool {
	if len(c.args) != 3 || c.args[1] != "id" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) GetHabit(h *Handler) {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("strconv id: [%w]", err))
	}
	habit, err := h.service.GetHabit(int64(id))
	if err != nil {
		log.Fatal(fmt.Errorf("get habit in self-process: [%w]", err))
	}
	fmt.Printf("Идентификатор: %d\n", habit.HabitID)
	fmt.Printf("Название: %s\n", habit.Name)
	fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
	fmt.Printf("Последнее повторение: %s\n\n", habit.LastRepetition)
}

func (c *Command) isDone() bool {
	if len(c.args) != 3 || c.args[1] != "done" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) Done(h *Handler) {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("strconv id: [%w]", err))
	}
	err = h.service.MarkHabitDone(int64(id))
	if err != nil {
		log.Fatal(fmt.Errorf("habit done: [%w]", err))
	}
}

func (c *Command) isDelete() bool {
	if len(c.args) != 3 || c.args[1] != "delete" || c.args[2] == "" {
		return false
	}
	return true
}

func (c *Command) Delete(h *Handler) {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("strconv id: [%w]", err))
	}

	habit, err := h.service.GetHabit(int64(id))
	if err != nil {
		log.Fatal(fmt.Errorf("get habit in delete: [%w]", err))
	}

	err = h.service.DeleteHabit(int64(id))
	if err != nil {
		log.Fatal(fmt.Errorf("delete habit: [%w]", err))
	}
	fmt.Printf("Привычка снизу была удалена!\n")
	fmt.Printf("Идентификатор: %d\n", habit.HabitID)
	fmt.Printf("Название: %s\n", habit.Name)
	fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
	fmt.Printf("Последнее повторение: %s\n\n", habit.LastRepetition)
}

func (c *Command) isHelp() bool {
	if len(os.Args) != 2 || os.Args[1] != "help" {
		return false
	}
	return true
}

func (c *Command) Help() {
	
}
