package repository

import (
	"database/sql"
	"errors"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
)

type BarangRepository struct {
	DB *sql.DB
}

func NewBarangRepository(db *sql.DB) *BarangRepository {
	return &BarangRepository{DB: db}
}

func (r *BarangRepository) Create(barang *model.Barang) error {
	query := `INSERT INTO barang (nama, harga, tanggal_beli, kategori_id, total_penggunaan)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.DB.QueryRow(query, barang.Nama, barang.Harga, barang.TanggalBeli, barang.KategoriID, barang.TotalPenggunaan).Scan(&barang.ID)
}

func (r *BarangRepository) GetAll() ([]model.Barang, error) {
	rows, err := r.DB.Query("SELECT id, nama, harga, tanggal_beli, kategori_id, total_penggunaan FROM barang")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var barangs []model.Barang
	for rows.Next() {
		var b model.Barang
		if err := rows.Scan(&b.ID, &b.Nama, &b.Harga, &b.TanggalBeli, &b.KategoriID, &b.TotalPenggunaan); err != nil {
			return nil, err
		}
		barangs = append(barangs, b)
	}
	return barangs, nil
}

func (r *BarangRepository) GetByID(id int) (*model.Barang, error) {
	var b model.Barang
	err := r.DB.QueryRow("SELECT id, nama, harga, tanggal_beli, kategori_id, total_penggunaan FROM barang WHERE id = $1", id).
		Scan(&b.ID, &b.Nama, &b.Harga, &b.TanggalBeli, &b.KategoriID, &b.TotalPenggunaan)
	if err == sql.ErrNoRows {
		return nil, errors.New("barang tidak ditemukan")
	}
	return &b, err
}

func (r *BarangRepository) Update(barang *model.Barang) error {
	_, err := r.DB.Exec("UPDATE barang SET nama = $1, harga = $2, tanggal_beli = $3, kategori_id = $4, total_penggunaan = $5 WHERE id = $6",
		barang.Nama, barang.Harga, barang.TanggalBeli, barang.KategoriID, barang.TotalPenggunaan, barang.ID)
	return err
}

func (r *BarangRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM barang WHERE id = $1", id)
	return err
}

func (r *BarangRepository) SearchByName(keyword string) ([]model.Barang, error) {
	rows, err := r.DB.Query("SELECT id, nama, harga, tanggal_beli, kategori_id FROM barang WHERE LOWER(nama) LIKE LOWER($1)", "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var barangs []model.Barang
	for rows.Next() {
		var b model.Barang
		if err := rows.Scan(&b.ID, &b.Nama, &b.Harga, &b.TanggalBeli, &b.KategoriID); err != nil {
			return nil, err
		}
		barangs = append(barangs, b)
	}
	return barangs, nil
}
