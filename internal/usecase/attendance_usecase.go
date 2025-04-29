package usecase

import (
	"api/internal/entity"
	"api/internal/repository"
	"encoding/json"

	"gorm.io/gorm"
)

type AttendanceUsecase interface {
	GetAllAttendances() ([]entity.Attendance, error)
	CreateAttendance(attendance *entity.Attendance) error
	UpdateAttendance(id uint, attendance *entity.Attendance) error
	DeleteAttendance(id uint) error
	FindByIdAttendance(id uint) (entity.Attendance, error)
}

type attendanceUsecase struct {
	attendanceRepo repository.AttendanceRepository
	logRepo repository.LogRepository
}

func NewAttendanceUsecase(db *gorm.DB) AttendanceUsecase {
	return &attendanceUsecase{
		attendanceRepo: repository.NewAttendanceRepository(db),
		logRepo: repository.NewLogRepository(db),
	}
}

func (u *attendanceUsecase) FindByIdAttendance(id uint) (entity.Attendance, error) {
	
	return u.attendanceRepo.FindById(id)
	
}

func (u *attendanceUsecase) GetAllAttendances() ([]entity.Attendance, error) {
	return u.attendanceRepo.FindAll()
}

func (u *attendanceUsecase) CreateAttendance(attendance *entity.Attendance) error {
	// currentAttendance, err := osAttendance.Current()
	// if err != nil {
	// 	return err
	// }
	
	// idInt, err := strconv.Atoi(currentAttendance.Uid)
	// if err != nil {
	// 	return err
	// }

	// id := new(uint)
    // *id = uint(idInt)
	
	errCreate := u.attendanceRepo.Create(attendance)
	if errCreate != nil {
		return errCreate
	}

	jsonBytes, err := json.Marshal(attendance)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// AttendanceID: id,
		Action: "CREATE",
		EntityType: "ATTENDANCE",
		EntityID: attendance.ID,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return nil
}

func (u *attendanceUsecase) UpdateAttendance(id uint, attendance *entity.Attendance) error {
	
	errUpdate := u.attendanceRepo.Update(id, attendance)
	if errUpdate != nil {
		return errUpdate
	}

	attendance.ID = id

	jsonBytes, err := json.Marshal(attendance)
	if err != nil {
		return err
	}
	
	errCreateLog := u.logRepo.Create(&entity.Log{
		// AttendanceID: id,
		Action: "UPDATE",
		EntityType: "ATTENDANCE",
		EntityID: id,
		Details: string(jsonBytes),
	})

	if errCreateLog != nil{
		return errCreateLog
	}
	
	return nil
}

func (u *attendanceUsecase) DeleteAttendance(id uint) error {
	errCreateLog := u.logRepo.Create(&entity.Log{
		// AttendanceID: id,
		Action: "DELETE",
		EntityType: "ATTENDANCE",
		EntityID: id,
	})

	if errCreateLog != nil{
		return errCreateLog
	}

	return u.attendanceRepo.Delete(id)

}
