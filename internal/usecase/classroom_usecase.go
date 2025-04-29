package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type ClassroomUsecase interface {
	GetAllClassrooms() ([]entity.Classroom, error)
	CreateClassroom(classroom *entity.Classroom) error
	UpdateClassroom(id uint, classroom *entity.Classroom) error
	DeleteClassroom(id uint) error
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

func (u *classroomUsecase) CreateClassroom(classroom *entity.Classroom) error {
	// currentClassroom, err := osClassroom.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentClassroom.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.classroomRepo.Create(classroom)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(classroom)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// ClassroomID: id,
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

func (u *classroomUsecase) UpdateClassroom(id uint, classroom *entity.Classroom) error {
	
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
		// ClassroomID: id,
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

func (u *classroomUsecase) DeleteClassroom(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// ClassroomID: id,
		Action: "DELETE",
		EntityType: "CLASSROOM",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.classroomRepo.Delete(id)

}
