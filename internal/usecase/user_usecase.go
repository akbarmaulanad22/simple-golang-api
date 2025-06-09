package usecase

import (
	"api/internal/entity"
	"api/internal/helper"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type UserUsecase interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(user *entity.User, ctx context.Context) error
	UpdateUser(id uint, user *entity.User, ctx context.Context) error
	DeleteUser(id uint, ctx context.Context) error
	FindByIdUser(id uint) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	logRepo  repository.LogRepository
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &userUsecase{
		userRepo: repository.NewUserRepository(db),
		logRepo:  repository.NewLogRepository(db),
	}
}

func (u *userUsecase) FindByIdUser(id uint) (*entity.User, error) {

	return u.userRepo.FindById(id)

}

func (u *userUsecase) GetAllUsers() ([]entity.User, error) {
	return u.userRepo.FindAll()
}

func (u *userUsecase) CreateUser(user *entity.User, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
	if !ok {
		return errors.New("user not authenticated")
	}

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
		UserID:     userClaims.ID,
		Action:     "CREATE",
		EntityType: "USER",
		EntityID:   user.ID,
		Details:    string(jsonBytes),
	})

	if errCreateLog != nil {
		return errCreateLog
	}

	return nil
}

func (u *userUsecase) UpdateUser(id uint, user *entity.User, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
	if !ok {
		return errors.New("user not authenticated")
	}
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
		UserID:     userClaims.ID,
		Action:     "UPDATE",
		EntityType: "USER",
		EntityID:   id,
		Details:    string(jsonBytes),
	})

	if errCreateLog != nil {
		return errCreateLog
	}

	return nil
}

func (u *userUsecase) DeleteUser(id uint, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
	if !ok {
		return errors.New("user not authenticated")
	}
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID:     userClaims.ID,
		Action:     "DELETE",
		EntityType: "USER",
		EntityID:   id,
	})

	if errCreateLog != nil {
		return errCreateLog
	}

	return u.userRepo.Delete(id)

}
