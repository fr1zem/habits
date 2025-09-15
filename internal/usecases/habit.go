package usecases

import (
	"CLIappHabits/internal/entities"
	"time"
)

type HabitsRepository interface {
	CreateHabit(h *entities.Habit) (int64, error)
	GetHabit(ID int64) (entities.Habit, error)
	ListHabits() ([]entities.Habit, error)
	MarkHabitDone(ID int64) error
	DeleteHabit(ID int64) error
}

// ---------- Create Habit ----------
type CreateHabitInputDTO struct {
	Name string
}

type CreateHabitOutputDTO struct {
	HabitID int64
}

type CreateHabitUseCase struct {
	repo HabitsRepository
}

type HabitCreator interface {
	CreateHabit(CreateHabitInputDTO) (CreateHabitOutputDTO, error)
}

func NewCreateHabitUseCase(repo HabitsRepository) HabitCreator {
	return &CreateHabitUseCase{repo: repo}
}

func (uc *CreateHabitUseCase) CreateHabit(input CreateHabitInputDTO) (CreateHabitOutputDTO, error) {
	habit := &entities.Habit{
		Name:        input.Name,
		Repetitions: 0,
	}

	id, err := uc.repo.CreateHabit(habit)
	if err != nil {
		return CreateHabitOutputDTO{}, err
	}

	return CreateHabitOutputDTO{HabitID: id}, err
}

// ---------- Get Habit ----------
type GetHabitInputDTO struct {
	HabitID int64
}

type GetHabitOutputDTO struct {
	HabitID        int64
	Name           string
	Repetitions    int64
	LastRepetition time.Time
}

type HabitGetter interface {
	GetHabit(GetHabitInputDTO) (GetHabitOutputDTO, error)
}

func NewGetHabitUseCase(repo HabitsRepository) HabitGetter {
	return &GetHabitUseCase{repo: repo}
}

type GetHabitUseCase struct {
	repo HabitsRepository
}

func (uc *GetHabitUseCase) GetHabit(input GetHabitInputDTO) (GetHabitOutputDTO, error) {
	habit, err := uc.repo.GetHabit(input.HabitID)
	if err != nil {
		return GetHabitOutputDTO{}, err
	}

	return GetHabitOutputDTO{
		HabitID:        habit.HabitID,
		Name:           habit.Name,
		Repetitions:    habit.Repetitions,
		LastRepetition: habit.LastRepetition,
	}, nil
}

// ---------- List Habits ----------

type ListHabitsOutputDTO struct {
	Habits []GetHabitOutputDTO
}

type HabitLister interface {
	ListHabits() (ListHabitsOutputDTO, error)
}

func NewListHabitsUseCase(repo HabitsRepository) HabitLister {
	return &ListHabitsUseCase{repo: repo}
}

type ListHabitsUseCase struct {
	repo HabitsRepository
}

func (uc *ListHabitsUseCase) ListHabits() (ListHabitsOutputDTO, error) {
	habits, err := uc.repo.ListHabits()
	if err != nil {
		return ListHabitsOutputDTO{}, err
	}

	result := make([]GetHabitOutputDTO, 0, len(habits))
	for _, h := range habits {
		result = append(result, GetHabitOutputDTO{
			HabitID:        h.HabitID,
			Name:           h.Name,
			Repetitions:    h.Repetitions,
			LastRepetition: h.LastRepetition,
		})
	}

	return ListHabitsOutputDTO{Habits: result}, nil
}

// ---------- Mark Habit As Completed ----------
type MarkHabitInputDTO struct {
	HabitID int64
}

type HabitMarker interface {
	MarkHabit(MarkHabitInputDTO) error
}

func NewMarkHabitUseCase(repo HabitsRepository) HabitMarker {
	return &MarkHabitUseCase{repo: repo}
}

type MarkHabitUseCase struct {
	repo HabitsRepository
}

func (uc *MarkHabitUseCase) MarkHabit(input MarkHabitInputDTO) error {
	err := uc.repo.MarkHabitDone(input.HabitID)
	if err != nil {
		return err
	}

	return nil
}

// ---------- Delete Habit ----------
type DeleteHabitInputDTO struct {
	HabitID int64
}

type HabitDeleter interface {
	DeleteHabit(DeleteHabitInputDTO) error
}

func NewDeleteHabitUseCase(repo HabitsRepository) HabitDeleter {
	return &DeleteHabitUseCase{repo: repo}
}

type DeleteHabitUseCase struct {
	repo HabitsRepository
}

func (uc *DeleteHabitUseCase) DeleteHabit(input DeleteHabitInputDTO) error {
	err := uc.repo.DeleteHabit(input.HabitID)
	if err != nil {
		return err
	}

	return nil
}
