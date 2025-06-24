package test

import (
    "database/sql"
    "os"
    "testing"

    _ "github.com/lib/pq"
    "github.com/stretchr/testify/assert"
    "project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
    "project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"
)

func setupTestDB() *sql.DB {
    connStr := os.Getenv("TEST_DB_CONN")
    db, _ := sql.Open("postgres", connStr)
    db.Exec("TRUNCATE kategori RESTART IDENTITY CASCADE")
    return db
}

func TestKategoriRepository_CRUD(t *testing.T) {
    db := setupTestDB()
    repo := repository.NewKategoriRepository(db)

    // Create
    k := &model.Kategori{Nama: "Elektronik", Deskripsi: "Peralatan elektronik"}
    err := repo.Create(k)
    assert.Nil(t, err)
    assert.NotZero(t, k.ID)

    // GetAll
    list, err := repo.GetAll()
    assert.Nil(t, err)
    assert.True(t, len(list) > 0)

    // GetByID
    kat, err := repo.GetByID(k.ID)
    assert.Nil(t, err)
    assert.Equal(t, "Elektronik", kat.Nama)

    // Update
    k.Nama = "Elektronik Baru"
    err = repo.Update(k)
    assert.Nil(t, err)
    kat, _ = repo.GetByID(k.ID)
    assert.Equal(t, "Elektronik Baru", kat.Nama)

    // Delete
    err = repo.Delete(k.ID)
    assert.Nil(t, err)
    _, err = repo.GetByID(k.ID)
    assert.NotNil(t, err)
}