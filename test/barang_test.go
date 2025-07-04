package test

import (
	"bytes"
	"os"
	"testing"

	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/handler"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"
)

func TestKategoriHandler_ListKategori(t *testing.T) {
	db := setupTestDB()
	repo := repository.NewKategoriRepository(db)
	repo.Create(&model.Kategori{Nama: "Test", Deskripsi: "Test Desc"})
	h := &handler.KategoriHandler{Repo: repo}

	// Redirect stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	h.ListKategori()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if output == "" || output == "\n" {
		t.Error("Output ListKategori kosong")
	}
}