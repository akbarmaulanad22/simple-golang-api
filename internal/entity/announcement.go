package entity

import "time"

type Announcement struct {
	ID             	uint      `json:"id" gorm:"primaryKey"`
	Title          	string    `json:"title" gorm:"not null"`
	Description  	string    `json:"description" gorm:"not null"`
	CreatedAt 	 	time.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy      	uint      `json:"created_by" gorm:"not null"`

	// Relasi
	CreatedByUser User `json:"created_by_user" gorm:"foreignKey:CreatedBy;references:ID"`
}