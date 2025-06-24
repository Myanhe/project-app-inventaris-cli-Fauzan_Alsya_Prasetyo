package service

import (
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"
	"time"
)

type BarangService struct {
    Repo *repository.BarangRepository
}

func (s *BarangService) GetAll() ([]model.Barang, error) {
    return s.Repo.GetAll()
}

func (s *BarangService) GetByID(id int) (*model.Barang, error) {
    return s.Repo.GetByID(id)
}

func (s *BarangService) Create(nama string, harga float64, tanggalBeli time.Time, kategoriID int) error {
    barang := &model.Barang{
        Nama:        nama,
        Harga:       harga,
        TanggalBeli: tanggalBeli,
        KategoriID:  kategoriID,
    }
    return s.Repo.Create(barang)
}

func (s *BarangService) Update(id int, nama string, harga float64, tanggalBeli time.Time, kategoriID int) error {
    barang := &model.Barang{
        ID:          id,
        Nama:        nama,
        Harga:       harga,
        TanggalBeli: tanggalBeli,
        KategoriID:  kategoriID,
    }
    return s.Repo.Update(barang)
}

func (s *BarangService) Delete(id int) error {
    return s.Repo.Delete(id)
}