package entity

type Schedule struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	CourseID    uint   `json:"course_id" gorm:"not null"`
	LecturerID  uint   `json:"lecturer_id" gorm:"not null"`
	Day         string `json:"day" gorm:"not null"`
	StartAt     string `json:"start_at" gorm:"not null"`
	EndAt       string `json:"end_at" gorm:"not null"`
	ClassroomID uint   `json:"classroom_id" gorm:"not null"` // Foreign Key ke Classrooms

	// Relasi
	Course      Course       `json:"course" gorm:"foreignKey:CourseID;references:ID"`
	Lecturer    Lecturer     `json:"lecturer" gorm:"foreignKey:LecturerID;references:ID"`
	Attendances []Attendance `json:"attendances,omitempty" gorm:"foreignKey:ScheduleID"`
	Classroom   Classroom    `json:"classroom" gorm:"foreignKey:ClassroomID;references:ID"`
}