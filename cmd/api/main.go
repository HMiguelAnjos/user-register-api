package main

import (
	"log"
	"net/http"
	"os"

	httpadapter "userregisterapi/internal/adapters/http"
	"userregisterapi/internal/adapters/id"
	"userregisterapi/internal/adapters/logger"
	"userregisterapi/internal/adapters/repository/memory"
	"userregisterapi/internal/app/ports"
	app "userregisterapi/internal/app/usecase"
	"userregisterapi/internal/config"

	pgdb "userregisterapi/internal/infrastructure/db"
	pgrepo "userregisterapi/internal/infrastructure/repository/postgres"
)

func main() {
	cfg := config.Load()

	var repo ports.TaskRepository

	// Se existir DATABASE_URL no ambiente, usa Postgres; caso contrário, usa memória
	if os.Getenv("DATABASE_URL") != "" {
		db := pgdb.NewPostgresDB()
		defer db.Close()
		repo = pgrepo.NewTaskRepoPostgres(db)
		log.Println("[INFO] Using PostgreSQL repository")
	} else {
		repo = memory.NewTaskRepoMemory()
		log.Println("[WARN] DATABASE_URL not set, using in-memory repository")
	}

	// Adapters
	ids := id.NewUUIDGenerator()
	lg := logger.NewStdLogger()

	// Application service (use cases)
	svc := app.NewTaskService(repo, ids, lg)

	// Controllers + Router
	ctrl := httpadapter.NewTaskController(svc)
	router := httpadapter.NewRouter(ctrl)

	log.Printf("Starting server on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, router); err != nil {
		log.Fatal(err)
	}
}
