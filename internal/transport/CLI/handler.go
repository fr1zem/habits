package CLI

import (
	"CLIappHabits/internal/entities"
	"strconv"
)

type Service interface {
	CreateHabit(name string) (int64, error)
	GetHabit(ID int64) (entities.Habit, error)
	GetHabits() ([]entities.Habit, error)
	MarkHabitDone(ID int64) error
	DeleteHabit(ID int64) error
}

type Router interface {
	Register(pattern string, f func(args []string), usage string)
	Run()
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
	h.router.Run()
}

func (h *Handler) Add(arg []string) {
	ID, err := h.service.CreateHabit(arg[0])
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

func (h *Handler) List(args []string) {
	hs, err := h.service.GetHabits()
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	h.presenter.FormatList(hs)
}

func (h *Handler) GetHabit(args []string) {
	id, err := strconv.Atoi(args[0])
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

func (h *Handler) Done(args []string) {
	id, err := strconv.Atoi(args[0])
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

func (h *Handler) Delete(args []string) {
	id, err := strconv.Atoi(args[0])
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
