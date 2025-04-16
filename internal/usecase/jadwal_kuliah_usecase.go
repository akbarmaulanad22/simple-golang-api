package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type JadwalKuliahUsecase interface {
	GetAllJadwalKuliahs() ([]entity.JadwalKuliah, error)
	CreateJadwalKuliah(jadwalKuliah *entity.JadwalKuliah) error
	UpdateJadwalKuliah(id uint, jadwalKuliah *entity.JadwalKuliah) error
	DeleteJadwalKuliah(id uint) error
}

type jadwalKuliahUsecase struct {
	jadwalKuliahRepo repository.JadwalKuliahRepository
}

func NewJadwalKuliahUsecase(db *gorm.DB) JadwalKuliahUsecase {
	return &jadwalKuliahUsecase{
		jadwalKuliahRepo: repository.NewJadwalKuliahRepository(db),
	}
}

func (u *jadwalKuliahUsecase) GetAllJadwalKuliahs() ([]entity.JadwalKuliah, error) {
	return u.jadwalKuliahRepo.FindAll()
}

func (u *jadwalKuliahUsecase) CreateJadwalKuliah(jadwalKuliah *entity.JadwalKuliah) error {
	return u.jadwalKuliahRepo.Create(jadwalKuliah)
}

func (u *jadwalKuliahUsecase) UpdateJadwalKuliah(id uint, jadwalKuliah *entity.JadwalKuliah) error {
	return u.jadwalKuliahRepo.Update(id, jadwalKuliah)
}

func (u *jadwalKuliahUsecase) DeleteJadwalKuliah(id uint) error {
	return u.jadwalKuliahRepo.Delete(id)
}
