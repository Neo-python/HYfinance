package account

import (
	"errors"
	"finance/models"
	models_finance "finance/models/finance"
	"finance/plugins/redis"
	"fmt"
)

type RegisteredForm struct {
	Name     string `validate:"required,min=1,max=6" form:"name" error_message:"姓名~required:为必填项"`
	Phone    string `validate:"required" json:"phone" form:"phone" error_message:"手机号~required:为必填项"`
	Password string `validate:"required,min=1,max=24" json:"password" form:"password" error_message:"密码~required:为必填项;min:最短长度为1位;max:最大长度为24位"`
	Code     string `validate:"required,len=4" json:"code" form:"code" error_message:"验证码~required:为必填项;len:长度为4位"`
	AccountFormModel
}

func (form *RegisteredForm) Valid() (bool, error) {
	redis_key := form.RedisCodeKey("registered", form.Phone)
	redis_code, err := redis.Get(redis_key)

	if err != nil {

		fmt.Println(err.Error(), redis_key)
		return false, errors.New("验证码错误")
	}

	if redis_code != form.Code {
		return false, errors.New("验证码错误")
	}

	if form.GetFinance().ID == 0 {
		return false, errors.New("此手机号不能注册成为海粤财务,请与管理员联系")
	}

	return true, nil
}

func (form *RegisteredForm) GetFinance() *models_finance.Finance {
	var finance models_finance.Finance
	models.DB.Where("phone=?", form.Phone).Where("password=?", "").First(&finance)
	return &finance
}
