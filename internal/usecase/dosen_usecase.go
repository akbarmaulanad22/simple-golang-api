package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type DosenUsecase interface {
	GetAllDosens() ([]entity.Dosen, error)
	CreateDosen(dosen *entity.Dosen) error
	UpdateDosen(id uint, dosen *entity.Dosen) error
	DeleteDosen(id uint) error
}

type dosenUsecase struct {
	dosenRepo repository.DosenRepository
}

func NewDosenUsecase(db *gorm.DB) DosenUsecase {
	return &dosenUsecase{
		dosenRepo: repository.NewDosenRepository(db),
	}
}

func (u *dosenUsecase) GetAllDosens() ([]entity.Dosen, error) {
	return u.dosenRepo.FindAll()
}

func (u *dosenUsecase) CreateDosen(dosen *entity.Dosen) error {
	return u.dosenRepo.Create(dosen)
}

func (u *dosenUsecase) UpdateDosen(id uint, dosen *entity.Dosen) error {
	return u.dosenRepo.Update(id, dosen)
}

func (u *dosenUsecase) DeleteDosen(id uint) error {
	return u.dosenRepo.Delete(id)
}
