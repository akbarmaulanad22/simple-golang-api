package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type GradeUsecase interface {
	GetAllGrades() ([]entity.Grade, error)
	CreateGrade(grade *entity.Grade) error
	UpdateGrade(id uint, grade *entity.Grade) error
	DeleteGrade(id uint) error
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

func (u *gradeUsecase) CreateGrade(grade *entity.Grade) error {
	// currentGrade, err := osGrade.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentGrade.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.gradeRepo.Create(grade)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(grade)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// GradeID: id,
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

func (u *gradeUsecase) UpdateGrade(id uint, grade *entity.Grade) error {
	
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
		// GradeID: id,
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

func (u *gradeUsecase) DeleteGrade(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// GradeID: id,
		Action: "DELETE",
		EntityType: "GRADE",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.gradeRepo.Delete(id)

}
