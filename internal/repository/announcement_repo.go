package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	FindAll() ([]entity.Announcement, error)
	Create(announcement *entity.Announcement) error
	Update(id uint, announcement *entity.Announcement) error
	Delete(id uint) error
	FindById(id uint) (entity.Announcement, error)
}

type announcementRepository struct {
	db *gorm.DB
}

func NewAnnouncementRepository(db *gorm.DB) AnnouncementRepository {
	return &announcementRepository{db}
}

func (r *announcementRepository) FindAll() ([]entity.Announcement, error) {
	var announcements []entity.Announcement
	if err := r.db.Preload("CreatedByUser").Find(&announcements).Error; err != nil {
		return nil, err
	}
	return announcements, nil
}

func (r *announcementRepository) FindById(id uint) (entity.Announcement, error) {
	 // Inisialisasi variabel announcement
    var announcement entity.Announcement

    // Query untuk mencari announcement berdasarkan ID
    result := r.db.Preload("CreatedByUser").First(&announcement, id)

    if result.Error == nil {
    	return announcement, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return announcement, fmt.Errorf("announcement tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return announcement, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *announcementRepository) Create(announcement *entity.Announcement) error {
	return r.db.Create(announcement).Error
}

func (r *announcementRepository) Update(id uint, announcement *entity.Announcement) error {
	return r.db.Model(&entity.Announcement{}).Where("id = ?", id).Updates(announcement).Error
}

func (r *announcementRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Announcement{}, id).Error
}
