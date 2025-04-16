package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type DosenRepository interface {
	FindAll() ([]entity.Dosen, error)
	Create(dosen *entity.Dosen) error
	Update(id uint, dosen *entity.Dosen) error
	Delete(id uint) error
}

type dosenRepository struct {
	db *gorm.DB
}

func NewDosenRepository(db *gorm.DB) DosenRepository {
	return &dosenRepository{db}
}

func (r *dosenRepository) FindAll() ([]entity.Dosen, error) {
	var dosens []entity.Dosen
	if err := r.db.Find(&dosens).Error; err != nil {
		return nil, err
	}
	return dosens, nil
}

func (r *dosenRepository) Create(dosen *entity.Dosen) error {
	return r.db.Create(dosen).Error
}

func (r *dosenRepository) Update(id uint, dosen *entity.Dosen) error {
	return r.db.Model(&entity.Dosen{}).Where("id = ?", id).Updates(dosen).Error
}

func (r *dosenRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Dosen{}, id).Error
}
