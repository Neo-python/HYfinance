package account

import (
	"errors"
	"finance/models"
	finance_model "finance/models/finance"
	"finance/plugins/common"
	"github.com/asaskevich/govalidator"
)

type SignInForm struct {
	Phone    string `valid:"required" json:"phone" form:"phone"`
	Password string `valid:"required,stringlength(1|24),alphanum" json:"password" form:"password"`
}

func (form *SignInForm) Valid() (finance_model.Finance, error) {
	var finance finance_model.Finance
	if _, err := govalidator.ValidateStruct(form); err != nil {
		// 表单验证失败,接口返回错误信息
		return finance, err
	}

	models.DB.First(&finance, "phone=?", form.Phone)

	if finance.ID == 0 {
		return finance, errors.New("账号错误")
	}

	if finance.Password != common.SHA1(form.Password) {
		return finance, errors.New("账号错误")
	}

	return finance, nil
}
