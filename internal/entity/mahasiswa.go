package entity

import "time"

type Mahasiswa struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id_mahasiswa"`
	NIM          string    `gorm:"unique;not null" json:"nim"`
	Nama         string    `gorm:"not null" json:"nama"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
	Email        string    `gorm:"unique;not null" json:"email"`
	NoTelepon    string    `json:"no_telepon"`
}
