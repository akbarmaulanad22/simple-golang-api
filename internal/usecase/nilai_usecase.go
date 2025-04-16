package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type NilaiUsecase interface {
	GetAllNilais() ([]entity.Nilai, error)
	CreateNilai(nilai *entity.Nilai) error
	UpdateNilai(id uint, nilai *entity.Nilai) error
	DeleteNilai(id uint) error
}

type nilaiUsecase struct {
	nilaiRepo repository.NilaiRepository
}

func NewNilaiUsecase(db *gorm.DB) NilaiUsecase {
	return &nilaiUsecase{
		nilaiRepo: repository.NewNilaiRepository(db),
	}
}

func (u *nilaiUsecase) GetAllNilais() ([]entity.Nilai, error) {
	return u.nilaiRepo.FindAll()
}

func (u *nilaiUsecase) CreateNilai(nilai *entity.Nilai) error {
	return u.nilaiRepo.Create(nilai)
}

func (u *nilaiUsecase) UpdateNilai(id uint, nilai *entity.Nilai) error {
	return u.nilaiRepo.Update(id, nilai)
}

func (u *nilaiUsecase) DeleteNilai(id uint) error {
	return u.nilaiRepo.Delete(id)
}
