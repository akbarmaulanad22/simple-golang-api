package entity

import "time"

type Log struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     *uint     `json:"user_id" gorm:"default:null"`           // Null jika tidak ada user yang terkait
	Action     string    `json:"action" gorm:"not null"`                // Contoh: "create", "update", "delete"
	EntityType string    `json:"entity_type" gorm:"not null"`           // Contoh: "user", "student", "course"
	EntityID   uint      `json:"entity_id" gorm:"not null"`             // ID dari entitas yang terpengaruh
	Details    string    `json:"details" gorm:"type:text;default:null"` // Informasi tambahan dalam JSON
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relasi
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
}