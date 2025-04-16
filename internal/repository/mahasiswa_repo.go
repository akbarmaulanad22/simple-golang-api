package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	FindAll() ([]entity.Mahasiswa, error)
	Create(mahasiswa *entity.Mahasiswa) error
	Update(id uint, mahasiswa *entity.Mahasiswa) error
	Delete(id uint) error
}

type mahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{db}
}

func (r *mahasiswaRepository) FindAll() ([]entity.Mahasiswa, error) {
	var mahasiswas []entity.Mahasiswa
	if err := r.db.Find(&mahasiswas).Error; err != nil {
		return nil, err
	}
	return mahasiswas, nil
}

func (r *mahasiswaRepository) Create(mahasiswa *entity.Mahasiswa) error {
	return r.db.Create(mahasiswa).Error
}

func (r *mahasiswaRepository) Update(id uint, mahasiswa *entity.Mahasiswa) error {
	return r.db.Model(&entity.Mahasiswa{}).Where("id = ?", id).Updates(mahasiswa).Error
}

func (r *mahasiswaRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Mahasiswa{}, id).Error
}
