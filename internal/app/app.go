package app

import (
	"CLIappHabits/internal/infrastructure/repository/postgres"
	"CLIappHabits/internal/transport/CLI"
	"CLIappHabits/internal/usecases"
	"CLIappHabits/pkg/CLIRouter"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Run() {
	connStr := "host=172.24.96.1 port=5432 user=postgres password=postgres dbname=habits sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("sql open: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("db ping: %v", err)
	}

	// 1. Репозиторий
	repo := postgres.NewHabitsRepo(db)

	// 2. Юзкейсы (каждый получает repo)
	creator := usecases.NewCreateHabitUseCase(repo)
	getter := usecases.NewGetHabitUseCase(repo)
	lister := usecases.NewListHabitsUseCase(repo)
	marker := usecases.NewMarkHabitUseCase(repo)
	deleter := usecases.NewDeleteHabitUseCase(repo)

	// 3. Презентер
	presenter := CLI.NewCLIPresenter()

	// 4. Роутер
	router := CLIRouter.NewRouter(os.Args)

	// 5. Хендлер
	handler := CLI.NewHabitHandler(
		router,
		creator,
		getter,
		lister,
		marker,
		deleter,
		presenter,
	)

	handler.Init()
	handler.Run()
}
