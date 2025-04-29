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

type FacultyController struct {
	facultyUsecase usecase.FacultyUsecase
}

func NewFacultyController(db *gorm.DB) *FacultyController {
	return &FacultyController{
		facultyUsecase: usecase.NewFacultyUsecase(db),
	}
}

func (c *FacultyController) GetFacultyById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.facultyUsecase.FindByIdFaculty(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *FacultyController) GetFacultys(w http.ResponseWriter, r *http.Request) {
	facultys, err := c.facultyUsecase.GetAllFacultys()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(facultys)
}

func (c *FacultyController) CreateFaculty(w http.ResponseWriter, r *http.Request) {
	var faculty entity.Faculty
	if err := json.NewDecoder(r.Body).Decode(&faculty); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.facultyUsecase.CreateFaculty(&faculty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	facultyResponse, err := c.facultyUsecase.FindByIdFaculty(faculty.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(facultyResponse)
}

func (c *FacultyController) UpdateFaculty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var faculty entity.Faculty
	if err := json.NewDecoder(r.Body).Decode(&faculty); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.facultyUsecase.FindByIdFaculty(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.facultyUsecase.UpdateFaculty(uint(id), &faculty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	facultyResponse, err := c.facultyUsecase.FindByIdFaculty(faculty.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(facultyResponse)
}

func (c *FacultyController) DeleteFaculty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.facultyUsecase.FindByIdFaculty(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.facultyUsecase.DeleteFaculty(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
