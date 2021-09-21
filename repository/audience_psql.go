package repository

import (
	"errors"

	"gorm.io/gorm"
	"platform2.0-go-challenge/models"
	"platform2.0-go-challenge/utils"
)

type AudienceRepository struct{}

func (a AudienceRepository) GetAudiences(user_id int) ([]models.Audience, error) {
	result := []models.Audience{}

	err := utils.DB.Where("user_id = ?", user_id).Where("favourite = true").Find(&result).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return []models.Audience{}, err
	}
	//fmt.Println(result)
	return result, nil
}

func (a AudienceRepository) GetAudiencesPagination(user_id, limit, offset int) ([]models.Audience, error) {
	result := []models.Audience{}

	err := utils.DB.Where("user_id = ?", user_id).Where("favourite = true").Limit(limit).Offset(offset).Find(&result).Error

	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return []models.Audience{}, err
	}
	//fmt.Println(result)
	return result, nil
}

func (a AudienceRepository) EditAudience(audience models.Audience, id int) (string, error) {
	result := []models.Audience{}
	//Find and update
	err := utils.DB.Where("id = ?", id).Find(&result).Error
	utils.DB.Save(&audience)

	return "Asset has been updated", err
}

func (a AudienceRepository) AddAudience(audience models.Audience) (int, error) {
	err := utils.DB.Save(&audience).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		return 0, err
	}

	return int(audience.ID), err
}
