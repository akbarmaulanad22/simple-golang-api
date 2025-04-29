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

type GradeController struct {
	gradeUsecase usecase.GradeUsecase
}

func NewGradeController(db *gorm.DB) *GradeController {
	return &GradeController{
		gradeUsecase: usecase.NewGradeUsecase(db),
	}
}

func (c *GradeController) GetGradeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.gradeUsecase.FindByIdGrade(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *GradeController) GetGrades(w http.ResponseWriter, r *http.Request) {
	grades, err := c.gradeUsecase.GetAllGrades()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grades)
}

func (c *GradeController) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade entity.Grade
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.gradeUsecase.CreateGrade(&grade, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gradeResponse, err := c.gradeUsecase.FindByIdGrade(grade.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(gradeResponse)
}

func (c *GradeController) UpdateGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var grade entity.Grade
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.gradeUsecase.FindByIdGrade(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.gradeUsecase.UpdateGrade(uint(id), &grade, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gradeResponse, err := c.gradeUsecase.FindByIdGrade(grade.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gradeResponse)
}

func (c *GradeController) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.gradeUsecase.FindByIdGrade(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.gradeUsecase.DeleteGrade(uint(id), r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
