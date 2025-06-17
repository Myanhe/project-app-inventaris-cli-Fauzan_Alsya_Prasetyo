package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

func Connect() *sql.DB {
    connStr := "user=postgres password=yourpassword dbname=yourdb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    return db
}