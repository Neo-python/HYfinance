package account

import (
	"errors"
	"finance/plugins/redis"
	"github.com/asaskevich/govalidator"
)

type RegisteredForm struct {
	Name     string `valid:"required,stringlength(1|6)" form:"name"`
	Phone    string `valid:"required" json:"phone" form:"phone"`
	Password string `valid:"required,stringlength(1|24),alphanum" json:"password" form:"password"`
	Code     string `valid:"required, stringlength(4|4)" json:"code" form:"code"`
	AccountFormModel
}

func (form *RegisteredForm) Valid() (bool, error) {

	if _, err := govalidator.ValidateStruct(form); err != nil {
		// 表单验证失败,接口返回错误信息
		return false, err
	}

	redis_code, err := redis.Get(form.RedisCodeKey("registered", form.Phone))

	if err != nil {
		return false, errors.New("code:验证码错误")
	}

	if redis_code != form.Code {
		return false, errors.New("phone:验证码错误")
	}

	return true, nil
}
