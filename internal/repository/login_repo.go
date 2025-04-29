package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type LoginRepository struct {
    DB *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
    return &LoginRepository{db}
}

func (r *LoginRepository) FindByUsername(username string) (*entity.User, error) {
    var user entity.User
    if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}