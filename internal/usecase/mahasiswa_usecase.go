package usecase

import (
	"api/internal/entity"
	"api/internal/repository"

	"gorm.io/gorm"
)

type MahasiswaUsecase interface {
	GetAllMahasiswas() ([]entity.Mahasiswa, error)
	CreateMahasiswa(mahasiswa *entity.Mahasiswa) error
	UpdateMahasiswa(id uint, mahasiswa *entity.Mahasiswa) error
	DeleteMahasiswa(id uint) error
}

type mahasiswaUsecase struct {
	mahasiswaRepo repository.MahasiswaRepository
}

func NewMahasiswaUsecase(db *gorm.DB) MahasiswaUsecase {
	return &mahasiswaUsecase{
		mahasiswaRepo: repository.NewMahasiswaRepository(db),
	}
}

func (u *mahasiswaUsecase) GetAllMahasiswas() ([]entity.Mahasiswa, error) {
	return u.mahasiswaRepo.FindAll()
}

func (u *mahasiswaUsecase) CreateMahasiswa(mahasiswa *entity.Mahasiswa) error {
	return u.mahasiswaRepo.Create(mahasiswa)
}

func (u *mahasiswaUsecase) UpdateMahasiswa(id uint, mahasiswa *entity.Mahasiswa) error {
	return u.mahasiswaRepo.Update(id, mahasiswa)
}

func (u *mahasiswaUsecase) DeleteMahasiswa(id uint) error {
	return u.mahasiswaRepo.Delete(id)
}
