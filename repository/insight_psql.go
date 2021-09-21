package repository

import (
	"errors"

	"gorm.io/gorm"
	"platform2.0-go-challenge/models"
	"platform2.0-go-challenge/utils"
)

type Insightrepository struct{}

func (a Insightrepository) GetInsights(user_id int) ([]models.Insight, error) {
	result := []models.Insight{}

	err := utils.DB.Where("user_id = ?", user_id).Where("favourite = true").Find(&result).Find(&result).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return []models.Insight{}, err
	}
	//fmt.Println(result)
	return result, nil
}

func (a Insightrepository) GetInsightsPagination(user_id, limit, offset int) ([]models.Insight, error) {
	result := []models.Insight{}

	err := utils.DB.Where("user_id = ?", user_id).Where("favourite = true").Find(&result).Limit(limit).Offset(offset).Find(&result).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return []models.Insight{}, err
	}
	//fmt.Println(result)
	return result, nil
}

func (a Insightrepository) EditInsight(insight models.Insight, id int) (string, error) {
	result := []models.Insight{}
	//Find and update
	err := utils.DB.Where("id = ?", id).Find(&result).Error
	utils.DB.Save(&insight)

	return "Asset has been updated", err
}

func (a Insightrepository) AddInsight(insight models.Insight) (int, error) {
	err := utils.DB.Save(&insight).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return 0, err
	}

	return int(insight.ID), err
}
