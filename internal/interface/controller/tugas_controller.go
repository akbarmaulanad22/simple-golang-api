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

type NilaiController struct {
	nilaiUsecase usecase.NilaiUsecase
}

func NewNilaiController(db *gorm.DB) *NilaiController {
	return &NilaiController{
		nilaiUsecase: usecase.NewNilaiUsecase(db),
	}
}

func (c *NilaiController) GetNilais(w http.ResponseWriter, r *http.Request) {
	nilais, err := c.nilaiUsecase.GetAllNilais()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilais)
}

func (c *NilaiController) CreateNilai(w http.ResponseWriter, r *http.Request) {
	var nilai entity.Nilai
	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.nilaiUsecase.CreateNilai(&nilai)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nilai)
}

func (c *NilaiController) UpdateNilai(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var nilai entity.Nilai
	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.nilaiUsecase.UpdateNilai(uint(id), &nilai)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nilai.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilai)
}

func (c *NilaiController) DeleteNilai(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.nilaiUsecase.DeleteNilai(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
