package CLI

import (
	"CLIappHabits/internal/entities"
	"errors"
	"fmt"
)

type Presenter struct {
}

func (p *Presenter) FormatError(err error) {
	fmt.Println("Произошла ошибка!")
	switch {
	case errors.Is(err, entities.ErrHabitNotExists):
		fmt.Println("Данной привычки не существует")
	case errors.Is(err, entities.ErrHabitAlreadyExists):
		fmt.Println("Данная привычка уже существует")
	case errors.Is(err, entities.ErrEmptyName):
		fmt.Println("Имя привычки не указано")
	default:
		fmt.Println("Неизвестная ошибка")
	}
}

func (p *Presenter) FormatGetHabit(habit entities.Habit) {
	fmt.Printf("Идентификатор: %d\n", habit.HabitID)
	fmt.Printf("Название: %s\n", habit.Name)
	fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
	fmt.Printf("Последнее повторение: %s\n\n", habit.LastRepetition)
}

func (p *Presenter) FormatAdd(habit entities.Habit) {
	fmt.Printf("Нынешняя запись:\n")
	fmt.Printf("Идентификатор: %d\n", habit.HabitID)
	fmt.Printf("Название: %s\n", habit.Name)
	fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
	fmt.Printf("Последнее повторение: %s\n", habit.LastRepetition)
}

func (p *Presenter) FormatList(hs []entities.Habit) {
	if len(hs) == 0 {
		fmt.Println("В вашем списке привычек ещё нету ни одной привычки!")
	}
	for _, habit := range hs {
		fmt.Printf("Идентификатор: %d\n", habit.HabitID)
		fmt.Printf("Название: %s\n", habit.Name)
		fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
		fmt.Printf("Последнее повторение: %s\n\n", habit.LastRepetition)
	}
}

func (p *Presenter) FormatDone(h entities.Habit) {
	fmt.Printf("Привычка %s была выполнена!\n", h.Name)
	fmt.Printf("Идентификатор: %d\n", h.HabitID)
	fmt.Printf("Название: %s\n", h.Name)
	fmt.Printf("Количество повторений: %d\n", h.Repetitions)
	fmt.Printf("Последнее повторение: %s\n\n", h.LastRepetition)
}

func (p *Presenter) FormatDelete(h entities.Habit) {
	fmt.Printf("Привычка была удалена!\n")
	fmt.Printf("Идентификатор: %d\n", h.HabitID)
	fmt.Printf("Название: %s\n", h.Name)
	fmt.Printf("Количество повторений: %d\n", h.Repetitions)
	fmt.Printf("Последнее повторение: %s\n\n", h.LastRepetition)
}
