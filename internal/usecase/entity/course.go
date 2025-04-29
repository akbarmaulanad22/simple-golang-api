package entity

type Course struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	CourseCode     string `json:"course_code" gorm:"unique;not null"`
	CourseName     string `json:"course_name" gorm:"not null"`
	SKS            int    `json:"sks" gorm:"not null"`
	Semester       int    `json:"semester" gorm:"not null"`
	StudyProgramID uint   `json:"study_program_id" gorm:"not null"`

	// Relasi
	StudyProgram StudyProgram `json:"study_program" gorm:"foreignKey:StudyProgramID;references:ID"`
	Schedules    []Schedule   `json:"schedules,omitempty" gorm:"foreignKey:CourseID"`
	Enrollments  []Enrollment `json:"enrollments,omitempty" gorm:"foreignKey:CourseID"`
}