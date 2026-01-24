package services_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	mrepositories "taska-core-me-go/cmd/api/mocks/repositories"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/services"
	"taska-core-me-go/cmd/api/utils"
	"testing"
)

func TestSkillsServices_Search(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockISkillsRepository := mrepositories.NewMockISkillsRepository(ctrl)
	testServices := &services.SkillsServices{
		SkillsRepository: mockISkillsRepository,
	}

	entity := getParamsSkillsSearch()
	response := getSkillsResponse()

	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		mockISkillsRepository.EXPECT().FindBy(gomock.Any(), entity).Return(response, nil)

		response, err := testServices.Search(ctx, entity)
		assert.NoError(err)
		assert.Equal(response[0].ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		mockISkillsRepository.EXPECT().FindBy(gomock.Any(), entity).Return(response, testError)

		_, err := testServices.Search(ctx, entity)
		assert.Error(err)
	})

}

func TestSkillsServices_List(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockISkillsRepository := mrepositories.NewMockISkillsRepository(ctrl)
	testServices := &services.SkillsServices{
		SkillsRepository: mockISkillsRepository,
	}

	response := getSkillsResponse()

	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		mockISkillsRepository.EXPECT().FindAll(gomock.Any()).Return(response, nil)

		response, err := testServices.List(ctx)
		assert.NoError(err)
		assert.Equal(response[0].ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		mockISkillsRepository.EXPECT().FindAll(gomock.Any()).Return(response, testError)

		_, err := testServices.List(ctx)
		assert.Error(err)
	})

}

func TestSkillsServices_Save(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockISkillsRepository := mrepositories.NewMockISkillsRepository(ctrl)
	testServices := &services.SkillsServices{
		SkillsRepository: mockISkillsRepository,
	}

	response := getSkills()
	request := getParamsSkillsSave()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		skills := models.Skills{
			Name:                 request.Name,
			Slug:                 request.Slug,
			Description:          request.Description,
			AvgPriceEstimate:     request.AvgPriceEstimate,
			RequiresVerification: request.RequiresVerification,
			RiskLevel:            request.RiskLevel,
			IsActive:             request.IsActive,
		}

		mockISkillsRepository.EXPECT().Upsert(gomock.Any(), skills).Return(response, nil)

		response, err := testServices.Save(ctx, request)
		assert.NoError(err)
		assert.Equal(response.ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		skills := models.Skills{
			Name:                 request.Name,
			Slug:                 request.Slug,
			Description:          request.Description,
			AvgPriceEstimate:     request.AvgPriceEstimate,
			RequiresVerification: request.RequiresVerification,
			RiskLevel:            request.RiskLevel,
			IsActive:             request.IsActive,
		}

		mockISkillsRepository.EXPECT().Upsert(gomock.Any(), skills).Return(response, testError)

		_, err := testServices.List(ctx)
		assert.Error(err)
	})

}

func TestSkillsServices_Update(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockISkillsRepository := mrepositories.NewMockISkillsRepository(ctrl)
	testServices := &services.SkillsServices{
		SkillsRepository: mockISkillsRepository,
	}

	request := getParamsSkillsSave()
	id := int64(1)
	response := getSkills()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		skills := models.Skills{
			ID:                   id,
			Name:                 request.Name,
			Slug:                 request.Slug,
			Description:          request.Description,
			AvgPriceEstimate:     request.AvgPriceEstimate,
			RequiresVerification: request.RequiresVerification,
			RiskLevel:            request.RiskLevel,
			IsActive:             request.IsActive,
		}

		mockISkillsRepository.EXPECT().Upsert(gomock.Any(), skills).Return(response, nil)

		response, err := testServices.Update(ctx, id, request)
		assert.NoError(err)
		assert.Equal(response.ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		skills := models.Skills{
			ID:                   id,
			Name:                 request.Name,
			Slug:                 request.Slug,
			Description:          request.Description,
			AvgPriceEstimate:     request.AvgPriceEstimate,
			RequiresVerification: request.RequiresVerification,
			RiskLevel:            request.RiskLevel,
			IsActive:             request.IsActive,
		}

		mockISkillsRepository.EXPECT().Upsert(gomock.Any(), skills).Return(response, testError)

		_, err := testServices.Update(ctx, id, request)
		assert.Error(err)
	})

}

func getParamsSkillsSearch() models.ParamsSkillsSearch {
	return models.ParamsSkillsSearch{
		ID: 1,
	}
}

func getSkillsResponse() []models.Skills {
	return []models.Skills{
		{
			ID: 1,
		},
	}
}

func getParamsSkillsSave() models.ParamsSkillsSave {
	return models.ParamsSkillsSave{
		Name: "test",
	}
}

func getSkills() models.Skills {
	return models.Skills{
		ID:                   1,
		Name:                 "instalación de televisores",
		Slug:                 "instalacion-televisores",
		Description:          "Montaje e instalación de televisores en pared o soporte, configuración básica incluida",
		AvgPriceEstimate:     60000,
		RequiresVerification: true,
		RiskLevel:            2,
		IsActive:             true,
	}
}
