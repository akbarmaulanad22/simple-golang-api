package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type StudentUsecase interface {
	GetAllStudents() ([]entity.Student, error)
	CreateStudent(student *entity.Student) error
	UpdateStudent(id uint, student *entity.Student) error
	DeleteStudent(id uint) error
	FindByIdStudent(id uint) (entity.Student, error)
}

type studentUsecase struct {
	studentRepo repository.StudentRepository
	logRepo repository.LogRepository
}

func NewStudentUsecase(db *gorm.DB) StudentUsecase {
	return &studentUsecase{
		studentRepo: repository.NewStudentRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *studentUsecase) FindByIdStudent(id uint) (entity.Student, error) {
	
	return u.studentRepo.FindById(id)
	
}

func (u *studentUsecase) GetAllStudents() ([]entity.Student, error) {
	return u.studentRepo.FindAll()
}

func (u *studentUsecase) CreateStudent(student *entity.Student) error {
	// currentStudent, err := osStudent.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentStudent.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.studentRepo.Create(student)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(student)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// StudentID: id,
		Action: "CREATE",
		EntityType: "STUDENT",
		EntityID: student.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *studentUsecase) UpdateStudent(id uint, student *entity.Student) error {
	
	errUpdate := u.studentRepo.Update(id, student)
	if errUpdate != nil {
		return errUpdate
	}

	student.ID = id

	jsonBytes, err := json.Marshal(student)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// StudentID: id,
		Action: "UPDATE",
		EntityType: "STUDENT",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *studentUsecase) DeleteStudent(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// StudentID: id,
		Action: "DELETE",
		EntityType: "STUDENT",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.studentRepo.Delete(id)

}
