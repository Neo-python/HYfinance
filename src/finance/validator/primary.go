package validator

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

var Valid *validator.Validate

func init() {
	Valid = validator.New()
	Valid.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("error_message")
	})
}
