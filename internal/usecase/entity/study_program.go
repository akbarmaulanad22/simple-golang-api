package entity

type StudyProgram struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	StudyProgramCode string `json:"study_program_code" gorm:"unique;not null"`
	StudyProgramName string `json:"study_program_name" gorm:"not null"`
	FacultyID        uint   `json:"faculty_id" gorm:"not null"`

	// Relasi
	Faculty  Faculty   `json:"faculties" gorm:"foreignKey:FacultyID;references:ID"`
	Courses  []Course  `json:"courses,omitempty" gorm:"foreignKey:StudyProgramID"`
	Students []Student `json:"students,omitempty" gorm:"foreignKey:StudyProgramID"`
}