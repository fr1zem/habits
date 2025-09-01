package CLI

import (
	"CLIappHabits/internal/usecases"
	"fmt"
	"log"
	"os"
	"strconv"
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
	case command.isList():
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
	case command.isDone():
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(fmt.Errorf("strconv id: [%w]", err))
		}
		err = h.service.MarkHabitDone(int64(id))
		if err != nil {
			log.Fatal(fmt.Errorf("habit done: [%w]", err))
		}
	case command.isGetHabitByID():
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
	default:
		fmt.Println("Неизвестная команда")
	}
}
