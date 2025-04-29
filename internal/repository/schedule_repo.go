package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	FindAll() ([]entity.Schedule, error)
	Create(schedule *entity.Schedule) error
	Update(id uint, schedule *entity.Schedule) error
	Delete(id uint) error
	FindById(id uint) (entity.Schedule, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db}
}

func (r *scheduleRepository) FindAll() ([]entity.Schedule, error) {
	var schedules []entity.Schedule
	if err := r.db.Preload("CreatedByUser").Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) FindById(id uint) (entity.Schedule, error) {
	 // Inisialisasi variabel schedule
    var schedule entity.Schedule

    // Query untuk mencari schedule berdasarkan ID
    result := r.db.Preload("CreatedByUser").First(&schedule, id)

    if result.Error == nil {
    	return schedule, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return schedule, fmt.Errorf("schedule tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return schedule, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *scheduleRepository) Create(schedule *entity.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) Update(id uint, schedule *entity.Schedule) error {
	return r.db.Model(&entity.Schedule{}).Where("id = ?", id).Updates(schedule).Error
}

func (r *scheduleRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Schedule{}, id).Error
}
