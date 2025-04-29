package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type LogUsecase interface {
	GetAllLogs() ([]entity.Log, error)
	CreateLog(log *entity.Log) error
}

type logUsecase struct {
	logRepo repository.LogRepository
}

func NewLogUsecase(db *gorm.DB) LogUsecase {
	return &logUsecase{
		logRepo: repository.NewLogRepository(db),
	}
}


func (u *logUsecase) GetAllLogs() ([]entity.Log, error) {
	return u.logRepo.FindAll()
}

func (u *logUsecase) CreateLog(log *entity.Log) error {
	// currentLog, err := osLog.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentLog.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	
	errCreate := u.logRepo.Create(log)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// LogID: id,
		Action: "CREATE",
		EntityType: "USER",
		EntityID: log.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

