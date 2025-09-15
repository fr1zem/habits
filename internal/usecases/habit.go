package usecases

import "CLIappHabits/internal/entities"

type HabitsRepository interface {
	CreateHabit(h *entities.Habit) (int64, error)
	GetHabit(ID int64) (entities.Habit, error)
	ListHabits() ([]entities.Habit, error)
	MarkHabitDone(ID int64) error
	DeleteHabit(ID int64) error
}

type HabitsService struct {
	repo HabitsRepository
}

func NewHabitsService(r HabitsRepository) *HabitsService {
	return &HabitsService{repo: r}
}

// ---------- Create Habit ----------
type CreateHabitInputDTO struct {
	name string
}

type CreateHabitOutputDTO struct {
	habitID int64
}

type CreateHabitUseCase struct {
	repo HabitsRepository
}

type HabitCreator interface {
	CreateHabit(CreateHabitInputDTO) (CreateHabitOutputDTO, error)
}

func (uc *CreateHabitUseCase) CreateHabit(input CreateHabitInputDTO) (CreateHabitOutputDTO, error) {
	habit := &entities.Habit{
		Name:        input.name,
		Repetitions: 0,
	}

	id, err := uc.repo.CreateHabit(habit)
	if err != nil {
		return CreateHabitOutputDTO{}, err
	}

	return CreateHabitOutputDTO{habitID: id}, err
}

// ---------- Get Habit ----------
type GetHabitInputDTO struct {
	HabitID int64
}

type GetHabitOutputDTO struct {
	ID          int64
	Name        string
	Repetitions int64
}

type HabitGetter interface {
	GetHabit(GetHabitInputDTO) (GetHabitOutputDTO, error)
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
		ID:          habit.HabitID,
		Name:        habit.Name,
		Repetitions: habit.Repetitions,
	}, nil
}

// ---------- List Habits ----------

type ListHabitsOutputDTO struct {
	Habits []GetHabitOutputDTO
}

type HabitLister interface {
	ListHabits() (ListHabitsOutputDTO, error)
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
			ID:          h.HabitID,
			Name:        h.Name,
			Repetitions: h.Repetitions,
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
