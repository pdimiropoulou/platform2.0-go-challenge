package controllers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
	"platform2.0-go-challenge/models"
	Repository "platform2.0-go-challenge/repository"
	"platform2.0-go-challenge/utils"
)

type ChartController struct{}

func (c ChartController) UpdateChart(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chart models.Chart
		var error models.Error

		json.NewDecoder(r.Body).Decode(&chart)

		if chart.ID == 0 || chart.Title == "" || chart.UserId == 0 || chart.XAxes == "" || chart.YAxes == "" || chart.Data == "" {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		chartRepo := Repository.ChartRepository{}
		rowsUpdated, err := chartRepo.EditChart(chart, int(chart.ID))

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c ChartController) AddChart(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chart models.Chart
		var error models.Error

		json.NewDecoder(r.Body).Decode(&chart)

		if chart.Title == "" || chart.UserId == 0 || chart.XAxes == "" || chart.YAxes == "" || chart.Data == "" {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		chartRepo := Repository.ChartRepository{}
		rowsUpdated, err := chartRepo.AddChart(chart)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}
