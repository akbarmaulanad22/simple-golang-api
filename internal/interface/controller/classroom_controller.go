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

type ClassroomController struct {
	classroomUsecase usecase.ClassroomUsecase
}

func NewClassroomController(db *gorm.DB) *ClassroomController {
	return &ClassroomController{
		classroomUsecase: usecase.NewClassroomUsecase(db),
	}
}

func (c *ClassroomController) GetClassroomById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.classroomUsecase.FindByIdClassroom(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *ClassroomController) GetClassrooms(w http.ResponseWriter, r *http.Request) {
	classrooms, err := c.classroomUsecase.GetAllClassrooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classrooms)
}

func (c *ClassroomController) CreateClassroom(w http.ResponseWriter, r *http.Request) {
	var classroom entity.Classroom
	if err := json.NewDecoder(r.Body).Decode(&classroom); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.classroomUsecase.CreateClassroom(&classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	classroomResponse, err := c.classroomUsecase.FindByIdClassroom(classroom.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(classroomResponse)
}

func (c *ClassroomController) UpdateClassroom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var classroom entity.Classroom
	if err := json.NewDecoder(r.Body).Decode(&classroom); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.classroomUsecase.FindByIdClassroom(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.classroomUsecase.UpdateClassroom(uint(id), &classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	classroomResponse, err := c.classroomUsecase.FindByIdClassroom(classroom.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classroomResponse)
}

func (c *ClassroomController) DeleteClassroom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.classroomUsecase.FindByIdClassroom(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.classroomUsecase.DeleteClassroom(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
