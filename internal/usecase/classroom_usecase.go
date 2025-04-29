package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type ClassroomUsecase interface {
	GetAllClassrooms() ([]entity.Classroom, error)
	CreateClassroom(classroom *entity.Classroom, ctx context.Context) error
	UpdateClassroom(id uint, classroom *entity.Classroom, ctx context.Context) error
	DeleteClassroom(id uint, ctx context.Context) error
	FindByIdClassroom(id uint) (entity.Classroom, error)
}

type classroomUsecase struct {
	classroomRepo repository.ClassroomRepository
	logRepo repository.LogRepository
}

func NewClassroomUsecase(db *gorm.DB) ClassroomUsecase {
	return &classroomUsecase{
		classroomRepo: repository.NewClassroomRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *classroomUsecase) FindByIdClassroom(id uint) (entity.Classroom, error) {
	
	return u.classroomRepo.FindById(id)
	
}

func (u *classroomUsecase) GetAllClassrooms() ([]entity.Classroom, error) {
	return u.classroomRepo.FindAll()
}

func (u *classroomUsecase) CreateClassroom(classroom *entity.Classroom, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errCreate := u.classroomRepo.Create(classroom)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(classroom)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "CREATE",
		EntityType: "CLASSROOM",
		EntityID: classroom.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *classroomUsecase) UpdateClassroom(id uint, classroom *entity.Classroom, ctx context.Context) error {

	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errUpdate := u.classroomRepo.Update(id, classroom)
	if errUpdate != nil {
		return errUpdate
	}

	classroom.ID = id

	jsonBytes, err := json.Marshal(classroom)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "CLASSROOM",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *classroomUsecase) DeleteClassroom(id uint, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "DELETE",
		EntityType: "CLASSROOM",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.classroomRepo.Delete(id)

}
