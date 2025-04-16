package entity

type Dosen struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id_dosen"`
	NIDN            string `gorm:"unique;not null" json:"nidn"`
	Nama            string `gorm:"not null" json:"nama"`
	JabatanAkademik string `json:"jabatan_akademik"`
	Email           string `gorm:"unique;not null" json:"email"`
	NoTelepon       string `json:"no_telepon"`
}
