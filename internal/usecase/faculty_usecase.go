package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type FacultyUsecase interface {
	GetAllFacultys() ([]entity.Faculty, error)
	CreateFaculty(faculty *entity.Faculty) error
	UpdateFaculty(id uint, faculty *entity.Faculty) error
	DeleteFaculty(id uint) error
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

func (u *facultyUsecase) CreateFaculty(faculty *entity.Faculty) error {
	// currentFaculty, err := osFaculty.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentFaculty.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.facultyRepo.Create(faculty)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(faculty)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// FacultyID: id,
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

func (u *facultyUsecase) UpdateFaculty(id uint, faculty *entity.Faculty) error {
	
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
		// FacultyID: id,
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

func (u *facultyUsecase) DeleteFaculty(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// FacultyID: id,
		Action: "DELETE",
		EntityType: "FACULTY",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.facultyRepo.Delete(id)

}
