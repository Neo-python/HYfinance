package account

import (
	"errors"
	"finance/plugins/redis"
	"fmt"
	"github.com/asaskevich/govalidator"
)

type Registered struct {
	Name     string `valid:"required,stringlength(1|6)" form:"name"`
	Phone    string `valid:"required" json:"phone" form:"phone"`
	Password string `valid:"required,stringlength(1|24),alphanum" json:"password" form:"password"`
	Code     string `valid:"required, stringlength(4|4)" json:"code" form:"code"`
}

func (form *Registered) Valid() (bool, error) {

	if _, err := govalidator.ValidateStruct(form); err != nil {
		// 表单验证失败,接口返回错误信息
		return false, err
	}

	redis_code, err := redis.Get(form.redis_key())

	if err != nil {
		return false, errors.New("code:验证码错误")
	}

	if redis_code != form.Code {
		return false, errors.New("phone:验证码错误")
	}

	return true, nil
}

func (form *Registered) Complete() {
	// 注册完成,后续工作.
	// 清理redis短信验证码
	redis.Delete(form.redis_key())
}

func (form *Registered) redis_key() string {
	// 生成redis短信验证码缓存键名
	return fmt.Sprintf("Registered_%s", form.Phone)
}
