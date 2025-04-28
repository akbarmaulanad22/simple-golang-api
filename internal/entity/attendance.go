package entity

type Attendance struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	StudentID   uint   `json:"student_id" gorm:"not null"`
	ScheduleID  uint   `json:"schedule_id" gorm:"not null"`
	AbsenceDate string `json:"absence_date" gorm:"not null"`
	Status      string `json:"status" gorm:"not null"`
	Description string `json:"description" gorm:"default:null"`

	// Relasi
	Student  Student  `json:"student" gorm:"foreignKey:StudentID;references:ID"`
	Schedule Schedule `json:"schedule" gorm:"foreignKey:ScheduleID;references:ID"`
}