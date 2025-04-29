package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type ScheduleUsecase interface {
	GetAllSchedules() ([]entity.Schedule, error)
	CreateSchedule(schedule *entity.Schedule, ctx context.Context) error
	UpdateSchedule(id uint, schedule *entity.Schedule, ctx context.Context) error
	DeleteSchedule(id uint,ctx context.Context) error
	FindByIdSchedule(id uint) (entity.Schedule, error)
}

type scheduleUsecase struct {
	scheduleRepo repository.ScheduleRepository
	logRepo repository.LogRepository
}

func NewScheduleUsecase(db *gorm.DB) ScheduleUsecase {
	return &scheduleUsecase{
		scheduleRepo: repository.NewScheduleRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *scheduleUsecase) FindByIdSchedule(id uint) (entity.Schedule, error) {
	
	return u.scheduleRepo.FindById(id)
	
}

func (u *scheduleUsecase) GetAllSchedules() ([]entity.Schedule, error) {
	return u.scheduleRepo.FindAll()
}

func (u *scheduleUsecase) CreateSchedule(schedule *entity.Schedule, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreate := u.scheduleRepo.Create(schedule)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID:userClaims.ID ,
		Action: "CREATE",
		EntityType: "SCHEDULE",
		EntityID: schedule.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *scheduleUsecase) UpdateSchedule(id uint, schedule *entity.Schedule, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errUpdate := u.scheduleRepo.Update(id, schedule)
	if errUpdate != nil {
		return errUpdate
	}

	schedule.ID = id

	jsonBytes, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "UPDATE",
		EntityType: "SCHEDULE",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *scheduleUsecase) DeleteSchedule(id uint, ctx context.Context) error {
	userClaims, ok := ctx.Value("user").(*entity.CustomClaims)
    if !ok {
        return errors.New("user not authenticated")
    }
	errCreateLog := u.logRepo.Create(&entity.Log{
		UserID: userClaims.ID,
		Action: "DELETE",
		EntityType: "SCHEDULE",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.scheduleRepo.Delete(id)

}
