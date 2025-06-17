package repository

import (
    "database/sql"
    "project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
)

type BarangRepository struct {
    DB *sql.DB
}

func (r *BarangRepository) GetAll() ([]model.Barang, error) {
    rows, err := r.DB.Query("SELECT id, nama, harga, tanggal_beli, kategori_id FROM barang")
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

func (r *BarangRepository) GetByID(id int) (*model.Barang, error) {
    row := r.DB.QueryRow("SELECT id, nama, harga, tanggal_beli, kategori_id FROM barang WHERE id=$1", id)
    var b model.Barang
    if err := row.Scan(&b.ID, &b.Nama, &b.Harga, &b.TanggalBeli, &b.KategoriID); err != nil {
        return nil, err
    }
    return &b, nil
}

func (r *BarangRepository) Create(nama string, harga float64, tanggalBeli string, kategoriID int) error {
    _, err := r.DB.Exec("INSERT INTO barang (nama, harga, tanggal_beli, kategori_id) VALUES ($1, $2, $3, $4)", nama, harga, tanggalBeli, kategoriID)
    return err
}

func (r *BarangRepository) Update(id int, nama string, harga float64, tanggalBeli string, kategoriID int) error {
    _, err := r.DB.Exec("UPDATE barang SET nama=$1, harga=$2, tanggal_beli=$3, kategori_id=$4 WHERE id=$5", nama, harga, tanggalBeli, kategoriID, id)
    return err
}

func (r *BarangRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM barang WHERE id=$1", id)
    return err
}