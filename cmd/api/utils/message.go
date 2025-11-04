package utils

import (
	"net/http"
	"taska-core-me-go/cmd/api/models"
)

func NewMessage[T any](data T) models.Message {
	return models.Message{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}
}
