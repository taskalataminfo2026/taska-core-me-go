package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
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
)

func TestAuthService_SkillsSearch(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockISkillsServices := mservices.NewMockISkillsServices(ctrl)
	mockIValidator := mvalidator.NewMockIValidator(ctrl)

	testController := &controllers.SkillsController{
		SkillsServices: mockISkillsServices,
		Validator:      mockIValidator,
	}

	entity := getParamsSkillsSearchDto()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Skills{}
		err = json.Unmarshal(json_mocks.GetJSONFile("skills", "skills_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockISkillsServices.EXPECT().SkillsSearch(gomock.Any(), entity.ToModel()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills/search", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.SkillsSearch(c)

		assert.NoError(err, "UpdateUser no debe devolver un error")
		assert.Equal(http.StatusOK, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Skills{}
		err = json.Unmarshal(json_mocks.GetJSONFile("skills", "skills_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(nil)

		mockISkillsServices.EXPECT().SkillsSearch(gomock.Any(), entity.ToModel()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills/search", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.SkillsSearch(c)

		assert.NoError(err, "UpdateUser no debe devolver un error")
		assert.Equal(http.StatusInternalServerError, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Bind_Error", func(t *testing.T) {
		e := echo.New()
		invalidJSON := `{"userName": "testUser"`

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills/search", strings.NewReader(invalidJSON))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := testController.SkillsSearch(c)

		assert.NoError(err, "UpdateUser no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Validate_Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Skills{}
		err = json.Unmarshal(json_mocks.GetJSONFile("skills", "skills_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockIValidator.EXPECT().Validate(&entity).Return(testError)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills/search", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.SkillsSearch(c)

		assert.NoError(err, "UpdateUser no debe devolver un error")
		assert.Equal(http.StatusBadRequest, rec.Code, "El código de estado debe ser 200 OK")
	})

}

func TestAuthService_SkillsList(t *testing.T) {
	logger.Init()
	defer logger.Sync()

	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockISkillsServices := mservices.NewMockISkillsServices(ctrl)

	testController := &controllers.SkillsController{
		SkillsServices: mockISkillsServices,
	}

	entity := getParamsSkillsSearchDto()
	var testError = errors.New("api_error")

	t.Run("Ok", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Skills{}
		err = json.Unmarshal(json_mocks.GetJSONFile("skills", "skills_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockISkillsServices.EXPECT().SkillsList(gomock.Any()).Return(expectedResponse, nil).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills/lis", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.SkillsList(c)

		assert.NoError(err, "UpdateUser no debe devolver un error")
		assert.Equal(http.StatusOK, rec.Code, "El código de estado debe ser 200 OK")
	})

	t.Run("Error", func(t *testing.T) {
		e := echo.New()
		reqBody, err := json.Marshal(entity)
		assert.NoError(err, "No se pudo serializar el cuerpo de la solicitud")

		expectedResponse := []models.Skills{}
		err = json.Unmarshal(json_mocks.GetJSONFile("skills", "skills_ok.json"), &expectedResponse)
		assert.Nil(err)

		mockISkillsServices.EXPECT().SkillsList(gomock.Any()).Return(expectedResponse, testError).Times(1)

		req := httptest.NewRequest(http.MethodPost, "/v1/api/core/skills/lis", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = testController.SkillsList(c)

		assert.NoError(err, "UpdateUser no debe devolver un error")
		assert.Equal(http.StatusInternalServerError, rec.Code, "El código de estado debe ser 200 OK")
	})

}

func getParamsSkillsSearchDto() dto.ParamsSkillsSearchDto {
	return dto.ParamsSkillsSearchDto{}
}
