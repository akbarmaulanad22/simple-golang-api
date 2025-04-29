package usecase

import (
	"api/internal/entity"
	"api/internal/helper"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type UserUsecase interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(id uint, user *entity.User) error
	DeleteUser(id uint) error
	FindByIdUser(id uint) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	logRepo repository.LogRepository
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &userUsecase{
		userRepo: repository.NewUserRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *userUsecase) FindByIdUser(id uint) (*entity.User, error) {
	
	return u.userRepo.FindById(id)
	
}

func (u *userUsecase) GetAllUsers() ([]entity.User, error) {
	return u.userRepo.FindAll()
}

func (u *userUsecase) CreateUser(user *entity.User) error {
	// currentUser, err := osUser.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentUser.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	password, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	
	user.Password = string(password)
	
	errCreate := u.userRepo.Create(user)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// UserID: id,
		Action: "CREATE",
		EntityType: "USER",
		EntityID: user.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *userUsecase) UpdateUser(id uint, user *entity.User) error {
	password, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	
	user.ID = id
	user.Password = string(password)

	errUpdate := u.userRepo.Update(id, user)
	if errUpdate != nil {
		return errUpdate
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// UserID: id,
		Action: "UPDATE",
		EntityType: "USER",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *userUsecase) DeleteUser(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// UserID: id,
		Action: "DELETE",
		EntityType: "USER",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.userRepo.Delete(id)

}
