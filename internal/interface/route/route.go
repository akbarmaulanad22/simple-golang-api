package route

import (
	"api/internal/interface/controller"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	userController := controller.NewUserController(db)

	router.HandleFunc("/api/v1/user", userController.GetUsers).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/api/v1/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/{id}", userController.DeleteUser).Methods("DELETE")

	announcementController := controller.NewAnnouncementController(db)

	router.HandleFunc("/api/v1/announcement", announcementController.GetAnnouncements).Methods("GET")
	router.HandleFunc("/api/v1/announcement/{id}", announcementController.GetAnnouncementById).Methods("GET")
	router.HandleFunc("/api/v1/announcement", announcementController.CreateAnnouncement).Methods("POST")
	router.HandleFunc("/api/v1/announcement/{id}", announcementController.UpdateAnnouncement).Methods("PUT")
	router.HandleFunc("/api/v1/announcement/{id}", announcementController.DeleteAnnouncement).Methods("DELETE")

}
