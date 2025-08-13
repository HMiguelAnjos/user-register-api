package main

import (
	"log"
	"net/http"

	controllers "userregisterapi/internal/adapters/http/controllers"
	routers "userregisterapi/internal/adapters/http/routes"
	"userregisterapi/internal/adapters/id"
	"userregisterapi/internal/adapters/logger"
	"userregisterapi/internal/app/ports"
	app "userregisterapi/internal/app/usecase"
	"userregisterapi/internal/config"
	"userregisterapi/internal/infrastructure/repository/memory"

	pgdb "userregisterapi/internal/infrastructure/db"
	pgrepo "userregisterapi/internal/infrastructure/repository/postgres"
)

func main() {
	cfg := config.Load()

	var repo ports.UserRepository

	if cfg.DatabaseURL != "" {
		db := pgdb.NewPostgresDB()
		defer db.Close()
		repo = pgrepo.NewUserRepoPostgres(db)
		log.Println("[INFO] Using PostgreSQL repository")
	} else {
		repo = memory.NewUserRepoMemory()
		log.Println("[WARN] DATABASE_URL not set, using in-memory repository")
	}

	ids := id.NewUUIDGenerator()
	lg := logger.NewStdLogger()

	svc := app.NewUserService(repo, ids, lg)

	ctrl := controllers.NewUserController(svc)
	router := routers.NewRouter(ctrl)

	log.Printf("Starting server on %s", cfg.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, router))
}
