package entities

import "time"

type Habit struct {
	Habit_ID       int64
	Name           string
	Count          int64
	LastRepetition time.Time
}

func (h *Habit) MarkDone(now time.Time) {
	h.Count++
	h.LastRepetition = now
}
