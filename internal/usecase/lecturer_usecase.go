package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type LecturerUsecase interface {
	GetAllLecturers() ([]entity.Lecturer, error)
	CreateLecturer(lecture *entity.Lecturer, ctx context.Context) error
	UpdateLecturer(id uint, lecture *entity.Lecturer, ctx context.Context) error
	DeleteLecturer(id uint, ctx context.Context) error
	FindByIdLecturer(id uint) (entity.Lecturer, error)
}

type lectureUsecase struct {
	lectureRepo repository.LecturerRepository
	logRepo repository.LogRepository
}

func NewLecturerUsecase(db *gorm.DB) LecturerUsecase {
	return &lectureUsecase{
		lectureRepo: repository.NewLecturerRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *lectureUsecase) FindByIdLecturer(id uint) (entity.Lecturer, error) {
	
	return u.lectureRepo.FindById(id)
	
}

func (u *lectureUsecase) GetAllLecturers() ([]entity.Lecturer, error) {
	return u.lectureRepo.FindAll()
}

func (u *lectureUsecase) CreateLecturer(lecture *entity.Lecturer, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errCreate := u.lectureRepo.Create(lecture)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(lecture)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "CREATE",
		EntityType: "LECTURER",
		EntityID: lecture.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *lectureUsecase) UpdateLecturer(id uint, lecture *entity.Lecturer, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errUpdate := u.lectureRepo.Update(id, lecture)
	if errUpdate != nil {
		return errUpdate
	}

	lecture.ID = id

	jsonBytes, err := json.Marshal(lecture)
	if err != nil {
		return err
	}

	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "LECTURER",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *lectureUsecase) DeleteLecturer(id uint, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "DELETE",
		EntityType: "LECTURER",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.lectureRepo.Delete(id)

}
