package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect membuka koneksi ke database Postgres menggunakan connection string.
func Connect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke database: %w", err)
	}

	// Cek koneksi
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
