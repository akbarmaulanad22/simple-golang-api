package entity

type Faculty struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FacultyCode string `json:"faculty_code" gorm:"unique;not null"`
	FacultyName string `json:"faculty_name" gorm:"not null"`

	// Relasi
	Programs []StudyProgram `json:"study_programs,omitempty" gorm:"foreignKey:FacultyID"`
}