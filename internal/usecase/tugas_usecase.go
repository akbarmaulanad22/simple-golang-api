package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type TugasUsecase interface {
	GetAllTugass() ([]entity.Tugas, error)
	CreateTugas(tugas *entity.Tugas) error
	UpdateTugas(id uint, tugas *entity.Tugas) error
	DeleteTugas(id uint) error
}

type tugasUsecase struct {
	tugasRepo repository.TugasRepository
}

func NewTugasUsecase(db *gorm.DB) TugasUsecase {
	return &tugasUsecase{
		tugasRepo: repository.NewTugasRepository(db),
	}
}

func (u *tugasUsecase) GetAllTugass() ([]entity.Tugas, error) {
	return u.tugasRepo.FindAll()
}

func (u *tugasUsecase) CreateTugas(tugas *entity.Tugas) error {
	return u.tugasRepo.Create(tugas)
}

func (u *tugasUsecase) UpdateTugas(id uint, tugas *entity.Tugas) error {
	return u.tugasRepo.Update(id, tugas)
}

func (u *tugasUsecase) DeleteTugas(id uint) error {
	return u.tugasRepo.Delete(id)
}
