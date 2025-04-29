package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type LecturerUsecase interface {
	GetAllLecturers() ([]entity.Lecturer, error)
	CreateLecturer(lecture *entity.Lecturer) error
	UpdateLecturer(id uint, lecture *entity.Lecturer) error
	DeleteLecturer(id uint) error
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

func (u *lectureUsecase) CreateLecturer(lecture *entity.Lecturer) error {
	// currentLecturer, err := osLecturer.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentLecturer.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.lectureRepo.Create(lecture)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(lecture)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// LecturerID: id,
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

func (u *lectureUsecase) UpdateLecturer(id uint, lecture *entity.Lecturer) error {
	
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
		// LecturerID: id,
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

func (u *lectureUsecase) DeleteLecturer(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// LecturerID: id,
		Action: "DELETE",
		EntityType: "LECTURER",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.lectureRepo.Delete(id)

}
