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

type CourseController struct {
	courseUsecase usecase.CourseUsecase
}

func NewCourseController(db *gorm.DB) *CourseController {
	return &CourseController{
		courseUsecase: usecase.NewCourseUsecase(db),
	}
}

func (c *CourseController) GetCourseById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.courseUsecase.FindByIdCourse(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *CourseController) GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := c.courseUsecase.GetAllCourses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func (c *CourseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course entity.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.courseUsecase.CreateCourse(&course, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	courseResponse, err := c.courseUsecase.FindByIdCourse(course.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(courseResponse)
}

func (c *CourseController) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var course entity.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.courseUsecase.FindByIdCourse(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.courseUsecase.UpdateCourse(uint(id), &course, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	courseResponse, err := c.courseUsecase.FindByIdCourse(course.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courseResponse)
}

func (c *CourseController) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.courseUsecase.FindByIdCourse(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.courseUsecase.DeleteCourse(uint(id),r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
