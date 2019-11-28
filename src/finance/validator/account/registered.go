package account

import (
	"errors"
	"finance/plugins/redis"
)

type RegisteredForm struct {
	Name     string `validate:"required,min=1,max=6" form:"name" error_message:"姓名~required:为必填项"`
	Phone    string `validate:"required" json:"phone" form:"phone" error_message:"手机号~required:为必填项"`
	Password string `validate:"required,min=1,max=24" json:"password" form:"password" error_message:"密码~required:为必填项;min:最短长度为1位;max:最大长度为24位"`
	Code     string `validate:"required,len=4" json:"code" form:"code" error_message:"验证码~required:为必填项;len:长度为4位"`
	AccountFormModel
}

func (form *RegisteredForm) Valid() (bool, error) {

	redis_code, err := redis.Get(form.RedisCodeKey("registered", form.Phone))

	if err != nil {
		return false, errors.New("验证码错误")
	}

	if redis_code != form.Code {
		return false, errors.New("验证码错误")
	}

	return true, nil
}
