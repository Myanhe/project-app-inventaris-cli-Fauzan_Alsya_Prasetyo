package repository

import (
	"database/sql"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
)

type KategoriRepository struct {
	DB *sql.DB
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
	row := r.DB.QueryRow("SELECT id, nama, deskripsi FROM kategori WHERE id=$1", id)
	var k model.Kategori
	if err := row.Scan(&k.ID, &k.Nama, &k.Deskripsi); err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *KategoriRepository) Create(nama, deskripsi string) error {
	_, err := r.DB.Exec("INSERT INTO kategori (nama, deskripsi) VALUES ($1, $2)", nama, deskripsi)
	return err
}

func (r *KategoriRepository) Update(id int, nama, deskripsi string) error {
	_, err := r.DB.Exec("UPDATE kategori SET nama=$1, deskripsi=$2 WHERE id=$3", nama, deskripsi, id)
	return err
}

func (r *KategoriRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM kategori WHERE id=$1", id)
	return err
}
