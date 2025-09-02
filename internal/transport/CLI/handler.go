package CLI

import (
	"CLIappHabits/internal/entities"
	"os"
)

/*
Нужно сделать интерфейс UseCase(Service), который и работать здесь уже с этим интерфейсом
Возможно ещё, что контроллер должен быть не handler, а command и presenter добавить.
Только сувать его сюда же или нет?
*/

type Service interface {
	CreateHabit(name string) (int64, error)
	GetHabit(ID int64) (entities.Habit, error)
	GetHabits() ([]entities.Habit, error)
	MarkHabitDone(ID int64) error
	DeleteHabit(ID int64) error
}

type Handler struct {
	service Service
}

/*Так же нужно сделать конструктор, который бы принимал интерфейс, а возвращал хэндлер*/
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Run() {

	command := NewCommand(os.Args, h.service)

	command.Run()

}
