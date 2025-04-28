package entity

type Lecturer struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	UserID           uint   `json:"user_id" gorm:"unique;not null"`
	NIDN             string `json:"nidn" gorm:"unique;not null"`
	Name             string `json:"name" gorm:"not null"`
	DateOfBirth      string `json:"date_of_birth" gorm:"not null"`
	Address          string `json:"address" gorm:"not null"`
	AcademicPosition string `json:"academic_position" gorm:"not null"`

	// Relasi
	User      User       `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Schedules []Schedule `json:"schedules,omitempty" gorm:"foreignKey:LecturerID"`
}