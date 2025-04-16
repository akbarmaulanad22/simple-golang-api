package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type BeasiswaRepository interface {
	FindAll() ([]entity.Beasiswa, error)
	Create(beasiswa *entity.Beasiswa) error
	Update(id uint, beasiswa *entity.Beasiswa) error
	Delete(id uint) error
}

type beasiswaRepository struct {
	db *gorm.DB
}

func NewBeasiswaRepository(db *gorm.DB) BeasiswaRepository {
	return &beasiswaRepository{db}
}

func (r *beasiswaRepository) FindAll() ([]entity.Beasiswa, error) {
	var beasiswas []entity.Beasiswa
	if err := r.db.Find(&beasiswas).Error; err != nil {
		return nil, err
	}
	return beasiswas, nil
}

func (r *beasiswaRepository) Create(beasiswa *entity.Beasiswa) error {
	return r.db.Create(beasiswa).Error
}

func (r *beasiswaRepository) Update(id uint, beasiswa *entity.Beasiswa) error {
	return r.db.Model(&entity.Beasiswa{}).Where("id = ?", id).Updates(beasiswa).Error
}

func (r *beasiswaRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Beasiswa{}, id).Error
}
