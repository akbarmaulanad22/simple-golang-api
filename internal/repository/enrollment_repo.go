package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type EnrollmentRepository interface {
	FindAll() ([]entity.Enrollment, error)
	Create(enrollment *entity.Enrollment) error
	Update(id uint, enrollment *entity.Enrollment) error
	Delete(id uint) error
	FindById(id uint) (entity.Enrollment, error)
}

type enrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) EnrollmentRepository {
	return &enrollmentRepository{db}
}

func (r *enrollmentRepository) FindAll() ([]entity.Enrollment, error) {
	var enrollments []entity.Enrollment
	if err := r.db.Preload("CreatedByUser").Find(&enrollments).Error; err != nil {
		return nil, err
	}
	return enrollments, nil
}

func (r *enrollmentRepository) FindById(id uint) (entity.Enrollment, error) {
	 // Inisialisasi variabel enrollment
    var enrollment entity.Enrollment

    // Query untuk mencari enrollment berdasarkan ID
    result := r.db.Preload("CreatedByUser").First(&enrollment, id)

    if result.Error == nil {
    	return enrollment, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return enrollment, fmt.Errorf("enrollment tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return enrollment, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *enrollmentRepository) Create(enrollment *entity.Enrollment) error {
	return r.db.Create(enrollment).Error
}

func (r *enrollmentRepository) Update(id uint, enrollment *entity.Enrollment) error {
	return r.db.Model(&entity.Enrollment{}).Where("id = ?", id).Updates(enrollment).Error
}

func (r *enrollmentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Enrollment{}, id).Error
}
