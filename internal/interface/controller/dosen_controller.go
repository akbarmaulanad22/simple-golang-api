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

type DosenController struct {
	dosenUsecase usecase.DosenUsecase
}

func NewDosenController(db *gorm.DB) *DosenController {
	return &DosenController{
		dosenUsecase: usecase.NewDosenUsecase(db),
	}
}

func (c *DosenController) GetDosens(w http.ResponseWriter, r *http.Request) {
	dosens, err := c.dosenUsecase.GetAllDosens()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dosens)
}

func (c *DosenController) CreateDosen(w http.ResponseWriter, r *http.Request) {
	var dosen entity.Dosen
	if err := json.NewDecoder(r.Body).Decode(&dosen); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.dosenUsecase.CreateDosen(&dosen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dosen)
}

func (c *DosenController) UpdateDosen(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var dosen entity.Dosen
	if err := json.NewDecoder(r.Body).Decode(&dosen); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.dosenUsecase.UpdateDosen(uint(id), &dosen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dosen.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dosen)
}

func (c *DosenController) DeleteDosen(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.dosenUsecase.DeleteDosen(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
