package config

import (
	"fmt"
	"log"
	"os"

	models "api/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	db.AutoMigrate(&models.Absensi{})
	db.AutoMigrate(&models.Beasiswa{})
	db.AutoMigrate(&models.Dosen{})
	db.AutoMigrate(&models.JadwalKuliah{})
	db.AutoMigrate(&models.Mahasiswa{})
	db.AutoMigrate(&models.MataKuliah{})
	db.AutoMigrate(&models.Nilai{})
	db.AutoMigrate(&models.Pengumuman{})
	db.AutoMigrate(&models.RuangKelas{})
	db.AutoMigrate(&models.Tugas{})


	return db
}
