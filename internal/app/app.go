package app

import (
	"CLIappHabits/internal/infrastructure/repository/postgres"
	"CLIappHabits/internal/transport/CLI"
	"CLIappHabits/internal/usecases"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Run() {

	db, err := sql.Open("postgres", "host=172.24.96.1 port=5432 user=postgres "+
		"password=postgres dbname=habits sslmode=disable")
	if err != nil {
		log.Fatal("sql open: [%w]", err)
	}

	repo := postgres.NewHabitsRepo(db)

	services := usecases.NewHabitsService(repo)

	handler := CLI.NewHandler(services)

	handler.Run()
}
