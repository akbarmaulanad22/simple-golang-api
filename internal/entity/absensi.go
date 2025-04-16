package entity

import "time"

type Absensi struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id_absensi"`
	NIM             string    `gorm:"not null" json:"nim"`
	Tanggal         time.Time `gorm:"not null" json:"tanggal"`
	StatusKehadiran string    `gorm:"not null" json:"status_kehadiran"`
}
