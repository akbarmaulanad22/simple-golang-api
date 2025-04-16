package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type PengumumanRepository interface {
	FindAll() ([]entity.Pengumuman, error)
	Create(pengumuman *entity.Pengumuman) error
	Update(id uint, pengumuman *entity.Pengumuman) error
	Delete(id uint) error
}

type pengumumanRepository struct {
	db *gorm.DB
}

func NewPengumumanRepository(db *gorm.DB) PengumumanRepository {
	return &pengumumanRepository{db}
}

func (r *pengumumanRepository) FindAll() ([]entity.Pengumuman, error) {
	var pengumumans []entity.Pengumuman
	if err := r.db.Find(&pengumumans).Error; err != nil {
		return nil, err
	}
	return pengumumans, nil
}

func (r *pengumumanRepository) Create(pengumuman *entity.Pengumuman) error {
	return r.db.Create(pengumuman).Error
}

func (r *pengumumanRepository) Update(id uint, pengumuman *entity.Pengumuman) error {
	return r.db.Model(&entity.Pengumuman{}).Where("id = ?", id).Updates(pengumuman).Error
}

func (r *pengumumanRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Pengumuman{}, id).Error
}
