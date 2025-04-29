package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type StudyProgramUsecase interface {
	GetAllStudyPrograms() ([]entity.StudyProgram, error)
	CreateStudyProgram(studyProgram *entity.StudyProgram) error
	UpdateStudyProgram(id uint, studyProgram *entity.StudyProgram) error
	DeleteStudyProgram(id uint) error
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

func (u *studyProgramUsecase) CreateStudyProgram(studyProgram *entity.StudyProgram) error {
	// currentStudyProgram, err := osStudyProgram.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentStudyProgram.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.studyProgramRepo.Create(studyProgram)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(studyProgram)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// StudyProgramID: id,
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

func (u *studyProgramUsecase) UpdateStudyProgram(id uint, studyProgram *entity.StudyProgram) error {
	
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
		// StudyProgramID: id,
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

func (u *studyProgramUsecase) DeleteStudyProgram(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// StudyProgramID: id,
		Action: "DELETE",
		EntityType: "STUDYPROGRAM",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.studyProgramRepo.Delete(id)

}
