package hashs_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"taska-core-me-go/cmd/api/utils/hashs"
	"testing"
)

func TestGeneratePasswordHash(t *testing.T) {
	ctx := context.Background()

	t.Run("Genera hash correctamente", func(t *testing.T) {
		password := "MiContraseñaSegura123"
		hash, err := hashs.GeneratePasswordHash(ctx, password)
		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.NotEqual(t, password, hash)
	})

	t.Run("Retorna error si bcrypt falla", func(t *testing.T) {
	})
}

func TestCheckPassword(t *testing.T) {
	ctx := context.Background()
	password := "MiContraseñaSegura123"
	hash, err := hashs.GeneratePasswordHash(ctx, password)
	assert.NoError(t, err)

	t.Run("Contraseña correcta", func(t *testing.T) {
		ok := hashs.CheckPassword(password, hash)
		assert.True(t, ok)
	})

	t.Run("Contraseña incorrecta", func(t *testing.T) {
		ok := hashs.CheckPassword("ContraseñaErrónea", hash)
		assert.False(t, ok)
	})
}
