package usecases

import "CLIappHabits/internal/entities"

type HabitsRepository interface {
	CreateHabit(h *entities.Habit) (int64, error)
	GetHabit(ID int64) (entities.Habit, error)
	GetHabits() ([]entities.Habit, error)
	MarkHabitDone(ID int64) error
}

type HabitsService struct {
	repo HabitsRepository
}

func NewHabitsService(r HabitsRepository) *HabitsService {
	return &HabitsService{repo: r}
}

func (s *HabitsService) CreateHabit(name string) (int64, error) {
	habit := &entities.Habit{}
	return s.repo.CreateHabit(habit)
}

func (s *HabitsService) GetHabit(ID int64) (entities.Habit, error) {
	return s.repo.GetHabit(ID)
}

func (s *HabitsService) GetHabits() ([]entities.Habit, error) {
	return s.repo.GetHabits()
}

func (s *HabitsService) MarkHabitDone(ID int64) error {
	return s.MarkHabitDone(ID)
}
