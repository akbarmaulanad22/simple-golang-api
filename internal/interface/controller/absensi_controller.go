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

type AbsensiController struct {
	absensiUsecase usecase.AbsensiUsecase
}

func NewAbsensiController(db *gorm.DB) *AbsensiController {
	return &AbsensiController{
		absensiUsecase: usecase.NewAbsensiUsecase(db),
	}
}

func (c *AbsensiController) GetAbsensis(w http.ResponseWriter, r *http.Request) {
	absensis, err := c.absensiUsecase.GetAllAbsensis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(absensis)
}

func (c *AbsensiController) CreateAbsensi(w http.ResponseWriter, r *http.Request) {
	var absensi entity.Absensi
	if err := json.NewDecoder(r.Body).Decode(&absensi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.absensiUsecase.CreateAbsensi(&absensi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(absensi)
}

func (c *AbsensiController) UpdateAbsensi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var absensi entity.Absensi
	if err := json.NewDecoder(r.Body).Decode(&absensi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.absensiUsecase.UpdateAbsensi(uint(id), &absensi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	absensi.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(absensi)
}

func (c *AbsensiController) DeleteAbsensi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.absensiUsecase.DeleteAbsensi(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
