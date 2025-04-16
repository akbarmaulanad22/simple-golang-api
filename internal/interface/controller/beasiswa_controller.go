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

type BeasiswaController struct {
	beasiswaUsecase usecase.BeasiswaUsecase
}

func NewBeasiswaController(db *gorm.DB) *BeasiswaController {
	return &BeasiswaController{
		beasiswaUsecase: usecase.NewBeasiswaUsecase(db),
	}
}

func (c *BeasiswaController) GetBeasiswas(w http.ResponseWriter, r *http.Request) {
	beasiswas, err := c.beasiswaUsecase.GetAllBeasiswas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beasiswas)
}

func (c *BeasiswaController) CreateBeasiswa(w http.ResponseWriter, r *http.Request) {
	var beasiswa entity.Beasiswa
	if err := json.NewDecoder(r.Body).Decode(&beasiswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.beasiswaUsecase.CreateBeasiswa(&beasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(beasiswa)
}

func (c *BeasiswaController) UpdateBeasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var beasiswa entity.Beasiswa
	if err := json.NewDecoder(r.Body).Decode(&beasiswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.beasiswaUsecase.UpdateBeasiswa(uint(id), &beasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	beasiswa.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beasiswa)
}

func (c *BeasiswaController) DeleteBeasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.beasiswaUsecase.DeleteBeasiswa(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
