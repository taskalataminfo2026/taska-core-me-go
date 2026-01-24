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

func TestSkillsServices_SkillsSearch(t *testing.T) {
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

		response, err := testServices.SkillsSearch(ctx, entity)
		assert.NoError(err)
		assert.Equal(response[0].ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		mockISkillsRepository.EXPECT().FindBy(gomock.Any(), entity).Return(response, testError)

		_, err := testServices.SkillsSearch(ctx, entity)
		assert.Error(err)
	})

}

func TestSkillsServices_SkillsList(t *testing.T) {
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

		response, err := testServices.SkillsList(ctx)
		assert.NoError(err)
		assert.Equal(response[0].ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		mockISkillsRepository.EXPECT().FindAll(gomock.Any()).Return(response, testError)

		_, err := testServices.SkillsList(ctx)
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
