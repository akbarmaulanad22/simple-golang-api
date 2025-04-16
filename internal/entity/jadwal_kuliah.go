package entity

import "time"

type JadwalKuliah struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id_jadwal"`
	Hari       string    `gorm:"not null" json:"hari"`
	JamMulai   time.Time `json:"jam_mulai"`
	JamSelesai time.Time `json:"jam_selesai"`
	Ruangan    string    `json:"ruangan"`
	Semester   string    `json:"semester"`
}
