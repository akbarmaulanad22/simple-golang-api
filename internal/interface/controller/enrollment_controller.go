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

type EnrollmentController struct {
	enrollmentUsecase usecase.EnrollmentUsecase
}

func NewEnrollmentController(db *gorm.DB) *EnrollmentController {
	return &EnrollmentController{
		enrollmentUsecase: usecase.NewEnrollmentUsecase(db),
	}
}

func (c *EnrollmentController) GetEnrollmentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.enrollmentUsecase.FindByIdEnrollment(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *EnrollmentController) GetEnrollments(w http.ResponseWriter, r *http.Request) {
	enrollments, err := c.enrollmentUsecase.GetAllEnrollments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrollments)
}

func (c *EnrollmentController) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	var enrollment entity.Enrollment
	if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.enrollmentUsecase.CreateEnrollment(&enrollment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enrollmentResponse, err := c.enrollmentUsecase.FindByIdEnrollment(enrollment.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(enrollmentResponse)
}

func (c *EnrollmentController) UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var enrollment entity.Enrollment
	if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.enrollmentUsecase.FindByIdEnrollment(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.enrollmentUsecase.UpdateEnrollment(uint(id), &enrollment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enrollmentResponse, err := c.enrollmentUsecase.FindByIdEnrollment(enrollment.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrollmentResponse)
}

func (c *EnrollmentController) DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.enrollmentUsecase.FindByIdEnrollment(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.enrollmentUsecase.DeleteEnrollment(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
