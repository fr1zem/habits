package entities

import (
	"testing"
	"time"
)

func TestHabit_MarkDone(t *testing.T) {
	h := &Habit{
		HabitID:        1,
		Name:           "Test",
		Repetitions:    0,
		LastRepetition: time.Now(),
	}

	before := h.LastRepetition

	h.MarkDone()

	if h.Repetitions != 1 {
		t.Errorf("Ожидалось 1 повторение, было %d", h.Repetitions)
	}

	if h.LastRepetition.Sub(before) <= 0 {
		t.Errorf("Время должно было увеличиться, но оно уменьшилось на %s", h.LastRepetition)
	}

}
