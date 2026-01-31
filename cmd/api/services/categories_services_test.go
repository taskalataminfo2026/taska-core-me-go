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

func TestCategoriesServices_Search(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockICategoriesRepository := mrepositories.NewMockICategoriesRepository(ctrl)
	testServices := &services.CategoriesServices{
		CategoriesRepository: mockICategoriesRepository,
	}

	entity := getParamsCategorySearch()
	response := getCategoriesResponse()

	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		mockICategoriesRepository.EXPECT().FindBy(gomock.Any(), entity).Return(response, nil)

		response, err := testServices.Search(ctx, entity)
		assert.NoError(err)
		assert.Equal(response[0].ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		mockICategoriesRepository.EXPECT().FindBy(gomock.Any(), entity).Return(response, testError)

		_, err := testServices.Search(ctx, entity)
		assert.Error(err)
	})

}

func TestCategoriesServices_List(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockICategoriesRepository := mrepositories.NewMockICategoriesRepository(ctrl)
	testServices := &services.CategoriesServices{
		CategoriesRepository: mockICategoriesRepository,
	}

	response := getCategoriesResponse()

	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		mockICategoriesRepository.EXPECT().FindAll(gomock.Any()).Return(response, nil)

		response, err := testServices.List(ctx)
		assert.NoError(err)
		assert.Equal(response[0].ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		mockICategoriesRepository.EXPECT().FindAll(gomock.Any()).Return(response, testError)

		_, err := testServices.List(ctx)
		assert.Error(err)
	})

}

func TestCategoriesServices_Save(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockICategoriesRepository := mrepositories.NewMockICategoriesRepository(ctrl)
	testServices := &services.CategoriesServices{
		CategoriesRepository: mockICategoriesRepository,
	}

	response := getCategory()
	request := getParamsCategorySave()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		category := models.Category{
			RootID:      request.RootID,
			ParentID:    request.ParentID,
			Name:        request.Name,
			Slug:        request.Slug,
			Description: request.Description,
			Icon:        request.Icon,
			IsActive:    request.IsActive,
			SortOrder:   request.SortOrder,
		}

		mockICategoriesRepository.EXPECT().Upsert(gomock.Any(), category).Return(response, nil)

		response, err := testServices.Save(ctx, request)
		assert.NoError(err)
		assert.Equal(response.ID, int64(1))
	})

	t.Run("Error", func(t *testing.T) {

		category := models.Category{
			RootID:      request.RootID,
			ParentID:    request.ParentID,
			Name:        request.Name,
			Slug:        request.Slug,
			Description: request.Description,
			Icon:        request.Icon,
			IsActive:    request.IsActive,
			SortOrder:   request.SortOrder,
		}

		mockICategoriesRepository.EXPECT().Upsert(gomock.Any(), category).Return(response, testError)

		_, err := testServices.Save(ctx, request)
		assert.Error(err)
	})

}

func TestCategoriesServices_Update(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	mockICategoriesRepository := mrepositories.NewMockICategoriesRepository(ctrl)
	testServices := &services.CategoriesServices{
		CategoriesRepository: mockICategoriesRepository,
	}

	request := getParamsCategorySave()
	id := int64(1)
	response := getCategory()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {

		mockICategoriesRepository.EXPECT().FirstBy(gomock.Any(), models.ParamsCategorySearch{ID: id}).Return(response, nil)

		category := models.Category{
			ID:          id,
			RootID:      request.RootID,
			ParentID:    request.ParentID,
			Name:        request.Name,
			Slug:        request.Slug,
			Description: request.Description,
			Icon:        request.Icon,
			IsActive:    request.IsActive,
			SortOrder:   request.SortOrder,
		}

		mockICategoriesRepository.EXPECT().Upsert(gomock.Any(), category).Return(response, nil)

		response, err := testServices.Update(ctx, id, request)
		assert.NoError(err)
		assert.Equal(response.ID, int64(1))
	})

	t.Run("Error_FirstBy", func(t *testing.T) {

		mockICategoriesRepository.EXPECT().FirstBy(gomock.Any(), models.ParamsCategorySearch{ID: id}).Return(response, testError)

		_, err := testServices.Update(ctx, id, request)
		assert.Error(err)
	})

	t.Run("Error_Upsert", func(t *testing.T) {

		mockICategoriesRepository.EXPECT().FirstBy(gomock.Any(), models.ParamsCategorySearch{ID: id}).Return(response, nil)

		category := models.Category{
			ID:          id,
			RootID:      request.RootID,
			ParentID:    request.ParentID,
			Name:        request.Name,
			Slug:        request.Slug,
			Description: request.Description,
			Icon:        request.Icon,
			IsActive:    request.IsActive,
			SortOrder:   request.SortOrder,
		}

		mockICategoriesRepository.EXPECT().Upsert(gomock.Any(), category).Return(response, testError)

		_, err := testServices.Update(ctx, id, request)
		assert.Error(err)
	})

}

func getParamsCategorySearch() models.ParamsCategorySearch {
	return models.ParamsCategorySearch{
		ID: 1,
	}
}

func getCategoriesResponse() []models.Category {
	return []models.Category{
		{
			ID: 1,
		},
	}
}

func getParamsCategorySave() models.ParamsCategorySave {
	return models.ParamsCategorySave{
		RootID:      0,
		ParentID:    0,
		Name:        "Hogar",
		Slug:        "hogar",
		Description: "Servicios para el hogar",
		Icon:        "home",
		IsActive:    true,
		SortOrder:   1,
	}
}

func getCategory() models.Category {
	return models.Category{
		ID:          1,
		RootID:      0,
		ParentID:    0,
		Name:        "Hogar",
		Slug:        "hogar",
		Description: "Servicios para el hogar",
		Icon:        "home",
		IsActive:    true,
		SortOrder:   1,
	}
}
