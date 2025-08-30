package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type repository struct {
	db *sql.DB
}

func (repo *repository) initRepository() error {
	var err error
	repo.db, err = sql.Open("postgres", "host=172.24.96.1 port=5432 user=postgres "+
		"password=postgres dbname=habits sslmode=disable")
	if err != nil {
		return fmt.Errorf("open db: [%w]", err)
	}
	return err
}

func (repo *repository) GetHabit(name string) (*habit, error) {
	var h habit

	err := repo.db.QueryRow("SELECT * FROM habits WHERE name = $1", name).
		Scan(&h.habitID, &h.name, &h.count, &h.lastRepetetion)
	if err != nil {
		return nil, fmt.Errorf("query habit by name: [%w]", err)
	}

	return &h, nil
}

func (repo *repository) GetHabits() ([]*habit, error) {
	var listLen int64
	err := repo.db.QueryRow("SELECT COUNT(*) FROM habits").Scan(&listLen)
	if err != nil {
		return nil, fmt.Errorf("query row list len: [%w]", err)
	}

	rows, err := repo.db.Query("SELECT * FROM habits")
	if err != nil {
		return nil, fmt.Errorf("query all habits info: [%w]", err)
	}
	defer rows.Close()

	var habits []*habit
	for rows.Next() {
		h := new(habit)
		err = rows.Scan(&h.habitID, &h.name, &h.count, &h.lastRepetetion)
		if err != nil {
			return nil, fmt.Errorf("scan habit in all: [%w]", err)
		}
		habits = append(habits, h)
	}

	return habits, nil
}

func (repo *repository) UpdateRepetetionOfHabit(name string, count int64) error {
	_, err := repo.db.Exec(`UPDATE habits SET count = $1, last_repetetion = $2
               WHERE name = $3`, count, time.Now(), name)
	if err != nil {
		return fmt.Errorf("update habit: [%w]", err)
	}

	return nil
}

func (repo *repository) AddNewHabit(name string) error {
	_, err := repo.db.Exec("INSERT INTO habits (name, count, last_repetetion) VALUES ($1, $2, $3)",
		name, 0, time.Now())
	if err != nil {
		return fmt.Errorf("exec add new habit: [%w]", err)
	}

	return nil
}

type Resources struct {
	db *sql.DB
}

func initResources(repo *repository) {
	err := repo.initRepository()
	if err != nil {
		log.Fatal(fmt.Errorf("init repository [%w]", err))
	}
}

type habit struct {
	habitID        int64
	name           string
	count          int64
	lastRepetetion time.Time
}

func PrintHabitInfo(name string, repo *repository) error {

	h, err := repo.GetHabit(name)
	if err != nil {
		return err
	}

	fmt.Printf("Habit %s has %d repetetion with last time %s\n", h.name, h.count, h.lastRepetetion)

	return nil
}

func PrintAllHabitsInfo(repo *repository) error {
	hs, err := repo.GetHabits()
	if err != nil {
		return err
	}

	for _, h := range hs {
		fmt.Printf("Habit %s has %d repetetion with last time %s\n", h.name, h.count, h.lastRepetetion)
	}

	return nil
}

func AddNewHabit(name string, repo *repository) error {
	err := repo.AddNewHabit(name)
	if err != nil {
		return err
	}

	return nil
}

func AddRepetetionOfHabit(name string, repo *repository) error {
	h, err := repo.GetHabit(name)
	if err != nil {
		return err
	}

	h.count++

	err = repo.UpdateRepetetionOfHabit(name, h.count)
	if err != nil {
		return err
	}

	return nil
}

func Scan(repo *repository) {

	for i, v := range os.Args {
		if i == 0 {
			continue
		}

		if i == 1 {
			switch v {
			case "add":
				i++
				err := AddNewHabit(os.Args[i], repo)
				if err != nil {
					log.Fatal(err)
				}
			case "done":
				i++
				err := AddRepetetionOfHabit(os.Args[i], repo)
				if err != nil {
					log.Fatal(err)
				}
			case "list":
				err := PrintAllHabitsInfo(repo)
				if err != nil {
					log.Fatal(err)
				}
			case "info":
				i++
				err := PrintHabitInfo(os.Args[i], repo)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}

func Run() {

	var repo repository
	var err error

	repo.db, err = sql.Open("postgres", "host=172.24.96.1 port=5432 user=postgres "+
		"password=postgres dbname=habits sslmode=disable")
	if err != nil {
		log.Fatal(fmt.Errorf("open db: [%w]", err))
	}
	Scan(&repo)
}

func main() {
	Run()
}
