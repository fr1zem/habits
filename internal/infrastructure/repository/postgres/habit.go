package postgres

import (
	"CLIappHabits/internal/entities"
	"database/sql"
)

type HabitsRepo struct {
	db *sql.DB
}

func NewHabitsRepo(db *sql.DB) *HabitsRepo {
	return &HabitsRepo{db: db}
}

func (r *HabitsRepo) CreateHabit(h entities.Habit) (int64, error) {
	var id int64
	err := r.db.QueryRow(
		`INSERT INTO habits (name, done) VALUES ($1, $2) RETURNING id`,
		h.Name, h.Done,
	).Scan(&id)
	return id, err
}

func (r *HabitsRepo) GetHabit(id int64) (entities.Habit, error) {
	var h entities.Habit
	err := r.db.QueryRow(
		`SELECT id, name, done FROM habits WHERE id=$1`, id,
	).Scan(&h.ID, &h.Name, &h.Done)
	return h, err
}

func (r *HabitsRepo) GetHabits() ([]entities.Habit, error) {
	rows, err := r.db.Query(`SELECT id, name, done FROM habits`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var habits []entities.Habit
	for rows.Next() {
		var h entities.Habit
		if err := rows.Scan(&h.ID, &h.Name, &h.Done); err != nil {
			return nil, err
		}
		habits = append(habits, h)
	}
	return habits, nil
}

func (r *HabitsRepo) MarkHabitDone(id int64) error {
	_, err := r.db.Exec(`UPDATE habits SET done = true WHERE id=$1`, id)
	return err
}
