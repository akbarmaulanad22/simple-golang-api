package repository

import (
	"api/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
	Create(user *entity.User) error
	Update(id uint, user *entity.User) error
	Delete(id uint) error
	FindById(id uint) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
// 	var user entity.User
//     if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
//         return nil, err
//     }
//     return &user, nil
// }

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindById(id uint) (*entity.User, error) {
	var user entity.User
    if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
	
	//  // Inisialisasi variabel user
    // var user entity.User

    // // Query untuk mencari user berdasarkan ID
    // result := r.db.First(&user, id)

    // if result.Error == nil {
    // 	return user, nil
    // }

	// // Jika record tidak ditemukan
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return user, fmt.Errorf("user tidak ditemukan")
	// }
	
	// // Untuk kesalahan lainnya
	// return user, fmt.Errorf("terjadi kesalahan: %v", result.Error)
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
