package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

//go:generate mockgen -destination=../mocks/validator/$GOFILE -package=mvalidator -source=./$GOFILE

type IValidator interface {
	Validate(input interface{}) error
}

type Validator struct {
}

func (s *Validator) Validate(input interface{}) error {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	if err := v.Struct(input); err != nil {
		return err
	}

	return nil
}
