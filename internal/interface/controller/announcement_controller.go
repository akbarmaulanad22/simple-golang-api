package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/internal/entity"
	"api/internal/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AnnouncementController struct {
	announcementUsecase usecase.AnnouncementUsecase
}

func NewAnnouncementController(db *gorm.DB) *AnnouncementController {
	return &AnnouncementController{
		announcementUsecase: usecase.NewAnnouncementUsecase(db),
	}
}

func (c *AnnouncementController) GetAnnouncementById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.announcementUsecase.FindByIdAnnouncement(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *AnnouncementController) GetAnnouncements(w http.ResponseWriter, r *http.Request) {
	announcements, err := c.announcementUsecase.GetAllAnnouncements()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(announcements)
}

func (c *AnnouncementController) CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	var announcement entity.Announcement
	if err := json.NewDecoder(r.Body).Decode(&announcement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.announcementUsecase.CreateAnnouncement(&announcement, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	announcementResponse, err := c.announcementUsecase.FindByIdAnnouncement(announcement.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(announcementResponse)
}

func (c *AnnouncementController) UpdateAnnouncement(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var announcement entity.Announcement
	if err := json.NewDecoder(r.Body).Decode(&announcement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.announcementUsecase.FindByIdAnnouncement(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.announcementUsecase.UpdateAnnouncement(uint(id), &announcement, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	announcementResponse, err := c.announcementUsecase.FindByIdAnnouncement(announcement.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(announcementResponse)
}

func (c *AnnouncementController) DeleteAnnouncement(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.announcementUsecase.FindByIdAnnouncement(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.announcementUsecase.DeleteAnnouncement(uint(id), r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
