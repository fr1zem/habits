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

func GetHabbitInfo(name string, repo repository) (*habbit, error) {
	var h *habbit

	err := repo.db.QueryRow("SELECT * FROM habbits WHERE name = $1", name).
		Scan(&h.habbitID, &h.name, &h.count, &h.lastRepeteation)
	if err != nil {
		return nil, fmt.Errorf("query habbit by name: [%w]", err)
	}

	return h, nil
}

func PrintHabbitInfo(name string, repo repository) error {

	h, err := GetHabbitInfo(name, repo)
	if err != nil {
		return err
	}

	fmt.Printf("Habbit %s has %d repetetion with last time %s\n", h.name, h.count, h.lastRepeteation)

	return nil
}

func PrintAllHabbitsInfo(repo repository) error {
	rows, err := repo.db.Query("SELECT * FROM habbits")
	if err != nil {
		return fmt.Errorf("query all habbits info: [%w]", err)
	}

	for rows.Next() {
		var h *habbit
		err = rows.Scan(&h.habbitID, &h.name, &h.count, &h.lastRepeteation)
		if err != nil {
			return fmt.Errorf("scan habbit in all: [%w]", err)
		}

		fmt.Printf("Habbit %s has %d repetetion with last time %s\n", h.name, h.count, h.lastRepeteation)
	}

	return nil
}

func AddNewHabbit(name string, repo repository) error {
	_, err := repo.db.Exec("INSERT INTO habbits (name, count, lastRepeteation) VALUES ($1, $2, $3)",
		name, 0, time.Now())
	if err != nil {
		return fmt.Errorf("exec add new habbit: [%w]", err)
	}
	return nil
}

func AddRepeteationOfHabbit(name string, repo repository) error {

	h, err := GetHabbitInfo(name, repo)
	if err != nil {
		return err
	}

	h.count++

	_, err = repo.db.Exec(`UPDATE habbits SET count = $1, last_repeteation = $2
               WHERE name = $3`, h.count, time.Now(), name)
	if err != nil {
		return fmt.Errorf("update habbit: [%w]", err)
	}


	return nil
}

func Scan() {
	
	for i, v := range os.Args {
		if i == 0 {
			continue
		}

		switch v {
		case "-add":
			i++
			AddNewHabbit(os.Args[i], )
		case "-"
		}

	}
}

func Run() {
	var repo *repository
	initResources(repo)

}

func main() {
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
