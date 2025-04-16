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

type RuangKelasController struct {
	ruangKelasUsecase usecase.RuangKelasUsecase
}

func NewRuangKelasController(db *gorm.DB) *RuangKelasController {
	return &RuangKelasController{
		ruangKelasUsecase: usecase.NewRuangKelasUsecase(db),
	}
}

func (c *RuangKelasController) GetRuangKelass(w http.ResponseWriter, r *http.Request) {
	ruangKelass, err := c.ruangKelasUsecase.GetAllRuangKelass()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ruangKelass)
}

func (c *RuangKelasController) CreateRuangKelas(w http.ResponseWriter, r *http.Request) {
	var ruangKelas entity.RuangKelas
	if err := json.NewDecoder(r.Body).Decode(&ruangKelas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.ruangKelasUsecase.CreateRuangKelas(&ruangKelas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ruangKelas)
}

func (c *RuangKelasController) UpdateRuangKelas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var ruangKelas entity.RuangKelas
	if err := json.NewDecoder(r.Body).Decode(&ruangKelas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.ruangKelasUsecase.UpdateRuangKelas(uint(id), &ruangKelas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ruangKelas.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ruangKelas)
}

func (c *RuangKelasController) DeleteRuangKelas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.ruangKelasUsecase.DeleteRuangKelas(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
