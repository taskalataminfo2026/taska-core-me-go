package services_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	mrepositories "taska-core-me-go/cmd/api/mocks/repositories"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/services"
	"taska-core-me-go/cmd/api/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSkillsCategoriesServices_Save(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockISkillsCategoriesRepository := mrepositories.NewMockISkillsCategoriesRepository(ctrl)
	testServices := &services.SkillsCategoriesServices{
		SkillsCategoriesRepository: mockISkillsCategoriesRepository,
	}

	request := getParamsSkillsCategorySave()
	response := getSkillsCategory()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		skillCategory := models.SkillCategory{
			SkillID:    request.SkillID,
			CategoryID: request.CategoryID,
			IsPrimary:  request.IsPrimary,
			IsActive:   request.IsActive,
		}

		mockISkillsCategoriesRepository.EXPECT().Upsert(gomock.Any(), skillCategory).Return(response, nil)

		response, err := testServices.Save(ctx, request)
		assert.NoError(err)
		assert.Equal(response.ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		skillCategory := models.SkillCategory{
			SkillID:    request.SkillID,
			CategoryID: request.CategoryID,
			IsPrimary:  request.IsPrimary,
			IsActive:   request.IsActive,
		}

		mockISkillsCategoriesRepository.EXPECT().Upsert(gomock.Any(), skillCategory).Return(response, testError)

		_, err := testServices.Save(ctx, request)
		assert.Error(err)
	})

}

func TestSkillsCategoriesServices_Update(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockISkillsCategoriesRepository := mrepositories.NewMockISkillsCategoriesRepository(ctrl)
	testServices := &services.SkillsCategoriesServices{
		SkillsCategoriesRepository: mockISkillsCategoriesRepository,
	}

	request := getParamsSkillsCategorySave()
	id := int64(1)
	response := getSkillsCategory()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		mockISkillsCategoriesRepository.EXPECT().FirstBy(gomock.Any(), models.ParamsSkillsCategorySearch{ID: id}).Return(response, nil)

		skillCategory := models.SkillCategory{
			ID:         id,
			SkillID:    request.SkillID,
			CategoryID: request.CategoryID,
			IsPrimary:  request.IsPrimary,
			IsActive:   request.IsActive,
		}

		mockISkillsCategoriesRepository.EXPECT().Upsert(gomock.Any(), skillCategory).Return(response, nil)

		response, err := testServices.Update(ctx, id, request)
		assert.NoError(err)
		assert.Equal(response.ID, int64(1))
	})

	t.Run("Error_FirstBy", func(t *testing.T) {

		mockISkillsCategoriesRepository.EXPECT().FirstBy(gomock.Any(), models.ParamsSkillsCategorySearch{ID: id}).Return(response, testError)

		_, err := testServices.Update(ctx, id, request)
		assert.Error(err)
	})

	t.Run("Error_Upsert", func(t *testing.T) {

		mockISkillsCategoriesRepository.EXPECT().FirstBy(gomock.Any(), models.ParamsSkillsCategorySearch{ID: id}).Return(response, nil)

		skillCategory := models.SkillCategory{
			ID:         id,
			SkillID:    request.SkillID,
			CategoryID: request.CategoryID,
			IsPrimary:  request.IsPrimary,
			IsActive:   request.IsActive,
		}

		mockISkillsCategoriesRepository.EXPECT().Upsert(gomock.Any(), skillCategory).Return(response, testError)

		_, err := testServices.Update(ctx, id, request)
		assert.Error(err)
	})

}

func getParamsSkillsCategorySave() models.ParamsSkillsCategorySave {
	return models.ParamsSkillsCategorySave{
		SkillID:    10,
		CategoryID: 20,
		IsPrimary:  true,
		IsActive:   true,
	}
}

func getSkillsCategory() models.SkillCategory {
	return models.SkillCategory{
		ID:         1,
		SkillID:    10,
		CategoryID: 20,
		IsPrimary:  true,
		IsActive:   true,
	}
}
