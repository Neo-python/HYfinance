package validator

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

var Valid *validator.Validate

type ListPage struct {
	Page  int `json:"page" form:"page" validate:"required,min=1" error_message:"页码~required:为必须填写的.;min:最小值为1"`
	Limit int `json:"limit" form:"limit" validate:"required,min=1,max=50" error_message:"limit~required:为必须填写的.;min:最小值为1;max:最大值为50" `
	Total int
}

func init() {
	Valid = validator.New()
	Valid.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("error_message")
	})
}
