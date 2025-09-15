package CLI

import (
	"CLIappHabits/internal/usecases"
	"strconv"
)

type Router interface {
	Register(pattern string, f func(args []string), usage string)
	Run()
}

type HabitPresenter interface {
	FormatGetHabit(usecases.GetHabitOutputDTO)
	FormatList(usecases.ListHabitsOutputDTO)
	FormatCompleted(usecases.GetHabitOutputDTO)
	FormatDelete(usecases.GetHabitOutputDTO)
	FormatError(error)
}

type HabitHandler struct {
	router    Router
	creator   usecases.HabitCreator
	getter    usecases.HabitGetter
	lister    usecases.HabitLister
	marker    usecases.HabitMarker
	deleter   usecases.HabitDeleter
	presenter HabitPresenter
}

func NewHabitHandler(
	router Router,
	creator usecases.HabitCreator,
	getter usecases.HabitGetter,
	lister usecases.HabitLister,
	marker usecases.HabitMarker,
	deleter usecases.HabitDeleter,
	presenter HabitPresenter,
) *HabitHandler {
	return &HabitHandler{
		router:    router,
		creator:   creator,
		getter:    getter,
		lister:    lister,
		marker:    marker,
		deleter:   deleter,
		presenter: presenter,
	}
}

func (h *HabitHandler) Init() {

	h.router.Register("add", h.Add, "Добавить новую привычку по названию")
	h.router.Register("list", h.List, "Список привычек")
	h.router.Register("id", h.GetHabit, "Информация о привычке по айди")
	h.router.Register("done", h.Completed, "Выполнить привычку по айди")
	h.router.Register("del", h.Delete, "Удалить привычку по айди")

}

func (h *HabitHandler) Run() {
	h.router.Run()
}

func (h *HabitHandler) Add(args []string) {
	output, err := h.creator.CreateHabit(usecases.CreateHabitInputDTO{Name: args[0]})
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	habit, err := h.getter.GetHabit(usecases.GetHabitInputDTO{HabitID: output.HabitID})
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	h.presenter.FormatGetHabit(habit)
}

func (h *HabitHandler) List(args []string) {
	output, err := h.lister.ListHabits()
	if err != nil {
		h.presenter.FormatError(err)
		return
	}
	h.presenter.FormatList(output)
}

func (h *HabitHandler) GetHabit(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	output, err := h.getter.GetHabit(usecases.GetHabitInputDTO{HabitID: int64(id)})
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	h.presenter.FormatGetHabit(output)
}

func (h *HabitHandler) Completed(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	err = h.marker.MarkHabit(usecases.MarkHabitInputDTO{HabitID: int64(id)})
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	// достаём свежую привычку
	output, err := h.getter.GetHabit(usecases.GetHabitInputDTO{HabitID: int64(id)})
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	h.presenter.FormatCompleted(output)
}

func (h *HabitHandler) Delete(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	habit, err := h.getter.GetHabit(usecases.GetHabitInputDTO{HabitID: int64(id)})
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	err = h.deleter.DeleteHabit(usecases.DeleteHabitInputDTO{HabitID: int64(id)})
	if err != nil {
		h.presenter.FormatError(err)
		return
	}

	h.presenter.FormatDelete(habit)
}
