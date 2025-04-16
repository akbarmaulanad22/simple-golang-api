package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type BeasiswaUsecase interface {
	GetAllBeasiswas() ([]entity.Beasiswa, error)
	CreateBeasiswa(beasiswa *entity.Beasiswa) error
	UpdateBeasiswa(id uint, beasiswa *entity.Beasiswa) error
	DeleteBeasiswa(id uint) error
}

type beasiswaUsecase struct {
	beasiswaRepo repository.BeasiswaRepository
}

func NewBeasiswaUsecase(db *gorm.DB) BeasiswaUsecase {
	return &beasiswaUsecase{
		beasiswaRepo: repository.NewBeasiswaRepository(db),
	}
}

func (u *beasiswaUsecase) GetAllBeasiswas() ([]entity.Beasiswa, error) {
	return u.beasiswaRepo.FindAll()
}

func (u *beasiswaUsecase) CreateBeasiswa(beasiswa *entity.Beasiswa) error {
	return u.beasiswaRepo.Create(beasiswa)
}

func (u *beasiswaUsecase) UpdateBeasiswa(id uint, beasiswa *entity.Beasiswa) error {
	return u.beasiswaRepo.Update(id, beasiswa)
}

func (u *beasiswaUsecase) DeleteBeasiswa(id uint) error {
	return u.beasiswaRepo.Delete(id)
}
