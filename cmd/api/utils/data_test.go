package utils_test

import (
	"github.com/stretchr/testify/assert"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/utils"
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	t.Run("Debe formatear correctamente una fecha válida", func(t *testing.T) {
		fecha := time.Date(2025, 10, 15, 0, 0, 0, 0, time.UTC)
		result := utils.FormatDate(fecha)
		assert.Equal(t, "2025-10-15", result)
	})

	t.Run("Debe devolver string vacío si la fecha está vacía", func(t *testing.T) {
		var fecha time.Time
		result := utils.FormatDate(fecha)
		assert.Equal(t, "", result)
	})
}

func TestFormatDateTime(t *testing.T) {
	t.Run("Debe formatear correctamente fecha y hora", func(t *testing.T) {
		fecha := time.Date(2025, 10, 15, 22, 30, 45, 0, time.UTC)
		result := utils.FormatDateTime(fecha)
		assert.Equal(t, "2025-10-15 22:30:45", result)
	})

	t.Run("Debe devolver string vacío si la fecha está vacía", func(t *testing.T) {
		var fecha time.Time
		result := utils.FormatDateTime(fecha)
		assert.Equal(t, "", result)
	})
}

func TestFormatConstants(t *testing.T) {
	assert.Equal(t, "2006-01-02", constants.DateFormatISO)
	assert.Equal(t, "2006-01-02 15:04:05", constants.DateTimeFormatFull)
}
