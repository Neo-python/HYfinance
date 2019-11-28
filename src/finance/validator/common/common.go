package common

import "fmt"

type SMSSend struct {
	Phone string `validate:"required" form:"phone"`
	Genre string `validate:"required,oneof=registered edit_password" form:"genre"`
}

type QueryAreaForm struct {
	SuperiorId *int `validate:"required" form:"superior_id" json:"superior_id" error_message:"上级地区编号~required:此为必填;"`
}

func (form QueryAreaForm) Valid() {
	fmt.Println(form.SuperiorId == nil)
}
