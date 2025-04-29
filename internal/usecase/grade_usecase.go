package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type GradeUsecase interface {
	GetAllGrades() ([]entity.Grade, error)
	CreateGrade(grade *entity.Grade, ctx context.Context) error
	UpdateGrade(id uint, grade *entity.Grade, ctx context.Context) error
	DeleteGrade(id uint, ctx context.Context) error
	FindByIdGrade(id uint) (entity.Grade, error)
}

type gradeUsecase struct {
	gradeRepo repository.GradeRepository
	logRepo repository.LogRepository
}

func NewGradeUsecase(db *gorm.DB) GradeUsecase {
	return &gradeUsecase{
		gradeRepo: repository.NewGradeRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *gradeUsecase) FindByIdGrade(id uint) (entity.Grade, error) {
	
	return u.gradeRepo.FindById(id)
	
}

func (u *gradeUsecase) GetAllGrades() ([]entity.Grade, error) {
	return u.gradeRepo.FindAll()
}

func (u *gradeUsecase) CreateGrade(grade *entity.Grade, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errCreate := u.gradeRepo.Create(grade)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(grade)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "CREATE",
		EntityType: "GRADE",
		EntityID: grade.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *gradeUsecase) UpdateGrade(id uint, grade *entity.Grade, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errUpdate := u.gradeRepo.Update(id, grade)
	if errUpdate != nil {
		return errUpdate
	}

	grade.ID = id

	jsonBytes, err := json.Marshal(grade)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "GRADE",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *gradeUsecase) DeleteGrade(id uint, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID:userClaims.ID ,
		Action: "DELETE",
		EntityType: "GRADE",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.gradeRepo.Delete(id)

}
