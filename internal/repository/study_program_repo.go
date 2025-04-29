package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type StudyProgramRepository interface {
	FindAll() ([]entity.StudyProgram, error)
	Create(studyProgram *entity.StudyProgram) error
	Update(id uint, studyProgram *entity.StudyProgram) error
	Delete(id uint) error
	FindById(id uint) (entity.StudyProgram, error)
}

type studyProgramRepository struct {
	db *gorm.DB
}

func NewStudyProgramRepository(db *gorm.DB) StudyProgramRepository {
	return &studyProgramRepository{db}
}

func (r *studyProgramRepository) FindAll() ([]entity.StudyProgram, error) {
	var studyPrograms []entity.StudyProgram
	if err := r.db.Preload("Faculty").Find(&studyPrograms).Error; err != nil {
		return nil, err
	}
	return studyPrograms, nil
}

func (r *studyProgramRepository) FindById(id uint) (entity.StudyProgram, error) {
	 // Inisialisasi variabel studyProgram
    var studyProgram entity.StudyProgram

    // Query untuk mencari studyProgram berdasarkan ID
    result := r.db.Preload("Faculty").First(&studyProgram, id)

    if result.Error == nil {
    	return studyProgram, nil
    }

	// Jika record tidak ditemukan
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return studyProgram, fmt.Errorf("studyProgram tidak ditemukan")
	}
	
	// Untuk kesalahan lainnya
	return studyProgram, fmt.Errorf("terjadi kesalahan: %v", result.Error)
}

func (r *studyProgramRepository) Create(studyProgram *entity.StudyProgram) error {
	return r.db.Create(studyProgram).Error
}

func (r *studyProgramRepository) Update(id uint, studyProgram *entity.StudyProgram) error {
	return r.db.Model(&entity.StudyProgram{}).Where("id = ?", id).Updates(studyProgram).Error
}

func (r *studyProgramRepository) Delete(id uint) error {
	return r.db.Delete(&entity.StudyProgram{}, id).Error
}
