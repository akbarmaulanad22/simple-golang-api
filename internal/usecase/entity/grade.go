package entity

type Grade struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	EnrollmentID uint    `json:"enrollment_id" gorm:"not null"`
	NilaiAkhir   float64 `json:"nilai_akhir" gorm:"not null"`
	GradeLetter  string  `json:"grade_letter" gorm:"not null"`

	// Relasi
	Enrollment Enrollment `json:"enrollment" gorm:"foreignKey:EnrollmentID;references:ID"`
}