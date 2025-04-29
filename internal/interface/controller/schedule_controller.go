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

type ScheduleController struct {
	scheduleUsecase usecase.ScheduleUsecase
}

func NewScheduleController(db *gorm.DB) *ScheduleController {
	return &ScheduleController{
		scheduleUsecase: usecase.NewScheduleUsecase(db),
	}
}

func (c *ScheduleController) GetScheduleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := c.scheduleUsecase.FindByIdSchedule(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *ScheduleController) GetSchedules(w http.ResponseWriter, r *http.Request) {
	schedules, err := c.scheduleUsecase.GetAllSchedules()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

func (c *ScheduleController) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule entity.Schedule
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.scheduleUsecase.CreateSchedule(&schedule, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scheduleResponse, err := c.scheduleUsecase.FindByIdSchedule(schedule.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(scheduleResponse)
}

func (c *ScheduleController) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var schedule entity.Schedule
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.scheduleUsecase.FindByIdSchedule(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}


	err = c.scheduleUsecase.UpdateSchedule(uint(id), &schedule, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scheduleResponse, err := c.scheduleUsecase.FindByIdSchedule(schedule.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scheduleResponse)
}

func (c *ScheduleController) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = c.scheduleUsecase.FindByIdSchedule(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.scheduleUsecase.DeleteSchedule(uint(id), r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
