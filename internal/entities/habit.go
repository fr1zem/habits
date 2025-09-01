package entities

import "time"

type Habit struct {
	HabitID        int64
	Name           string
	Repetitions    int64
	LastRepetition time.Time
}

/*func (h *Habit) MarkDone(now time.Time) {
	h.Repetitions++
	h.LastRepetition = now
}*/
