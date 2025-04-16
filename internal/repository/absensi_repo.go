package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type AbsensiRepository interface {
	FindAll() ([]entity.Absensi, error)
	Create(absensi *entity.Absensi) error
	Update(id uint, absensi *entity.Absensi) error
	Delete(id uint) error
}

type absensiRepository struct {
	db *gorm.DB
}

func NewAbsensiRepository(db *gorm.DB) AbsensiRepository {
	return &absensiRepository{db}
}

func (r *absensiRepository) FindAll() ([]entity.Absensi, error) {
	var absensis []entity.Absensi
	if err := r.db.Find(&absensis).Error; err != nil {
		return nil, err
	}
	return absensis, nil
}

func (r *absensiRepository) Create(absensi *entity.Absensi) error {
	return r.db.Create(absensi).Error
}

func (r *absensiRepository) Update(id uint, absensi *entity.Absensi) error {
	return r.db.Model(&entity.Absensi{}).Where("id = ?", id).Updates(absensi).Error
}

func (r *absensiRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Absensi{}, id).Error
}
