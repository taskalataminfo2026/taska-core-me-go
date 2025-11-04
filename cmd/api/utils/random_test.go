package utils_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"taska-core-me-go/cmd/api/utils"
)

func TestGenerateUniqueCode_Success(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	code, err := utils.GenerateUniqueCode(ctx)
	assert.NoError(err, "no debería haber error al generar código")
	assert.Len(code, 6, "el código debe tener 6 dígitos")
}

func TestGenerateRandomCode_Success(t *testing.T) {
	assert := assert.New(t)

	code, err := utils.GenerateRandomCode()
	assert.NoError(err, "no debería haber error al generar el código")
	assert.Len(code, 6, "el código debe tener 6 dígitos")
}
