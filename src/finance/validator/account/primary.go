package account

import (
	"finance/plugins"
	"finance/plugins/redis"
	"fmt"
)

type AccountFormBase interface {
	RedisCodeKey()
}

type AccountFormModel struct {
}

func (form *AccountFormModel) RedisCodeKey(genre string, phone string) string {
	// 生成redis短信验证码缓存键名
	redis_key := fmt.Sprintf(plugins.Config.SMSCodeGenre[genre], phone)
	return redis_key
}

func (form *AccountFormModel) Complete(redis_key string) {
	// 注册完成,后续工作.
	// 清理redis短信验证码
	redis.Delete(redis_key)
}
