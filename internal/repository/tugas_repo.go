package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type TugasRepository interface {
	FindAll() ([]entity.Tugas, error)
	Create(tugas *entity.Tugas) error
	Update(id uint, tugas *entity.Tugas) error
	Delete(id uint) error
}

type tugasRepository struct {
	db *gorm.DB
}

func NewTugasRepository(db *gorm.DB) TugasRepository {
	return &tugasRepository{db}
}

func (r *tugasRepository) FindAll() ([]entity.Tugas, error) {
	var tugass []entity.Tugas
	if err := r.db.Find(&tugass).Error; err != nil {
		return nil, err
	}
	return tugass, nil
}

func (r *tugasRepository) Create(tugas *entity.Tugas) error {
	return r.db.Create(tugas).Error
}

func (r *tugasRepository) Update(id uint, tugas *entity.Tugas) error {
	return r.db.Model(&entity.Tugas{}).Where("id = ?", id).Updates(tugas).Error
}

func (r *tugasRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Tugas{}, id).Error
}
