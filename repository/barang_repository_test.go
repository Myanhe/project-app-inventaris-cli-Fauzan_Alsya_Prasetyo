package repository

import (
    "testing"
    "database/sql"
    _ "github.com/lib/pq"
)

func TestGetAll(t *testing.T) {
    db, err := sql.Open("postgres", "user=postgres password=yourpassword dbname=yourdb sslmode=disable")
    if err != nil {
        t.Fatal(err)
    }
    repo := &BarangRepository{DB: db}
    barangs, err := repo.GetAll()
    if err != nil {
        t.Errorf("GetAll() error = %v", err)
    }
    if barangs == nil {
        t.Errorf("GetAll() returned nil")
    }
}