package entity

type MataKuliah struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id_matkul"`
	KodeMatkul string `gorm:"unique;not null" json:"kode_matkul"`
	NamaMatkul string `gorm:"not null" json:"nama_matkul"`
	SKS        int    `gorm:"not null" json:"sks"`
	Deskripsi  string `json:"deskripsi"`
}
