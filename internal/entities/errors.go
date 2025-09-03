package entities

import "errors"

var (
	ErrHabitNotExists     = errors.New("habit doesn't exist")
	ErrHabitAlreadyExists = errors.New("habit is already exists")
	ErrEmptyName          = errors.New("name is empty")
)
