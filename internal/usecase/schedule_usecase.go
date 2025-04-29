package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type ScheduleUsecase interface {
	GetAllSchedules() ([]entity.Schedule, error)
	CreateSchedule(schedule *entity.Schedule) error
	UpdateSchedule(id uint, schedule *entity.Schedule) error
	DeleteSchedule(id uint) error
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

func (u *scheduleUsecase) CreateSchedule(schedule *entity.Schedule) error {
	// currentSchedule, err := osSchedule.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentSchedule.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.scheduleRepo.Create(schedule)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// ScheduleID: id,
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

func (u *scheduleUsecase) UpdateSchedule(id uint, schedule *entity.Schedule) error {
	
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
		// ScheduleID: id,
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

func (u *scheduleUsecase) DeleteSchedule(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// ScheduleID: id,
		Action: "DELETE",
		EntityType: "SCHEDULE",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.scheduleRepo.Delete(id)

}
