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
	Name     string `json:"name"`
	Phone    string `json:"phone" gorm:"unique_index"`
	Password string `structs:",remove"`
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
		Iat:   iat})

	redis.Set(finance.RedisKey(), iat, plugins.Config.JWTExpired)

	return token

}

func (finance *Finance) RedisKey() string {
	return fmt.Sprint("FinanceIat_", finance.Phone)
}
