package entities

import "errors"

var (
	ErrHabitNotExists     = errors.New("habit doesn't exist")
	ErrHabitsNotExists    = errors.New("habits doesn't exists")
	ErrHabitAlreadyExists = errors.New("habit is already exists")
	ErrEmptyName          = errors.New("name is empty")
)
