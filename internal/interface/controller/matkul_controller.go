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

type MataKuliahController struct {
	mataKuliahUsecase usecase.MataKuliahUsecase
}

func NewMataKuliahController(db *gorm.DB) *MataKuliahController {
	return &MataKuliahController{
		mataKuliahUsecase: usecase.NewMataKuliahUsecase(db),
	}
}

func (c *MataKuliahController) GetMataKuliahs(w http.ResponseWriter, r *http.Request) {
	mataKuliahs, err := c.mataKuliahUsecase.GetAllMataKuliahs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mataKuliahs)
}

func (c *MataKuliahController) CreateMataKuliah(w http.ResponseWriter, r *http.Request) {
	var mataKuliah entity.MataKuliah
	if err := json.NewDecoder(r.Body).Decode(&mataKuliah); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.mataKuliahUsecase.CreateMataKuliah(&mataKuliah)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mataKuliah)
}

func (c *MataKuliahController) UpdateMataKuliah(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var mataKuliah entity.MataKuliah
	if err := json.NewDecoder(r.Body).Decode(&mataKuliah); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.mataKuliahUsecase.UpdateMataKuliah(uint(id), &mataKuliah)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mataKuliah.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mataKuliah)
}

func (c *MataKuliahController) DeleteMataKuliah(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.mataKuliahUsecase.DeleteMataKuliah(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
