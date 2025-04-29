package entity

type Enrollment struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	StudentID    uint   `json:"student_id" gorm:"not null"`
	CourseID     uint   `json:"course_id" gorm:"not null"`
	AcademicYear string `json:"academic_year" gorm:"not null"`
	Semester     string `json:"semester" gorm:"not null"`

	// Relasi
	Student Student `json:"student" gorm:"foreignKey:StudentID;references:ID"`
	Course  Course  `json:"course" gorm:"foreignKey:CourseID;references:ID"`
	Grades  []Grade `json:"grades,omitempty" gorm:"foreignKey:EnrollmentID"`
}