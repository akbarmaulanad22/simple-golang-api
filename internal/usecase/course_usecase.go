package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type CourseUsecase interface {
	GetAllCourses() ([]entity.Course, error)
	CreateCourse(course *entity.Course, ctx context.Context) error
	UpdateCourse(id uint, course *entity.Course, ctx context.Context) error
	DeleteCourse(id uint,ctx context.Context) error
	FindByIdCourse(id uint) (entity.Course, error)
}

type courseUsecase struct {
	courseRepo repository.CourseRepository
	logRepo repository.LogRepository
}

func NewCourseUsecase(db *gorm.DB) CourseUsecase {
	return &courseUsecase{
		courseRepo: repository.NewCourseRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *courseUsecase) FindByIdCourse(id uint) (entity.Course, error) {
	
	return u.courseRepo.FindById(id)
	
}

func (u *courseUsecase) GetAllCourses() ([]entity.Course, error) {
	return u.courseRepo.FindAll()
}

func (u *courseUsecase) CreateCourse(course *entity.Course, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errCreate := u.courseRepo.Create(course)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(course)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "CREATE",
		EntityType: "COURSE",
		EntityID: course.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *courseUsecase) UpdateCourse(id uint, course *entity.Course, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errUpdate := u.courseRepo.Update(id, course)
	if errUpdate != nil {
		return errUpdate
	}

	course.ID = id

	jsonBytes, err := json.Marshal(course)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "COURSE",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *courseUsecase) DeleteCourse(id uint, ctx context.Context) error {
	
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	
	errCreateLog := u.logRepo.Create(&entity.Log{
	UserID: userClaims.ID,	
		Action: "DELETE",
		EntityType: "COURSE",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.courseRepo.Delete(id)

}
