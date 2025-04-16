package entity

import "time"

type Beasiswa struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id_beasiswa"`
	NamaBeasiswa     string    `gorm:"not null" json:"nama_beasiswa"`
	Deskripsi        string    `json:"deskripsi"`
	Persyaratan      string    `json:"persyaratan"`
	BatasPendaftaran time.Time `json:"batas_pendaftaran"`
}
