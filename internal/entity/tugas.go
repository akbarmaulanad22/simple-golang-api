package entity

import "time"

type Tugas struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id_tugas"`
	KodeMatkul string    `gorm:"not null" json:"kode_matkul"`
	JudulTugas string    `gorm:"not null" json:"judul_tugas"`
	Deskripsi  string    `json:"deskripsi"`
	Deadline   time.Time `json:"deadline"`
}
