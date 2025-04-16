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

type JadwalKuliahController struct {
	jadwalKuliahUsecase usecase.JadwalKuliahUsecase
}

func NewJadwalKuliahController(db *gorm.DB) *JadwalKuliahController {
	return &JadwalKuliahController{
		jadwalKuliahUsecase: usecase.NewJadwalKuliahUsecase(db),
	}
}

func (c *JadwalKuliahController) GetJadwalKuliahs(w http.ResponseWriter, r *http.Request) {
	jadwalKuliahs, err := c.jadwalKuliahUsecase.GetAllJadwalKuliahs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jadwalKuliahs)
}

func (c *JadwalKuliahController) CreateJadwalKuliah(w http.ResponseWriter, r *http.Request) {
	var jadwalKuliah entity.JadwalKuliah
	if err := json.NewDecoder(r.Body).Decode(&jadwalKuliah); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.jadwalKuliahUsecase.CreateJadwalKuliah(&jadwalKuliah)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(jadwalKuliah)
}

func (c *JadwalKuliahController) UpdateJadwalKuliah(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var jadwalKuliah entity.JadwalKuliah
	if err := json.NewDecoder(r.Body).Decode(&jadwalKuliah); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.jadwalKuliahUsecase.UpdateJadwalKuliah(uint(id), &jadwalKuliah)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jadwalKuliah.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jadwalKuliah)
}

func (c *JadwalKuliahController) DeleteJadwalKuliah(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.jadwalKuliahUsecase.DeleteJadwalKuliah(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
