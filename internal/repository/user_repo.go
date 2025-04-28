package repository

import (
	"api/internal/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	Create(user *entity.User) error
	Update(id uint, user *entity.User) error
	Delete(id uint) error
	FindById(id uint) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindById(id uint) (entity.User, error) {
	 // Inisialisasi variabel user
    var user entity.User

    // Query untuk mencari user berdasarkan ID
    result := r.db.First(&user, id)

    // Handle error
    if result.Error != nil {
        // Jika record tidak ditemukan
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return user, fmt.Errorf("user tidak ditemukan")
        }
        // Untuk kesalahan lainnya
        return user, fmt.Errorf("terjadi kesalahan: %v", result.Error)
    }

    // Jika berhasil, kembalikan user tanpa error
    return user, nil
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(id uint, user *entity.User) error {
	return r.db.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
