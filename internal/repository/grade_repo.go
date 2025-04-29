package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type GradeRepository interface {
	FindAll() ([]entity.Grade, error)
	Create(grade *entity.Grade) error
	Update(id uint, grade *entity.Grade) error
	Delete(id uint) error
	FindById(id uint) (entity.Grade, error)
}

type gradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository(db *gorm.DB) GradeRepository {
	return &gradeRepository{db}
}

func (r *gradeRepository) FindAll() ([]entity.Grade, error) {
	var grades []entity.Grade
	if err := r.db.Preload("CreatedByUser").Find(&grades).Error; err != nil {
		return nil, err
	}
	return grades, nil
}

func (r *gradeRepository) FindById(id uint) (entity.Grade, error) {
	 // Inisialisasi variabel grade
    var grade entity.Grade

    // Query untuk mencari grade berdasarkan ID
    result := r.db.Preload("CreatedByUser").First(&grade, id)

    if result.Error == nil {
    	return grade, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return grade, fmt.Errorf("grade tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return grade, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *gradeRepository) Create(grade *entity.Grade) error {
	return r.db.Create(grade).Error
}

func (r *gradeRepository) Update(id uint, grade *entity.Grade) error {
	return r.db.Model(&entity.Grade{}).Where("id = ?", id).Updates(grade).Error
}

func (r *gradeRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Grade{}, id).Error
}
