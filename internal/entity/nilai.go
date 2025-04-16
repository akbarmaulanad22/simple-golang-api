package entity

type Nilai struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id_nilai"`
	NIM        string  `gorm:"not null" json:"nim"`
	KodeMatkul string  `gorm:"not null" json:"kode_matkul"`
	NilaiUTS   float64 `json:"nilai_uts"`
	NilaiUAS   float64 `json:"nilai_uas"`
	NilaiAkhir float64 `json:"nilai_akhir"`
	Grade      string  `json:"grade"`
}
