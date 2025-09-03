package mocks

import "CLIappHabits/internal/entities"

type MockHabitRepo struct {
	CreateHabitFunc   func(h *entities.Habit) (int64, error)
	GetHabitFunc      func(ID int64) (entities.Habit, error)
	GetHabitsFunc     func() ([]entities.Habit, error)
	MarkHabitDoneFunc func(ID int64) error
	DeleteHabitFunc   func(ID int64) error
}

func (m *MockHabitRepo) CreateHabit(h *entities.Habit) (int64, error) {
	return m.CreateHabitFunc(h)
}

func (m *MockHabitRepo) GetHabit(ID int64) (entities.Habit, error) {
	return m.GetHabitFunc(ID)
}

func (m *MockHabitRepo) GetHabits() ([]entities.Habit, error) {
	return m.GetHabitsFunc()
}

func (m *MockHabitRepo) MarkHabitDone(ID int64) error {
	return m.MarkHabitDoneFunc(ID)
}

func (m *MockHabitRepo) DeleteHabit(ID int64) error {
	return m.DeleteHabitFunc(ID)
}
