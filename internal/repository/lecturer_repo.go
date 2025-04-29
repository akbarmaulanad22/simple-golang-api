package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type LecturerRepository interface {
	FindAll() ([]entity.Lecturer, error)
	Create(lecture *entity.Lecturer) error
	Update(id uint, lecture *entity.Lecturer) error
	Delete(id uint) error
	FindById(id uint) (entity.Lecturer, error)
}

type lectureRepository struct {
	db *gorm.DB
}

func NewLecturerRepository(db *gorm.DB) LecturerRepository {
	return &lectureRepository{db}
}

func (r *lectureRepository) FindAll() ([]entity.Lecturer, error) {
	var lectures []entity.Lecturer
	if err := r.db.Find(&lectures).Error; err != nil {
		return nil, err
	}
	return lectures, nil
}

func (r *lectureRepository) FindById(id uint) (entity.Lecturer, error) {
	 // Inisialisasi variabel lecture
    var lecture entity.Lecturer

    // Query untuk mencari lecture berdasarkan ID
    result := r.db.First(&lecture, id)

    if result.Error == nil {
    	return lecture, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return lecture, fmt.Errorf("lecture tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return lecture, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *lectureRepository) Create(lecture *entity.Lecturer) error {
	return r.db.Create(lecture).Error
}

func (r *lectureRepository) Update(id uint, lecture *entity.Lecturer) error {
	return r.db.Model(&entity.Lecturer{}).Where("id = ?", id).Updates(lecture).Error
}

func (r *lectureRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Lecturer{}, id).Error
}
