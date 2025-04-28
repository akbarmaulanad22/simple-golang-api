package entity

type Classroom struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	RoomCode string `json:"room_code" gorm:"unique;not null"`
	Capacity int    `json:"capacity" gorm:"not null"`

	// Relasi
	Schedules []Schedule `json:"schedules,omitempty" gorm:"foreignKey:ClassroomID"`
}