package jwt_auth

import (
	"finance/plugins/redis"
	"fmt"
)

// 校验token是否被刷新
func (claims *CustomClaims) AuthToken() bool {
	redis_key := fmt.Sprint("FinanceIat_", claims.Phone)
	redis_iat, _ := redis.Get(redis_key)
	if claims.Iat != redis_iat {
		return false
	} else {
		return true
	}
}

func (claims *CustomClaims) Clear() {
	redis_key := claims.RedisKey()
	redis.Delete(redis_key)
}

func (claims *CustomClaims) RedisKey() string {
	return fmt.Sprint("FinanceIat_", claims.Phone)
}
