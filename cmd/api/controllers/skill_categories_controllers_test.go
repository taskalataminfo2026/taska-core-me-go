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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
)

func TestSkillsCategoriesController_Save(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockISkillsCategoriesServices := mservices.NewMockISkillsCategoriesServices(ctrl)
	mockIValidator := mvalidator.NewMockIValidator(ctrl)

	testController := &controllers.SkillsCategoriesController{
		SkillsCategoriesServices: mockISkillsCategoriesServices,
		Validator:                mockIValidator,
	}

	entity := getParamsSkillsCategorySaveDto()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := models.SkillCategory{
			ID:         1,
			SkillID:    entity.SkillID,
			CategoryID: entity.CategoryID,
			IsPrimary:  entity.IsPrimary,
			IsActive:   entity.IsActive,
		}

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockISkillsCategoriesServices.EXPECT().Save(gomock.Any(), entity.ToModel()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills-categories/save", bytes.NewReader(reqBody))
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

		expectedResponse := models.SkillCategory{}

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockISkillsCategoriesServices.EXPECT().Save(gomock.Any(), entity.ToModel()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills-categories/save", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Save(c)

		assert.NoError(err, "Save no debe devolver un error")
		assert.Equal(http.StatusInternalServerError, rec.Code, "El código de estado debe ser 500")
	})

	t.Run("Bind_Error", func(t *testing.T) {
		e := echo.New()
		invalidJSON := `{"skill_id": 1`

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills-categories/save", strings.NewReader(invalidJSON))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := testController.Save(c)

		assert.NoError(err, "Save no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

	t.Run("Validate_Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		mockIValidator.EXPECT().Validate(&entity).Return(testError)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills-categories/save", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Save(c)

		assert.NoError(err, "Save no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

}

func TestSkillsCategoriesController_Update(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockISkillsCategoriesServices := mservices.NewMockISkillsCategoriesServices(ctrl)
	mockIValidator := mvalidator.NewMockIValidator(ctrl)

	testController := &controllers.SkillsCategoriesController{
		SkillsCategoriesServices: mockISkillsCategoriesServices,
		Validator:                mockIValidator,
	}

	entity := getParamsSkillsCategorySaveDto()
	path := getParamsSkillsCategoryRequestDTO()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := models.SkillCategory{
			ID:         path.ID,
			SkillID:    entity.SkillID,
			CategoryID: entity.CategoryID,
			IsPrimary:  entity.IsPrimary,
			IsActive:   entity.IsActive,
		}

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockISkillsCategoriesServices.EXPECT().Update(gomock.Any(), path.ID, entity.ToModel()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/skills-categories/1", bytes.NewReader(reqBody))
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

		expectedResponse := models.SkillCategory{}

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockISkillsCategoriesServices.EXPECT().Update(gomock.Any(), path.ID, entity.ToModel()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/skills-categories/1", bytes.NewReader(reqBody))
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

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/skills-categories/1", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

	t.Run("ParseIDFromParam_Error_InvalidID", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/skills-categories/abc", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("abc")

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

	t.Run("ParseIDFromParam_Error_EmptyID", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/skills-categories/", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("")

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

	t.Run("ParseIDFromParam_Error_NegativeID", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/core/skills-categories/-1", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("-1")

		err = testController.Update(c)

		assert.NoError(err, "Update no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 400")
	})

}

func getParamsSkillsCategorySaveDto() dto.ParamsSkillsCategorySaveDto {
	return dto.ParamsSkillsCategorySaveDto{
		SkillID:    1,
		CategoryID: 2,
		IsPrimary:  true,
		IsActive:   true,
	}
}

func getParamsSkillsCategoryRequestDTO() dto.ParamsSkillsCategoryRequestDTO {
	return dto.ParamsSkillsCategoryRequestDTO{
		ID: 1,
	}
}
