package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"net/http"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/validator"
)

func BuildLoginResponse(params models.LoginResponseParams) models.LoginResponse {
	return models.LoginResponse{
		AccessToken:  params.AccessToken,
		RefreshToken: params.RefreshToken,
		Message:      params.Message,
	}
}

func BindAndValidate[T any](c echo.Context, validator validator.IValidator, dto *T) error {
	if err := c.Bind(dto); err != nil {
		return response_capture.NewErrorME(CreateRequestContext(c),
			http.StatusBadRequest, err, fmt.Sprintf(constants.ErrInvalidRequestFormat, err.Error()))
	}
	if err := validator.Validate(dto); err != nil {
		return response_capture.NewErrorME(CreateRequestContext(c),
			http.StatusBadRequest, err, fmt.Sprintf(constants.ErrInvalidInputData, err.Error()))
	}
	return nil
}
