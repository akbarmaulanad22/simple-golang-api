package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type FacultyUsecase interface {
	GetAllFacultys() ([]entity.Faculty, error)
	CreateFaculty(faculty *entity.Faculty, ctx context.Context) error
	UpdateFaculty(id uint, faculty *entity.Faculty, ctx context.Context) error	
	DeleteFaculty(id uint, ctx context.Context) error
	FindByIdFaculty(id uint) (entity.Faculty, error)
}

type facultyUsecase struct {
	facultyRepo repository.FacultyRepository
	logRepo repository.LogRepository
}

func NewFacultyUsecase(db *gorm.DB) FacultyUsecase {
	return &facultyUsecase{
		facultyRepo: repository.NewFacultyRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *facultyUsecase) FindByIdFaculty(id uint) (entity.Faculty, error) {
	
	return u.facultyRepo.FindById(id)
	
}

func (u *facultyUsecase) GetAllFacultys() ([]entity.Faculty, error) {
	return u.facultyRepo.FindAll()
}

func (u *facultyUsecase) CreateFaculty(faculty *entity.Faculty, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreate := u.facultyRepo.Create(faculty)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(faculty)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "CREATE",
		EntityType: "FACULTY",
		EntityID: faculty.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *facultyUsecase) UpdateFaculty(id uint, faculty *entity.Faculty, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errUpdate := u.facultyRepo.Update(id, faculty)
	if errUpdate != nil {
		return errUpdate
	}

	faculty.ID = id

	jsonBytes, err := json.Marshal(faculty)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "FACULTY",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *facultyUsecase) DeleteFaculty(id uint, ctx context.Context) error {

	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "DELETE",
		EntityType: "FACULTY",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.facultyRepo.Delete(id)

}
