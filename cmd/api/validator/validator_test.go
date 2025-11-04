package validator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"taska-core-me-go/cmd/api/validator"
)

type ExampleDTO struct {
	UserName string `json:"user_name" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func TestValidator_Validate(t *testing.T) {
	assert := assert.New(t)
	v := &validator.Validator{}

	t.Run("Ok", func(t *testing.T) {
		input := ExampleDTO{
			UserName: "usuario@ejemplo.com",
			Password: "123456",
		}

		err := v.Validate(input)
		assert.NoError(err, "No debe devolver error con datos válidos")
	})

	t.Run("Error - Campos faltantes", func(t *testing.T) {
		input := ExampleDTO{
			UserName: "",
			Password: "",
		}

		err := v.Validate(input)
		assert.Error(err, "Debe devolver error cuando faltan campos requeridos")
	})
}
