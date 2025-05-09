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
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Announcement{})
	db.AutoMigrate(&models.Classroom{})
	db.AutoMigrate(&models.Course{})
	db.AutoMigrate(&models.Lecturer{})
	db.AutoMigrate(&models.Schedule{})
	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.Attendance{})
	db.AutoMigrate(&models.Enrollment{})
	db.AutoMigrate(&models.Faculty{})
	db.AutoMigrate(&models.Grade{})
	db.AutoMigrate(&models.Log{})
	db.AutoMigrate(&models.StudyProgram{})


	return db
}
