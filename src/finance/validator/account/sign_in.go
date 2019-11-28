package account

import (
	"errors"
	"finance/models"
	finance_model "finance/models/finance"
	"finance/plugins/common"
)

type SignInForm struct {
	Phone    string `validate:"required" json:"phone" form:"phone" error_message:"密码~required:为必填项"`
	Password string `validate:"required,min=1,max=24" json:"password" form:"password" error_message:"密码~required:为必填项;min:最短长度为1位;max:最大长度为24位"`
}

func (form *SignInForm) Valid() (finance_model.Finance, error) {
	var finance finance_model.Finance

	models.DB.First(&finance, "phone=?", form.Phone)

	if finance.ID == 0 {
		return finance, errors.New("账号错误")
	}

	if finance.Password != common.SHA1(form.Password) {
		return finance, errors.New("账号错误")
	}

	return finance, nil
}
