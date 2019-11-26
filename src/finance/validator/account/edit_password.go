package account

import (
	"errors"
	"finance/models"
	"finance/models/finance"
	"finance/plugins/redis"
	"github.com/asaskevich/govalidator"
)

type EditPasswordForm struct {
	Phone          string `valid:"required" json:"phone" form:"phone"`
	Password       string `valid:"required,stringlength(1|24),alphanum" json:"password" form:"password"`
	RepeatPassword string `valid:"required,stringlength(1|24),alphanum" json:"repeat_password" form:"repeat_password"`
	Code           string `valid:"required, stringlength(4|4)" json:"code" form:"code"`
	AccountFormModel
}

func (form *EditPasswordForm) Valid() (finance.Finance, error) {

	var finance finance.Finance

	if _, err := govalidator.ValidateStruct(form); err != nil {
		// 表单验证失败,接口返回错误信息
		return finance, err
	}

	redis_code, err := redis.Get(form.RedisCodeKey("edit_password", form.Phone))

	if err != nil {
		return finance, errors.New("code:验证码错误")
	}

	if redis_code != form.Code {
		return finance, errors.New("phone:验证码错误")
	}

	if form.Password != form.RepeatPassword {
		return finance, errors.New("两次密码输入不一致,请检查后重新尝试")
	}

	models.DB.First(&finance, "phone=?", form.Phone)
	return finance, nil
}
