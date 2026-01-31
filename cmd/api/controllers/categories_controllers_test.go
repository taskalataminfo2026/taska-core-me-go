package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"taska-core-me-go/cmd/api/controllers"
	"taska-core-me-go/cmd/api/controllers/dto"
	mservices "taska-core-me-go/cmd/api/mocks/services"
	mvalidator "taska-core-me-go/cmd/api/mocks/validator"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/utils/json_mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
)

func TestCategoriesController_Search(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockICategoriesServices := mservices.NewMockICategoriesServices(ctrl)
	mockIValidator := mvalidator.NewMockIValidator(ctrl)

	testController := &controllers.CategoriesController{
		CategoriesServices: mockICategoriesServices,
		Validator:          mockIValidator,
	}

	entity := getParamsCategorySearchDto()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockICategoriesServices.EXPECT().Search(gomock.Any(), entity.ToModel()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/core/category/search", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Search(c)

		assert.NoError(err, "Search no debe devolver un error")
		assert.Equal(http.StatusOK, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockICategoriesServices.EXPECT().Search(gomock.Any(), entity.ToModel()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/core/category/search", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Search(c)

		assert.NoError(err, "Search no debe devolver un error")
		assert.Equal(http.StatusInternalServerError, rec.Code, "El código de estado debe ser 500")
	})

	t.Run("Bind_Error", func(t *testing.T) {
		e := echo.New()
		invalidJSON := `{"name": "testCategory"`

		req := httptest.NewRequest(http.MethodGet, "/v1/api/core/category/search", strings.NewReader(invalidJSON))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := testController.Search(c)

		assert.NoError(err, "Search no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

	t.Run("Validate_Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(testError)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/core/category/search", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Search(c)

		assert.NoError(err, "Search no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

}

func TestCategoriesController_List(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockICategoriesServices := mservices.NewMockICategoriesServices(ctrl)

	testController := &controllers.CategoriesController{
		CategoriesServices: mockICategoriesServices,
	}

	entity := getParamsCategorySearchDto()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockICategoriesServices.EXPECT().List(gomock.Any()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/core/category/list", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.List(c)

		assert.NoError(err, "List no debe devolver un error")
		assert.Equal(http.StatusOK, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockICategoriesServices.EXPECT().List(gomock.Any()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/core/category/list", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.List(c)

		assert.NoError(err, "List no debe devolver un error")
		assert.Equal(http.StatusInternalServerError, rec.Code, "El código de estado debe ser 500")
	})

}

func TestCategoriesController_Save(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockICategoriesServices := mservices.NewMockICategoriesServices(ctrl)
	mockIValidator := mvalidator.NewMockIValidator(ctrl)

	testController := &controllers.CategoriesController{
		CategoriesServices: mockICategoriesServices,
		Validator:          mockIValidator,
	}

	entity := getParamsCategorySaveDto()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_save_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockICategoriesServices.EXPECT().Save(gomock.Any(), entity.ToModel()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/category/save", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Save(c)

		assert.NoError(err, "Save no debe devolver un error")
		assert.Equal(http.StatusOK, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_save_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockICategoriesServices.EXPECT().Save(gomock.Any(), entity.ToModel()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/category/save", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Save(c)

		assert.NoError(err, "Save no debe devolver un error")
		assert.Equal(http.StatusInternalServerError, rec.Code, "El código de estado debe ser 500")
	})

	t.Run("BindAndValidate_Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		mockIValidator.EXPECT().Validate(&entity).Return(testError)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/category/save", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Save(c)

		assert.NoError(err, "Save no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

}

func TestCategoriesController_Update(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockICategoriesServices := mservices.NewMockICategoriesServices(ctrl)
	mockIValidator := mvalidator.NewMockIValidator(ctrl)

	testController := &controllers.CategoriesController{
		CategoriesServices: mockICategoriesServices,
		Validator:          mockIValidator,
	}

	entity := getParamsCategorySaveDto()
	path := getParamsCategoryRequestDTO()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_save_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockICategoriesServices.EXPECT().Update(gomock.Any(), path.ID, entity.ToModel()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/category/1", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusOK, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_save_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockICategoriesServices.EXPECT().Update(gomock.Any(), path.ID, entity.ToModel()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/category/1", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusInternalServerError, rec.Code, "El código de estado debe ser 500")
	})

	t.Run("BindAndValidate_Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		mockIValidator.EXPECT().Validate(&entity).Return(testError)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/category/1", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

	t.Run("ParseIDFromParam_Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := models.Category{}
		err = json.Unmarshal(json_mocks.GetJSONFile("category", "category_save_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/category/a", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("a")

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

}

func getParamsCategorySearchDto() dto.ParamsCategorySearchDto {
	return dto.ParamsCategorySearchDto{}
}

func getParamsCategorySaveDto() dto.ParamsCategorySaveDto {
	return dto.ParamsCategorySaveDto{
		RootID:      0,
		ParentID:    0,
		Name:        "Hogar y Mantenimiento",
		Slug:        "hogar-mantenimiento",
		Description: "Servicios relacionados con el hogar y mantenimiento general",
		Icon:        "home",
		IsActive:    true,
		SortOrder:   1,
	}
}

func getParamsCategoryRequestDTO() dto.ParamsCategoryRequestDTO {
	return dto.ParamsCategoryRequestDTO{
		ID: 1,
	}
}
