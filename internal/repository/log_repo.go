package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type LogRepository interface {
	FindAll() ([]entity.Log, error)
	Create(log *entity.Log) error
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db}
}

func (r *logRepository) FindAll() ([]entity.Log, error) {
	var logs []entity.Log
	if err := r.db.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *logRepository) Create(log *entity.Log) error {
	return r.db.Create(log).Error
}