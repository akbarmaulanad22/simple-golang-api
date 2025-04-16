package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type PengumumanUsecase interface {
	GetAllPengumumans() ([]entity.Pengumuman, error)
	CreatePengumuman(pengumuman *entity.Pengumuman) error
	UpdatePengumuman(id uint, pengumuman *entity.Pengumuman) error
	DeletePengumuman(id uint) error
}

type pengumumanUsecase struct {
	pengumumanRepo repository.PengumumanRepository
}

func NewPengumumanUsecase(db *gorm.DB) PengumumanUsecase {
	return &pengumumanUsecase{
		pengumumanRepo: repository.NewPengumumanRepository(db),
	}
}

func (u *pengumumanUsecase) GetAllPengumumans() ([]entity.Pengumuman, error) {
	return u.pengumumanRepo.FindAll()
}

func (u *pengumumanUsecase) CreatePengumuman(pengumuman *entity.Pengumuman) error {
	return u.pengumumanRepo.Create(pengumuman)
}

func (u *pengumumanUsecase) UpdatePengumuman(id uint, pengumuman *entity.Pengumuman) error {
	return u.pengumumanRepo.Update(id, pengumuman)
}

func (u *pengumumanUsecase) DeletePengumuman(id uint) error {
	return u.pengumumanRepo.Delete(id)
}
