package CLI

import (
	"CLIappHabits/internal/entities"
	"CLIappHabits/internal/usecases"
	"errors"
	"fmt"
)

type CLIPresenter struct {
}

func NewCLIPresenter() *CLIPresenter {
	return &CLIPresenter{}
}

func (p *CLIPresenter) FormatError(err error) {
	fmt.Println("Произошла ошибка!")
	switch {
	case errors.Is(err, entities.ErrHabitNotExists):
		fmt.Println("Данной привычки не существует")
	case errors.Is(err, entities.ErrHabitAlreadyExists):
		fmt.Println("Данная привычка уже существует")
	case errors.Is(err, entities.ErrEmptyName):
		fmt.Println("Имя привычки не указано")
	default:
		fmt.Println(err.Error())
	}
}

func (p *CLIPresenter) FormatAdd(output usecases.GetHabitOutputDTO) {
	fmt.Println("Была добавлена новая привычка!")
	fmt.Printf("Нынешняя запись:\n")
	fmt.Printf("Идентификатор: %d\n", output.HabitID)
	fmt.Printf("Название: %s\n", output.Name)
	fmt.Printf("Количество повторений: %d\n", output.Repetitions)
	fmt.Printf("Последнее повторение: ещё не было выполнено ни одного раза")
}

func (p *CLIPresenter) FormatGetHabit(output usecases.GetHabitOutputDTO) {
	fmt.Printf("Идентификатор: %d\n", output.HabitID)
	fmt.Printf("Название: %s\n", output.Name)
	if output.Repetitions == 0 {
		fmt.Printf("Количество повторений: %d\n", output.Repetitions)
	}
	if output.Repetitions != 0 {
		fmt.Printf("Количество повторений: %d\n", output.Repetitions)
		fmt.Printf("Последнее повторение: %s\n", output.LastRepetition)
	}
	fmt.Println("")
}

func (p *CLIPresenter) FormatList(output usecases.ListHabitsOutputDTO) {
	if len(output.Habits) == 0 {
		fmt.Println("В вашем списке привычек ещё нету ни одной привычки!")
	}
	for _, habit := range output.Habits {
		fmt.Printf("Идентификатор: %d\n", habit.HabitID)
		fmt.Printf("Название: %s\n", habit.Name)
		fmt.Printf("Количество повторений: %d\n", habit.Repetitions)
		if habit.Repetitions != 0 {
			fmt.Printf("Последнее повторение: %s\n", habit.LastRepetition)
		}
		fmt.Println("")
	}
}

func (p *CLIPresenter) FormatCompleted(output usecases.GetHabitOutputDTO) {
	fmt.Printf("Привычка %s была выполнена!\n", output.Name)
	fmt.Printf("Идентификатор: %d\n", output.HabitID)
	fmt.Printf("Название: %s\n", output.Name)
	fmt.Printf("Количество повторений: %d\n", output.Repetitions)
	fmt.Printf("Последнее повторение: %s\n\n", output.LastRepetition)
}

func (p *CLIPresenter) FormatDelete(output usecases.GetHabitOutputDTO) {
	fmt.Printf("Привычка была удалена!\n")
	fmt.Printf("Идентификатор: %d\n", output.HabitID)
	fmt.Printf("Название: %s\n", output.Name)
	fmt.Printf("Количество повторений: %d\n", output.Repetitions)
	if output.Repetitions != 0 {
		fmt.Printf("Последнее повторение: %s\n", output.LastRepetition)
	}
	fmt.Println()
}
