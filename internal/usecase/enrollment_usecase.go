package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type EnrollmentUsecase interface {
	GetAllEnrollments() ([]entity.Enrollment, error)
	CreateEnrollment(enrollment *entity.Enrollment, ctx context.Context) error
	UpdateEnrollment(id uint, enrollment *entity.Enrollment, ctx context.Context) error
	DeleteEnrollment(id uint, ctx context.Context) error
	FindByIdEnrollment(id uint) (entity.Enrollment, error)
}

type enrollmentUsecase struct {
	enrollmentRepo repository.EnrollmentRepository
	logRepo repository.LogRepository
}

func NewEnrollmentUsecase(db *gorm.DB) EnrollmentUsecase {
	return &enrollmentUsecase{
		enrollmentRepo: repository.NewEnrollmentRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *enrollmentUsecase) FindByIdEnrollment(id uint) (entity.Enrollment, error) {
	
	return u.enrollmentRepo.FindById(id)
	
}

func (u *enrollmentUsecase) GetAllEnrollments() ([]entity.Enrollment, error) {
	return u.enrollmentRepo.FindAll()
}

func (u *enrollmentUsecase) CreateEnrollment(enrollment *entity.Enrollment, ctx context.Context) error {

	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }	
	errCreate := u.enrollmentRepo.Create(enrollment)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(enrollment)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "CREATE",
		EntityType: "ENROLLMENT",
		EntityID: enrollment.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *enrollmentUsecase) UpdateEnrollment(id uint, enrollment *entity.Enrollment,ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errUpdate := u.enrollmentRepo.Update(id, enrollment)
	if errUpdate != nil {
		return errUpdate
	}

	enrollment.ID = id

	jsonBytes, err := json.Marshal(enrollment)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "ENROLLMENT",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *enrollmentUsecase) DeleteEnrollment(id uint, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "DELETE",
		EntityType: "ENROLLMENT",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.enrollmentRepo.Delete(id)

}
