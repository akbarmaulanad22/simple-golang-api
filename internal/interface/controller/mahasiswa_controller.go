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

type MahasiswaController struct {
	mahasiswaUsecase usecase.MahasiswaUsecase
}

func NewMahasiswaController(db *gorm.DB) *MahasiswaController {
	return &MahasiswaController{
		mahasiswaUsecase: usecase.NewMahasiswaUsecase(db),
	}
}

func (c *MahasiswaController) GetMahasiswas(w http.ResponseWriter, r *http.Request) {
	mahasiswas, err := c.mahasiswaUsecase.GetAllMahasiswas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mahasiswas)
}

func (c *MahasiswaController) CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var mahasiswa entity.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mahasiswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.mahasiswaUsecase.CreateMahasiswa(&mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mahasiswa)
}

func (c *MahasiswaController) UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var mahasiswa entity.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mahasiswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.mahasiswaUsecase.UpdateMahasiswa(uint(id), &mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mahasiswa.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mahasiswa)
}

func (c *MahasiswaController) DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.mahasiswaUsecase.DeleteMahasiswa(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
