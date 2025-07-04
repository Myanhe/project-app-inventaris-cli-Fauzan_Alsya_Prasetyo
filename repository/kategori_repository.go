package repository

import (
	"database/sql"
	"errors"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
)

type KategoriRepository struct {
	DB *sql.DB
}

func NewKategoriRepository(db *sql.DB) *KategoriRepository {
	return &KategoriRepository{DB: db}
}

func (r *KategoriRepository) Create(kategori *model.Kategori) error {
	query := "INSERT INTO kategori (nama, deskripsi) VALUES ($1, $2) RETURNING id"
	return r.DB.QueryRow(query, kategori.Nama, kategori.Deskripsi).Scan(&kategori.ID)
}

func (r *KategoriRepository) GetAll() ([]model.Kategori, error) {
	rows, err := r.DB.Query("SELECT id, nama, deskripsi FROM kategori")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kategoris []model.Kategori
	for rows.Next() {
		var k model.Kategori
		if err := rows.Scan(&k.ID, &k.Nama, &k.Deskripsi); err != nil {
			return nil, err
		}
		kategoris = append(kategoris, k)
	}
	return kategoris, nil
}

func (r *KategoriRepository) GetByID(id int) (*model.Kategori, error) {
	var k model.Kategori
	err := r.DB.QueryRow("SELECT id, nama, deskripsi FROM kategori WHERE id = $1", id).
		Scan(&k.ID, &k.Nama, &k.Deskripsi)
	if err == sql.ErrNoRows {
		return nil, errors.New("kategori tidak ditemukan")
	}
	return &k, err
}

func (r *KategoriRepository) Update(kategori *model.Kategori) error {
	_, err := r.DB.Exec("UPDATE kategori SET nama = $1, deskripsi = $2 WHERE id = $3",
		kategori.Nama, kategori.Deskripsi, kategori.ID)
	return err
}

func (r *KategoriRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM kategori WHERE id = $1", id)
	return err
}
