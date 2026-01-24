package utils_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"net/http"
	"net/http/httptest"
	"strings"
	dto2 "taska-core-me-go/cmd/api/controllers/dto"
	mvalidator "taska-core-me-go/cmd/api/mocks/validator"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/utils"
	"testing"
)

func TestBuildLoginResponse(t *testing.T) {
	t.Run("should build a valid LoginResponse", func(t *testing.T) {
		params := models.LoginResponseParams{
			AccessToken:  "access-token-123",
			RefreshToken: "refresh-token-456",
			Message:      "Login successful",
		}

		result := utils.BuildLoginResponse(params)

		assert.Equal(t, params.AccessToken, result.AccessToken, "AccessToken debe coincidir")
		assert.Equal(t, params.RefreshToken, result.RefreshToken, "RefreshToken debe coincidir")
		assert.Equal(t, params.Message, result.Message, "Message debe coincidir")
	})

	t.Run("should handle empty fields gracefully", func(t *testing.T) {
		params := models.LoginResponseParams{}

		result := utils.BuildLoginResponse(params)

		assert.Empty(t, result.AccessToken)
		assert.Empty(t, result.RefreshToken)
		assert.Empty(t, result.Message)
	})
}

func TestBindAndValidate(t *testing.T) {

	t.Run("should bind and validate successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		e := echo.New()
		body := `{"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6"}`
		req := httptest.NewRequest(http.MethodPost, "/auth/refresh", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		dto := new(dto2.SkillsDto)
		mockValidator := mvalidator.NewMockIValidator(ctrl)

		mockValidator.EXPECT().Validate(gomock.Any()).Return(nil)

		err := utils.BindAndValidate(c, mockValidator, dto)

		assert.NoError(t, err)
	})

	t.Run("should return error when JSON is invalid", func(t *testing.T) {
		e := echo.New()
		body := `{"refresh_token":}` // JSON malformado
		req := httptest.NewRequest(http.MethodPost, "/auth/refresh", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		dto := new(dto2.SkillsDto)

		// No se usa validator porque Bind falla antes
		err := utils.BindAndValidate(c, nil, dto)

		assert.Error(t, err, "Debe retornar error de parseo JSON")

		switch e := err.(type) {
		case response_capture.ErrorME:
			assert.Equal(t, http.StatusBadRequest, e.Code)
			assert.Contains(t, e.Message, "Formato de solicitud inválido")
		case *response_capture.ErrorME:
			assert.Equal(t, http.StatusBadRequest, e.Code)
			assert.Contains(t, e.Message, "Formato de solicitud inválido")
		default:
			t.Fatalf("Error inesperado de tipo: %T, valor: %v", err, err)
		}
	})

	t.Run("should return error when validation fails (empty token)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		e := echo.New()
		body := `{"refresh_token":""}`
		req := httptest.NewRequest(http.MethodPost, "/auth/refresh", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		dto := new(dto2.SkillsDto)
		mockValidator := mvalidator.NewMockIValidator(ctrl)

		// Simulamos error de validación
		mockValidator.EXPECT().Validate(gomock.Any()).Return(errors.New("refresh_token is required"))

		err := utils.BindAndValidate(c, mockValidator, dto)

		assert.Error(t, err, "Debe retornar error de validación")

		// Validamos el tipo del error sin causar panic
		switch e := err.(type) {
		case response_capture.ErrorME:
			assert.Equal(t, http.StatusBadRequest, e.Code)
			assert.Contains(t, e.Message, "Datos de entrada inválidos")
			assert.Contains(t, e.Message, "refresh_token is required")

		case *response_capture.ErrorME:
			assert.Equal(t, http.StatusBadRequest, e.Code)
			assert.Contains(t, e.Message, "Datos de entrada inválidos")
			assert.Contains(t, e.Message, "refresh_token is required")

		default:
			t.Fatalf("Error inesperado de tipo: %T, valor: %v", err, err)
		}
	})

}
