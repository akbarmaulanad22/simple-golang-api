package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type NilaiRepository interface {
	FindAll() ([]entity.Nilai, error)
	Create(nilai *entity.Nilai) error
	Update(id uint, nilai *entity.Nilai) error
	Delete(id uint) error
}

type nilaiRepository struct {
	db *gorm.DB
}

func NewNilaiRepository(db *gorm.DB) NilaiRepository {
	return &nilaiRepository{db}
}

func (r *nilaiRepository) FindAll() ([]entity.Nilai, error) {
	var nilais []entity.Nilai
	if err := r.db.Find(&nilais).Error; err != nil {
		return nil, err
	}
	return nilais, nil
}

func (r *nilaiRepository) Create(nilai *entity.Nilai) error {
	return r.db.Create(nilai).Error
}

func (r *nilaiRepository) Update(id uint, nilai *entity.Nilai) error {
	return r.db.Model(&entity.Nilai{}).Where("id = ?", id).Updates(nilai).Error
}

func (r *nilaiRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Nilai{}, id).Error
}
