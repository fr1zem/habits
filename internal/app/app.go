package app

import (
	"CLIappHabits/internal/infrastructure/repository/postgres"
	"CLIappHabits/internal/transport/CLI"
	"CLIappHabits/internal/usecases"
	"CLIappHabits/pkg/CLIRouter"
	"CLIappHabits/pkg/Postgres"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Run() {

	db, err := Postgres.NewPostgres(
		Postgres.Client{
			Host:     "172.24.96.1",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DBName:   "habits",
			SSLMode:  "disable",
		})
	if err != nil {
		log.Fatal("sql open: [%w]", err)
	}
	/*Это типа конструктор NewXXX(), но такой конструктор должен принимать интерфейс, а он принимает *sql.DB
	Получается это будет не конструктор с инъекцией, а конструктор БД исходя из конфига*/
	repo := postgres.NewHabitsRepo(db)

	services := usecases.NewHabitsService(repo) //Вот это уже правильный конструктор

	router := CLIRouter.NewRouter(os.Args)

	handler := CLI.NewHandler(services, router)

	handler.Init()

	handler.Run()
}
