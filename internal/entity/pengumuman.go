package entity

import "time"

type Pengumuman struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id_pengumuman"`
	Judul            string    `gorm:"not null" json:"judul"`
	IsiPengumuman    string    `gorm:"not null" json:"isi_pengumuman"`
	TanggalPublikasi time.Time `gorm:"not null" json:"tanggal_publikasi"`
	Penulis          string    `gorm:"not null" json:"penulis"`
}
