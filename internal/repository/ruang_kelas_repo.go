package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type RuangKelasRepository interface {
	FindAll() ([]entity.RuangKelas, error)
	Create(ruangKelas *entity.RuangKelas) error
	Update(id uint, ruangKelas *entity.RuangKelas) error
	Delete(id uint) error
}

type ruangKelasRepository struct {
	db *gorm.DB
}

func NewRuangKelasRepository(db *gorm.DB) RuangKelasRepository {
	return &ruangKelasRepository{db}
}

func (r *ruangKelasRepository) FindAll() ([]entity.RuangKelas, error) {
	var ruangKelass []entity.RuangKelas
	if err := r.db.Find(&ruangKelass).Error; err != nil {
		return nil, err
	}
	return ruangKelass, nil
}

func (r *ruangKelasRepository) Create(ruangKelas *entity.RuangKelas) error {
	return r.db.Create(ruangKelas).Error
}

func (r *ruangKelasRepository) Update(id uint, ruangKelas *entity.RuangKelas) error {
	return r.db.Model(&entity.RuangKelas{}).Where("id = ?", id).Updates(ruangKelas).Error
}

func (r *ruangKelasRepository) Delete(id uint) error {
	return r.db.Delete(&entity.RuangKelas{}, id).Error
}
