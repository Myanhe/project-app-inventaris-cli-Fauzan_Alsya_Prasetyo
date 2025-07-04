package service

import (
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"	
)

type KategoriService struct {
	Repo *repository.KategoriRepository
}

func (s *KategoriService) GetAll() ([]model.Kategori, error) {
	return s.Repo.GetAll()
}

func (s *KategoriService) GetByID(id int) (*model.Kategori, error) {
	return s.Repo.GetByID(id)
}

func (s *KategoriService) Create(nama, deskripsi string) error {
	kategori := &model.Kategori{
		Nama:      nama,
		Deskripsi: deskripsi,
	}
	return s.Repo.Create(kategori)
}

func (s *KategoriService) Update(id int, nama, deskripsi string) error {
	kategori := &model.Kategori{
		ID:        id,
		Nama:      nama,
		Deskripsi: deskripsi,
	}
	return s.Repo.Update(kategori)
}

func (s *KategoriService) Delete(id int) error {
	return s.Repo.Delete(id)
}
