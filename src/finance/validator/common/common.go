package common

type SMSSend struct {
	Phone string `valid:"required" form:"phone"`
	Genre string `valid:"required,in(registered|edit_password)~短信类型错误" form:"genre"`
}
