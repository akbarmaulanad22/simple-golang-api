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

type AttendanceController struct {
	attendanceUsecase usecase.AttendanceUsecase
}

func NewAttendanceController(db *gorm.DB) *AttendanceController {
	return &AttendanceController{
		attendanceUsecase: usecase.NewAttendanceUsecase(db),
	}
}

func (c *AttendanceController) GetAttendanceById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.attendanceUsecase.FindByIdAttendance(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *AttendanceController) GetAttendances(w http.ResponseWriter, r *http.Request) {
	attendances, err := c.attendanceUsecase.GetAllAttendances()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendances)
}

func (c *AttendanceController) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	var attendance entity.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.attendanceUsecase.CreateAttendance(&attendance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	attendanceResponse, err := c.attendanceUsecase.FindByIdAttendance(attendance.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attendanceResponse)
}

func (c *AttendanceController) UpdateAttendance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var attendance entity.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.attendanceUsecase.FindByIdAttendance(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.attendanceUsecase.UpdateAttendance(uint(id), &attendance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	attendanceResponse, err := c.attendanceUsecase.FindByIdAttendance(attendance.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(attendanceResponse)
}

func (c *AttendanceController) DeleteAttendance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.attendanceUsecase.FindByIdAttendance(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.attendanceUsecase.DeleteAttendance(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
