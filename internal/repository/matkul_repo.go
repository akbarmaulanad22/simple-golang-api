package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	FindAll() ([]entity.MataKuliah, error)
	Create(mataKuliah *entity.MataKuliah) error
	Update(id uint, mataKuliah *entity.MataKuliah) error
	Delete(id uint) error
}

type mataKuliahRepository struct {
	db *gorm.DB
}

func NewMataKuliahRepository(db *gorm.DB) MataKuliahRepository {
	return &mataKuliahRepository{db}
}

func (r *mataKuliahRepository) FindAll() ([]entity.MataKuliah, error) {
	var mataKuliahs []entity.MataKuliah
	if err := r.db.Find(&mataKuliahs).Error; err != nil {
		return nil, err
	}
	return mataKuliahs, nil
}

func (r *mataKuliahRepository) Create(mataKuliah *entity.MataKuliah) error {
	return r.db.Create(mataKuliah).Error
}

func (r *mataKuliahRepository) Update(id uint, mataKuliah *entity.MataKuliah) error {
	return r.db.Model(&entity.MataKuliah{}).Where("id = ?", id).Updates(mataKuliah).Error
}

func (r *mataKuliahRepository) Delete(id uint) error {
	return r.db.Delete(&entity.MataKuliah{}, id).Error
}
