package entity

type RuangKelas struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id_ruangan"`
	KodeRuangan string `gorm:"unique;not null" json:"kode_ruangan"`
	Kapasitas   int    `gorm:"not null" json:"kapasitas"`
	Fasilitas   string `json:"fasilitas"`
	Lokasi      string `json:"lokasi"`
}
