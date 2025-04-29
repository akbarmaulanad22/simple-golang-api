package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FindAll() ([]entity.Student, error)
	Create(student *entity.Student) error
	Update(id uint, student *entity.Student) error
	Delete(id uint) error
	FindById(id uint) (entity.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) FindAll() ([]entity.Student, error) {
	var students []entity.Student
	if err := r.db.Preload("User").Preload("StudyProgram").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) FindById(id uint) (entity.Student, error) {
	 // Inisialisasi variabel student
    var student entity.Student

    // Query untuk mencari student berdasarkan ID
    result := r.db.Preload("User").Preload("StudyProgram").First(&student, id)

    if result.Error == nil {
    	return student, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return student, fmt.Errorf("student tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return student, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *studentRepository) Create(student *entity.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) Update(id uint, student *entity.Student) error {
	return r.db.Model(&entity.Student{}).Where("id = ?", id).Updates(student).Error
}

func (r *studentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Student{}, id).Error
}
