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

type LecturerController struct {
	lectureUsecase usecase.LecturerUsecase
}

func NewLecturerController(db *gorm.DB) *LecturerController {
	return &LecturerController{
		lectureUsecase: usecase.NewLecturerUsecase(db),
	}
}

func (c *LecturerController) GetLecturerById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.lectureUsecase.FindByIdLecturer(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *LecturerController) GetLecturers(w http.ResponseWriter, r *http.Request) {
	lectures, err := c.lectureUsecase.GetAllLecturers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lectures)
}

func (c *LecturerController) CreateLecturer(w http.ResponseWriter, r *http.Request) {
	var lecture entity.Lecturer
	if err := json.NewDecoder(r.Body).Decode(&lecture); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.lectureUsecase.CreateLecturer(&lecture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lectureResponse, err := c.lectureUsecase.FindByIdLecturer(lecture.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lectureResponse)
}

func (c *LecturerController) UpdateLecturer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var lecture entity.Lecturer
	if err := json.NewDecoder(r.Body).Decode(&lecture); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.lectureUsecase.FindByIdLecturer(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.lectureUsecase.UpdateLecturer(uint(id), &lecture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lectureResponse, err := c.lectureUsecase.FindByIdLecturer(lecture.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lectureResponse)
}

func (c *LecturerController) DeleteLecturer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.lectureUsecase.FindByIdLecturer(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.lectureUsecase.DeleteLecturer(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
