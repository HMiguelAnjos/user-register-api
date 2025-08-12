package db

import (
    "database/sql"
    "log"
    _ "github.com/jackc/pgx/v5/stdlib"
    "os"
)

func NewPostgresDB() *sql.DB {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL not set")
    }
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    return db
}
