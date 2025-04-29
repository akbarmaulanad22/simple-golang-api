package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll() ([]entity.Course, error)
	Create(course *entity.Course) error
	Update(id uint, course *entity.Course) error
	Delete(id uint) error
	FindById(id uint) (entity.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) FindAll() ([]entity.Course, error) {
	var courses []entity.Course
	if err := r.db.Preload("CreatedByUser").Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) FindById(id uint) (entity.Course, error) {
	 // Inisialisasi variabel course
    var course entity.Course

    // Query untuk mencari course berdasarkan ID
    result := r.db.Preload("CreatedByUser").First(&course, id)

    if result.Error == nil {
    	return course, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return course, fmt.Errorf("course tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return course, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *courseRepository) Create(course *entity.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) Update(id uint, course *entity.Course) error {
	return r.db.Model(&entity.Course{}).Where("id = ?", id).Updates(course).Error
}

func (r *courseRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Course{}, id).Error
}
