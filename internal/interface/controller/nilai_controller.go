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

type TugasController struct {
	tugasUsecase usecase.TugasUsecase
}

func NewTugasController(db *gorm.DB) *TugasController {
	return &TugasController{
		tugasUsecase: usecase.NewTugasUsecase(db),
	}
}

func (c *TugasController) GetTugass(w http.ResponseWriter, r *http.Request) {
	tugass, err := c.tugasUsecase.GetAllTugass()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tugass)
}

func (c *TugasController) CreateTugas(w http.ResponseWriter, r *http.Request) {
	var tugas entity.Tugas
	if err := json.NewDecoder(r.Body).Decode(&tugas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.tugasUsecase.CreateTugas(&tugas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tugas)
}

func (c *TugasController) UpdateTugas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var tugas entity.Tugas
	if err := json.NewDecoder(r.Body).Decode(&tugas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.tugasUsecase.UpdateTugas(uint(id), &tugas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tugas.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tugas)
}

func (c *TugasController) DeleteTugas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.tugasUsecase.DeleteTugas(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
