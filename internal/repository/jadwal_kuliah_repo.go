package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type JadwalKuliahRepository interface {
	FindAll() ([]entity.JadwalKuliah, error)
	Create(jadwalKuliah *entity.JadwalKuliah) error
	Update(id uint, jadwalKuliah *entity.JadwalKuliah) error
	Delete(id uint) error
}

type jadwalKuliahRepository struct {
	db *gorm.DB
}

func NewJadwalKuliahRepository(db *gorm.DB) JadwalKuliahRepository {
	return &jadwalKuliahRepository{db}
}

func (r *jadwalKuliahRepository) FindAll() ([]entity.JadwalKuliah, error) {
	var jadwalKuliahs []entity.JadwalKuliah
	if err := r.db.Find(&jadwalKuliahs).Error; err != nil {
		return nil, err
	}
	return jadwalKuliahs, nil
}

func (r *jadwalKuliahRepository) Create(jadwalKuliah *entity.JadwalKuliah) error {
	return r.db.Create(jadwalKuliah).Error
}

func (r *jadwalKuliahRepository) Update(id uint, jadwalKuliah *entity.JadwalKuliah) error {
	return r.db.Model(&entity.JadwalKuliah{}).Where("id = ?", id).Updates(jadwalKuliah).Error
}

func (r *jadwalKuliahRepository) Delete(id uint) error {
	return r.db.Delete(&entity.JadwalKuliah{}, id).Error
}
