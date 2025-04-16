package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type MataKuliahUsecase interface {
	GetAllMataKuliahs() ([]entity.MataKuliah, error)
	CreateMataKuliah(mataKuliah *entity.MataKuliah) error
	UpdateMataKuliah(id uint, mataKuliah *entity.MataKuliah) error
	DeleteMataKuliah(id uint) error
}

type mataKuliahUsecase struct {
	mataKuliahRepo repository.MataKuliahRepository
}

func NewMataKuliahUsecase(db *gorm.DB) MataKuliahUsecase {
	return &mataKuliahUsecase{
		mataKuliahRepo: repository.NewMataKuliahRepository(db),
	}
}

func (u *mataKuliahUsecase) GetAllMataKuliahs() ([]entity.MataKuliah, error) {
	return u.mataKuliahRepo.FindAll()
}

func (u *mataKuliahUsecase) CreateMataKuliah(mataKuliah *entity.MataKuliah) error {
	return u.mataKuliahRepo.Create(mataKuliah)
}

func (u *mataKuliahUsecase) UpdateMataKuliah(id uint, mataKuliah *entity.MataKuliah) error {
	return u.mataKuliahRepo.Update(id, mataKuliah)
}

func (u *mataKuliahUsecase) DeleteMataKuliah(id uint) error {
	return u.mataKuliahRepo.Delete(id)
}
