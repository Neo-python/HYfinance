package account

import (
	"errors"
	"finance/models"
	"finance/models/finance"
	"finance/plugins/redis"
)

type EditPasswordForm struct {
	Phone          string `validate:"required" json:"phone" form:"phone" error_message:"手机号~required:为必填项"`
	Password       string `validate:"required,min=1,max=24" json:"password" form:"password" error_message:"密码~required:为必填项;min:最短长度为1位;max:最大长度为24位"`
	RepeatPassword string `validate:"required,eqfield=Password" json:"repeat_password" form:"repeat_password" error_message:"重复密码~required:为必填项;eqfield:两次输入不一致,请检查后重试!"`
	Code           string `validate:"required,len=4" json:"code" form:"code" error_message:"验证码~required:为必填项;len:长度为4位"`
	AccountFormModel
}

func (form *EditPasswordForm) Valid() (finance.Finance, error) {

	var finance finance.Finance

	redis_code, err := redis.Get(form.RedisCodeKey("edit_password", form.Phone))

	if err != nil {
		return finance, errors.New("验证码错误")
	}

	if redis_code != form.Code {
		return finance, errors.New("验证码错误")
	}

	if form.Password != form.RepeatPassword {
		return finance, errors.New("两次密码输入不一致,请检查后重新尝试")
	}

	models.DB.First(&finance, "phone=?", form.Phone)
	return finance, nil
}
