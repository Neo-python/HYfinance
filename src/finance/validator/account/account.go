package account

type Registered struct {
	Name     string `form:"name"`
	Phone    string `valid:"required" json:"phone" form:"phone"`
	Password string `valid:"required,stringlength(1|2)~密码长度错误,alphanum~只允许字母和数字" json:"password" form:"password"`
}
