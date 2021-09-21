package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"platform2.0-go-challenge/models"
	"platform2.0-go-challenge/utils"
)

type ChartRepository struct{}

func (a ChartRepository) GetCharts(user_id int) ([]models.Chart, error) {
	//var chart models.Chart
	result := []models.Chart{}

	err := utils.DB.Where("user_id = ?", user_id).Where("favourite = true").Find(&result).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return []models.Chart{}, err
	}
	//fmt.Println(result)
	return result, nil
}

func (a ChartRepository) GetChartsPagination(user_id, limit, offset int) ([]models.Chart, error) {
	//var chart models.Chart
	result := []models.Chart{}
	fmt.Println(limit)
	err := utils.DB.Where("user_id = ?", user_id).Where("favourite = true").Limit(limit).Offset(offset).Find(&result).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return []models.Chart{}, err
	}
	//fmt.Println(result)
	return result, nil
}

func (a ChartRepository) EditChart(chart models.Chart, id int) (string, error) {
	result := []models.Chart{}
	//Find and update
	err := utils.DB.Where("id = ?", id).Find(&result).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return "Error", err
	}

	utils.DB.Save(&chart)

	return "Asset has been updated", err
}

func (a ChartRepository) AddChart(chart models.Chart) (int, error) {
	err := utils.DB.Save(&chart).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return 0, err
	}

	return int(chart.ID), err
}
