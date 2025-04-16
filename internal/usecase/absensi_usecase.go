package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type AbsensiUsecase interface {
	GetAllAbsensis() ([]entity.Absensi, error)
	CreateAbsensi(absensi *entity.Absensi) error
	UpdateAbsensi(id uint, absensi *entity.Absensi) error
	DeleteAbsensi(id uint) error
}

type absensiUsecase struct {
	absensiRepo repository.AbsensiRepository
}

func NewAbsensiUsecase(db *gorm.DB) AbsensiUsecase {
	return &absensiUsecase{
		absensiRepo: repository.NewAbsensiRepository(db),
	}
}

func (u *absensiUsecase) GetAllAbsensis() ([]entity.Absensi, error) {
	return u.absensiRepo.FindAll()
}

func (u *absensiUsecase) CreateAbsensi(absensi *entity.Absensi) error {
	return u.absensiRepo.Create(absensi)
}

func (u *absensiUsecase) UpdateAbsensi(id uint, absensi *entity.Absensi) error {
	return u.absensiRepo.Update(id, absensi)
}

func (u *absensiUsecase) DeleteAbsensi(id uint) error {
	return u.absensiRepo.Delete(id)
}
