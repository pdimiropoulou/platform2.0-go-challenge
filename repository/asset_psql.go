package repository

import (
	"platform2.0-go-challenge/models"
)

type AssetRepository struct{}

func (a AssetRepository) GetUserAssetsPagination(user_id, limit, offset int) (*models.AssetReponse, error) {
	var response models.AssetReponse
	var err error
	errs := make(chan error)

	go getAllPagination(user_id, limit, offset, &response, errs)

	temp := <-errs
	if temp != nil {
		err = temp
	}
	response.UserId = user_id

	return &response, err
}

func (a AssetRepository) GetUserAssets(user_id int) (*models.AssetReponse, error) {
	var response models.AssetReponse
	var err error
	errs := make(chan error)

	go getAll(user_id, &response, errs)

	temp := <-errs
	if temp != nil {
		err = temp
	}

	response.UserId = user_id

	return &response, err
}

func getAll(user_id int, response *models.AssetReponse, errs chan error) {
	chartrepo := ChartRepository{}
	charts, err := chartrepo.GetCharts(user_id)
	response.Charts = charts

	insightsrepo := Insightrepository{}
	insights, err := insightsrepo.GetInsights(user_id)
	response.Insights = insights

	audiencesrepo := AudienceRepository{}
	audiences, err := audiencesrepo.GetAudiences(user_id)
	response.Audiences = audiences

	errs <- err
}

func getAllPagination(user_id, limit, offset int, response *models.AssetReponse, errs chan error) {
	chartrepo := ChartRepository{}
	charts, err := chartrepo.GetChartsPagination(user_id, limit, offset)
	response.Charts = charts

	insightsrepo := Insightrepository{}
	insights, err := insightsrepo.GetInsightsPagination(user_id, limit, offset)
	response.Insights = insights

	audiencesrepo := AudienceRepository{}
	audiences, err := audiencesrepo.GetAudiencesPagination(user_id, limit, offset)
	response.Audiences = audiences

	errs <- err
}
