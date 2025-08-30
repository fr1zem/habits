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
		"password=postgres dbname=habbits sslmode=disable")
	if err != nil {
		return fmt.Errorf("open db: [%w]", err)
	}
	return err
}

func (repo *repository) GetHabbit(name string) (*habbit, error) {
	var h habbit

	err := repo.db.QueryRow("SELECT * FROM habbits WHERE name = $1", name).
		Scan(&h.habbitID, &h.name, &h.count, &h.lastRepeteation)
	if err != nil {
		return nil, fmt.Errorf("query habbit by name: [%w]", err)
	}

	return &h, nil
}

func (repo *repository) GetHabbits() ([]*habbit, error) {
	var listLen int64
	err := repo.db.QueryRow("SELECT COUNT(*) FROM habbits").Scan(&listLen)
	if err != nil {
		return nil, fmt.Errorf("query row list len: [%w]", err)
	}

	rows, err := repo.db.Query("SELECT * FROM habbits")
	if err != nil {
		return nil, fmt.Errorf("query all habbits info: [%w]", err)
	}
	defer rows.Close()

	var habbits []*habbit
	for rows.Next() {
		h := new(habbit)
		err = rows.Scan(&h.habbitID, &h.name, &h.count, &h.lastRepeteation)
		if err != nil {
			return nil, fmt.Errorf("scan habbit in all: [%w]", err)
		}
		habbits = append(habbits, h)
	}

	return habbits, nil
}

func (repo *repository) UpdateRepeteationOfHabbit(name string, count int64) error {
	_, err := repo.db.Exec(`UPDATE habbits SET count = $1, last_repeteation = $2
               WHERE name = $3`, count, time.Now(), name)
	if err != nil {
		return fmt.Errorf("update habbit: [%w]", err)
	}

	return nil
}

func (repo *repository) AddNewHabbit(name string) error {
	_, err := repo.db.Exec("INSERT INTO habbits (name, count, last_repeteation) VALUES ($1, $2, $3)",
		name, 0, time.Now())
	if err != nil {
		return fmt.Errorf("exec add new habbit: [%w]", err)
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

type habbit struct {
	habbitID        int64
	name            string
	count           int64
	lastRepeteation time.Time
}

func PrintHabbitInfo(name string, repo *repository) error {

	h, err := repo.GetHabbit(name)
	if err != nil {
		return err
	}

	fmt.Printf("Habbit %s has %d repetetion with last time %s\n", h.name, h.count, h.lastRepeteation)

	return nil
}

func PrintAllHabbitsInfo(repo *repository) error {
	hs, err := repo.GetHabbits()
	if err != nil {
		return err
	}

	for _, h := range hs {
		fmt.Printf("Habbit %s has %d repetetion with last time %s\n", h.name, h.count, h.lastRepeteation)
	}

	return nil
}

func AddNewHabbit(name string, repo *repository) error {
	err := repo.AddNewHabbit(name)
	if err != nil {
		return err
	}

	return nil
}

func AddRepeteationOfHabbit(name string, repo *repository) error {
	h, err := repo.GetHabbit(name)
	if err != nil {
		return err
	}

	h.count++

	err = repo.UpdateRepeteationOfHabbit(name, h.count)
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
				err := AddNewHabbit(os.Args[i], repo)
				if err != nil {
					log.Fatal(err)
				}
			case "done":
				i++
				err := AddRepeteationOfHabbit(os.Args[i], repo)
				if err != nil {
					log.Fatal(err)
				}
			case "list":
				err := PrintAllHabbitsInfo(repo)
				if err != nil {
					log.Fatal(err)
				}
			case "info":
				i++
				err := PrintHabbitInfo(os.Args[i], repo)
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
		"password=postgres dbname=habbits sslmode=disable")
	if err != nil {
		log.Fatal(fmt.Errorf("open db: [%w]", err))
	}
	Scan(&repo)
}

func main() {
	Run()
}
