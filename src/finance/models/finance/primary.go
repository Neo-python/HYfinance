package finance

import (
	"finance/plugins"
	"finance/plugins/jwt_auth"
	"finance/plugins/redis"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type Finance struct {
	gorm.Model
	Name     string `json:"name" gorm:"COMMENT:'名称'"`
	Phone    string `json:"phone" gorm:"unique_index;COMMENT:'手机号'"`
	Password string `structs:",remove"`
	Level    int    `json:"level" gorm:"default:1;COMMENT:'权限等级.1:普通角色 2:管理员角色'"`
}

// 生成token
// 记录iat值
func (finance *Finance) Token() string {
	iat := strconv.Itoa(int(time.Now().Unix()))
	jwt_obj := jwt_auth.NewJWT()
	token, _ := jwt_obj.CreateToken(&jwt_auth.CustomClaims{
		ID:    strconv.Itoa(int(finance.ID)),
		Name:  finance.Name,
		Phone: finance.Phone,
		Iat:   iat,
		Level: finance.Level})

	redis.Set(finance.RedisKey(), iat, plugins.Config.JWTExpired)

	return token

}

func (finance *Finance) RedisKey() string {
	return fmt.Sprintf("FinanceIat_%s", finance.Phone)
}

func (finance *Finance) ToJson() *map[string]interface{} {
	data := make(map[string]interface{})
	data["name"] = finance.Name
	data["phone"] = finance.Phone
	data["level"] = finance.Level
	return &data
}
