package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type EnrollmentUsecase interface {
	GetAllEnrollments() ([]entity.Enrollment, error)
	CreateEnrollment(enrollment *entity.Enrollment) error
	UpdateEnrollment(id uint, enrollment *entity.Enrollment) error
	DeleteEnrollment(id uint) error
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

func (u *enrollmentUsecase) CreateEnrollment(enrollment *entity.Enrollment) error {
	// currentEnrollment, err := osEnrollment.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentEnrollment.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.enrollmentRepo.Create(enrollment)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(enrollment)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// EnrollmentID: id,
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

func (u *enrollmentUsecase) UpdateEnrollment(id uint, enrollment *entity.Enrollment) error {
	
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
		// EnrollmentID: id,
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

func (u *enrollmentUsecase) DeleteEnrollment(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// EnrollmentID: id,
		Action: "DELETE",
		EntityType: "ENROLLMENT",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.enrollmentRepo.Delete(id)

}
