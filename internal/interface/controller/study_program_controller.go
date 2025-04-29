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

type StudyProgramController struct {
	studyProgramUsecase usecase.StudyProgramUsecase
}

func NewStudyProgramController(db *gorm.DB) *StudyProgramController {
	return &StudyProgramController{
		studyProgramUsecase: usecase.NewStudyProgramUsecase(db),
	}
}

func (c *StudyProgramController) GetStudyProgramById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.studyProgramUsecase.FindByIdStudyProgram(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *StudyProgramController) GetStudyPrograms(w http.ResponseWriter, r *http.Request) {
	studyPrograms, err := c.studyProgramUsecase.GetAllStudyPrograms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studyPrograms)
}

func (c *StudyProgramController) CreateStudyProgram(w http.ResponseWriter, r *http.Request) {
	var studyProgram entity.StudyProgram
	if err := json.NewDecoder(r.Body).Decode(&studyProgram); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.studyProgramUsecase.CreateStudyProgram(&studyProgram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studyProgramResponse, err := c.studyProgramUsecase.FindByIdStudyProgram(studyProgram.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(studyProgramResponse)
}

func (c *StudyProgramController) UpdateStudyProgram(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var studyProgram entity.StudyProgram
	if err := json.NewDecoder(r.Body).Decode(&studyProgram); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.studyProgramUsecase.FindByIdStudyProgram(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.studyProgramUsecase.UpdateStudyProgram(uint(id), &studyProgram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studyProgramResponse, err := c.studyProgramUsecase.FindByIdStudyProgram(studyProgram.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studyProgramResponse)
}

func (c *StudyProgramController) DeleteStudyProgram(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.studyProgramUsecase.FindByIdStudyProgram(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.studyProgramUsecase.DeleteStudyProgram(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
