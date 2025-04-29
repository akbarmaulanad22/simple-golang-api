package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	FindAll() ([]entity.Attendance, error)
	Create(attendance *entity.Attendance) error
	Update(id uint, attendance *entity.Attendance) error
	Delete(id uint) error
	FindById(id uint) (entity.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db}
}

func (r *attendanceRepository) FindAll() ([]entity.Attendance, error) {
	var attendances []entity.Attendance
	if err := r.db.Preload("Student").Preload("Schedule").Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}

func (r *attendanceRepository) FindById(id uint) (entity.Attendance, error) {
	 // Inisialisasi variabel attendance
    var attendance entity.Attendance

    // Query untuk mencari attendance berdasarkan ID
    result := r.db.Preload("Student").Preload("Schedule").First(&attendance, id)

    if result.Error == nil {
    	return attendance, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return attendance, fmt.Errorf("attendance tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return attendance, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *attendanceRepository) Create(attendance *entity.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *attendanceRepository) Update(id uint, attendance *entity.Attendance) error {
	return r.db.Model(&entity.Attendance{}).Where("id = ?", id).Updates(attendance).Error
}

func (r *attendanceRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Attendance{}, id).Error
}
