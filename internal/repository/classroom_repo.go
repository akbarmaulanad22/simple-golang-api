package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ClassroomRepository interface {
	FindAll() ([]entity.Classroom, error)
	Create(classroom *entity.Classroom) error
	Update(id uint, classroom *entity.Classroom) error
	Delete(id uint) error
	FindById(id uint) (entity.Classroom, error)
}

type classroomRepository struct {
	db *gorm.DB
}

func NewClassroomRepository(db *gorm.DB) ClassroomRepository {
	return &classroomRepository{db}
}

func (r *classroomRepository) FindAll() ([]entity.Classroom, error) {
	var classrooms []entity.Classroom
	if err := r.db.Preload("CreatedByUser").Find(&classrooms).Error; err != nil {
		return nil, err
	}
	return classrooms, nil
}

func (r *classroomRepository) FindById(id uint) (entity.Classroom, error) {
	 // Inisialisasi variabel classroom
    var classroom entity.Classroom

    // Query untuk mencari classroom berdasarkan ID
    result := r.db.Preload("CreatedByUser").First(&classroom, id)

    if result.Error == nil {
    	return classroom, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return classroom, fmt.Errorf("classroom tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return classroom, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *classroomRepository) Create(classroom *entity.Classroom) error {
	return r.db.Create(classroom).Error
}

func (r *classroomRepository) Update(id uint, classroom *entity.Classroom) error {
	return r.db.Model(&entity.Classroom{}).Where("id = ?", id).Updates(classroom).Error
}

func (r *classroomRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Classroom{}, id).Error
}
