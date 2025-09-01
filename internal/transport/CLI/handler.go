package CLI

import (
	"CLIappHabits/internal/usecases"
	"fmt"
	"os"
)

type Handler struct {
	service *usecases.HabitsService
}

func NewHandler(service *usecases.HabitsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Run() {

	command := NewCommand(os.Args)

	switch {
	case command.isAdd():
		command.Add(h)
	case command.isList():
		command.List(h)
	case command.isGetHabit():
		command.GetHabit(h)
	case command.isDone():
		command.Done(h)
	case command.isDelete():
		command.Delete(h)
	default:
		fmt.Println("Неизвестная команда")
	}
}
