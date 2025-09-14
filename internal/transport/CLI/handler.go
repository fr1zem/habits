package CLI

import (
	"CLIappHabits/internal/entities"
	"os"
	"strconv"
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

type Router interface {
	Register(pattern string, f func(arg string), usage string)
	Serve()
}

type Handler struct {
	router    Router
	service   Service
	presenter Presenter
}

func NewHandler(service Service, router Router) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

func (h *Handler) Init() {

	h.router.Register("add", h.Add, "Добавить новую привычку по названию")
	h.router.Register("list", h.List, "Список привычек")
	h.router.Register("id", h.GetHabit, "Информация о привычке по айди")
	h.router.Register("done", h.Done, "Выполнить привычку по айди")
	h.router.Register("del", h.Delete, "Удалить привычку по айди")

}

func (h *Handler) Run() {
	h.router.Serve()
}

func (h *Handler) Add(arg string) {
	ID, err := h.service.CreateHabit(arg)
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	habit, err := h.service.GetHabit(ID)
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	h.presenter.FormatAdd(habit)
}

func (h *Handler) List(arg string) {
	hs, err := h.service.GetHabits()
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	h.presenter.FormatList(hs)
}

func (h *Handler) GetHabit(arg string) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		h.presenter.FormatError(err)
	}
	habit, err := h.service.GetHabit(int64(id))
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	h.presenter.FormatGetHabit(habit)
}

func (h *Handler) Done(arg string) {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	err = h.service.MarkHabitDone(int64(id))
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	habit, err := h.service.GetHabit(int64(id))
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	h.presenter.FormatDone(habit)
}

func (h *Handler) Delete(arg string) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	habit, err := h.service.GetHabit(int64(id))
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	err = h.service.DeleteHabit(int64(id))
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	h.presenter.FormatDelete(habit)
}
