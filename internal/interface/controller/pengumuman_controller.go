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

type PengumumanController struct {
	pengumumanUsecase usecase.PengumumanUsecase
}

func NewPengumumanController(db *gorm.DB) *PengumumanController {
	return &PengumumanController{
		pengumumanUsecase: usecase.NewPengumumanUsecase(db),
	}
}

func (c *PengumumanController) GetPengumumans(w http.ResponseWriter, r *http.Request) {
	pengumumans, err := c.pengumumanUsecase.GetAllPengumumans()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pengumumans)
}

func (c *PengumumanController) CreatePengumuman(w http.ResponseWriter, r *http.Request) {
	var pengumuman entity.Pengumuman
	if err := json.NewDecoder(r.Body).Decode(&pengumuman); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.pengumumanUsecase.CreatePengumuman(&pengumuman)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pengumuman)
}

func (c *PengumumanController) UpdatePengumuman(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var pengumuman entity.Pengumuman
	if err := json.NewDecoder(r.Body).Decode(&pengumuman); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.pengumumanUsecase.UpdatePengumuman(uint(id), &pengumuman)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pengumuman.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pengumuman)
}

func (c *PengumumanController) DeletePengumuman(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.pengumumanUsecase.DeletePengumuman(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
