package postgres

import (
	"CLIappHabits/internal/entities"
	"database/sql"
	"fmt"
)

type HabitsRepo struct {
	db *sql.DB
}

func NewHabitsRepo(db *sql.DB) *HabitsRepo {
	return &HabitsRepo{db: db}
}

func (r *HabitsRepo) CreateHabit(h *entities.Habit) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		`INSERT INTO habits (name, repetitions, last_repetition) VALUES ($1, $2, $3) RETURNING habit_id`,
		h.Name, h.Repetitions, h.LastRepetition,
	).Scan(&id)
	return id, err
}

func (r *HabitsRepo) GetHabit(id int64) (entities.Habit, error) {
	var h entities.Habit
	err := r.db.QueryRow(
		`SELECT * FROM habits WHERE habit_id=$1`, id,
	).Scan(&h.HabitID, &h.Name, &h.Repetitions, &h.LastRepetition)
	return h, err
}

func (r *HabitsRepo) GetHabits() ([]entities.Habit, error) {
	rows, err := r.db.Query(`SELECT * FROM habits`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var habits []entities.Habit
	for rows.Next() {
		var h entities.Habit
		if err = rows.Scan(&h.HabitID, &h.Name, &h.Repetitions, &h.LastRepetition); err != nil {
			return nil, err
		}
		habits = append(habits, h)
	}
	return habits, nil
}

func (r *HabitsRepo) MarkHabitDone(id int64) error {

	var h entities.Habit
	err := r.db.QueryRow(
		`SELECT * FROM habits WHERE habit_id=$1`, id,
	).Scan(&h.HabitID, &h.Name, &h.Repetitions, &h.LastRepetition)
	if err != nil {
		return fmt.Errorf("get habit inside done: [%w]", err)
	}
	h.MarkDone()

	_, err = r.db.Exec(`UPDATE habits SET last_repetition=$1,
                  repetitions=$2 WHERE habit_id=$3`, h.LastRepetition, h.Repetitions, id)
	return err
}

func (r *HabitsRepo) DeleteHabit(ID int64) error {
	_, err := r.db.Exec("DELETE FROM habits WHERE habit_id=$1", ID)
	return err
}
