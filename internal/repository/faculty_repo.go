package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type FacultyRepository interface {
	FindAll() ([]entity.Faculty, error)
	Create(faculty *entity.Faculty) error
	Update(id uint, faculty *entity.Faculty) error
	Delete(id uint) error
	FindById(id uint) (entity.Faculty, error)
}

type facultyRepository struct {
	db *gorm.DB
}

func NewFacultyRepository(db *gorm.DB) FacultyRepository {
	return &facultyRepository{db}
}

func (r *facultyRepository) FindAll() ([]entity.Faculty, error) {
	var facultys []entity.Faculty
	if err := r.db.Preload("CreatedByUser").Find(&facultys).Error; err != nil {
		return nil, err
	}
	return facultys, nil
}

func (r *facultyRepository) FindById(id uint) (entity.Faculty, error) {
	 // Inisialisasi variabel faculty
    var faculty entity.Faculty

    // Query untuk mencari faculty berdasarkan ID
    result := r.db.Preload("CreatedByUser").First(&faculty, id)

    if result.Error == nil {
    	return faculty, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return faculty, fmt.Errorf("faculty tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return faculty, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *facultyRepository) Create(faculty *entity.Faculty) error {
	return r.db.Create(faculty).Error
}

func (r *facultyRepository) Update(id uint, faculty *entity.Faculty) error {
	return r.db.Model(&entity.Faculty{}).Where("id = ?", id).Updates(faculty).Error
}

func (r *facultyRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Faculty{}, id).Error
}
