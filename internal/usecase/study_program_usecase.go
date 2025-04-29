package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type StudyProgramUsecase interface {
	GetAllStudyPrograms() ([]entity.StudyProgram, error)
	CreateStudyProgram(studyProgram *entity.StudyProgram, ctx context.Context) error
	UpdateStudyProgram(id uint, studyProgram *entity.StudyProgram, ctx context.Context) error
	DeleteStudyProgram(id uint, ctx context.Context) error
	FindByIdStudyProgram(id uint) (entity.StudyProgram, error)
}

type studyProgramUsecase struct {
	studyProgramRepo repository.StudyProgramRepository
	logRepo repository.LogRepository
}

func NewStudyProgramUsecase(db *gorm.DB) StudyProgramUsecase {
	return &studyProgramUsecase{
		studyProgramRepo: repository.NewStudyProgramRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *studyProgramUsecase) FindByIdStudyProgram(id uint) (entity.StudyProgram, error) {
	
	return u.studyProgramRepo.FindById(id)
	
}

func (u *studyProgramUsecase) GetAllStudyPrograms() ([]entity.StudyProgram, error) {
	return u.studyProgramRepo.FindAll()
}

func (u *studyProgramUsecase) CreateStudyProgram(studyProgram *entity.StudyProgram, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errCreate := u.studyProgramRepo.Create(studyProgram)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(studyProgram)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "CREATE",
		EntityType: "STUDYPROGRAM",
		EntityID: studyProgram.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *studyProgramUsecase) UpdateStudyProgram(id uint, studyProgram *entity.StudyProgram, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errUpdate := u.studyProgramRepo.Update(id, studyProgram)
	if errUpdate != nil {
		return errUpdate
	}

	studyProgram.ID = id

	jsonBytes, err := json.Marshal(studyProgram)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "STUDYPROGRAM",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *studyProgramUsecase) DeleteStudyProgram(id uint, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "DELETE",
		EntityType: "STUDYPROGRAM",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.studyProgramRepo.Delete(id)

}
