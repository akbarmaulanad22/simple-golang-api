package controller

import (
	"encoding/json"
	"net/http"

	"api/internal/entity"
	"api/internal/usecase"

	"gorm.io/gorm"
)

type LogController struct {
	logUsecase usecase.LogUsecase
}

func NewLogController(db *gorm.DB) *LogController {
	return &LogController{
		logUsecase: usecase.NewLogUsecase(db),
	}
}



func (c *LogController) GetLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := c.logUsecase.GetAllLogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func (c *LogController) CreateLog(w http.ResponseWriter, r *http.Request) {
	var log entity.Log
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.logUsecase.CreateLog(&log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(log)
}