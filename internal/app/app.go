package app

import (
	"CLIappHabits/internal/config"
	"CLIappHabits/internal/infrastructure/repository/postgres"
	"CLIappHabits/internal/transport/CLI"
	"CLIappHabits/internal/transport/Web/v1/httpGin"
	"CLIappHabits/internal/usecases"
	"CLIappHabits/pkg/CLIRouter"
	"CLIappHabits/pkg/Postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func RunCLI() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName,
		cfg.Database.SSLMode)
	db, err := Postgres.NewPostgres(connStr)
	if err != nil {
		log.Fatal(err)
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

func RunWebGin() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName,
		cfg.Database.SSLMode)
	db, err := Postgres.NewPostgres(connStr)
	if err != nil {
		log.Fatal(err)
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

	r := gin.Default()

	handler := httpGin.NewHabitHandler(creator, getter, lister, marker, deleter)

	handler.InitRoutes(r)

	serverConnStr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	err = r.Run(serverConnStr)
	if err != nil {
		log.Fatal(err)
	}

}
