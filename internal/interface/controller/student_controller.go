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

type StudentController struct {
	studentUsecase usecase.StudentUsecase
}

func NewStudentController(db *gorm.DB) *StudentController {
	return &StudentController{
		studentUsecase: usecase.NewStudentUsecase(db),
	}
}

func (c *StudentController) GetStudentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.studentUsecase.FindByIdStudent(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *StudentController) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := c.studentUsecase.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func (c *StudentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.studentUsecase.CreateStudent(&student, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentResponse, err := c.studentUsecase.FindByIdStudent(student.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(studentResponse)
}

func (c *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var student entity.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.studentUsecase.FindByIdStudent(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.studentUsecase.UpdateStudent(uint(id), &student, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentResponse, err := c.studentUsecase.FindByIdStudent(student.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentResponse)
}

func (c *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.studentUsecase.FindByIdStudent(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.studentUsecase.DeleteStudent(uint(id), r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
