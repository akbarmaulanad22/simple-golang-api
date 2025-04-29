package entity

type Student struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	UserID         uint   `json:"user_id" gorm:"unique;not null"`
	NIM            string `json:"nim" gorm:"unique;not null"`
	Name           string `json:"name" gorm:"not null"`
	DateOfBirth    string `json:"date_of_birth" gorm:"not null"`
	Address        string `json:"address" gorm:"not null"`
	StudyProgramID uint   `json:"study_program_id" gorm:"not null"`
	ClassYear      string `json:"class_year" gorm:"not null"`

	// Relasi
	User         User         `json:"user" gorm:"foreignKey:UserID;references:ID"`
	StudyProgram StudyProgram `json:"study_program" gorm:"foreignKey:StudyProgramID;references:ID"`
	Enrollments  []Enrollment `json:"enrollments,omitempty" gorm:"foreignKey:StudentID"`
	Attendances  []Attendance `json:"attendances,omitempty" gorm:"foreignKey:StudentID"`
}