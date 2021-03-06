package common

import (
	"crypto/sha1"
	"encoding/hex"
	plugins_pkg "finance/plugins"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 生成一个指定位数的数字码
func GenerateVerifyCode(length int) string {
	var build strings.Builder
	// 加入随机因子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := &length; *i > 0; *i-- {
		rand_number := r.Intn(9)
		string_number := strconv.Itoa(rand_number)
		build.WriteString(string_number)
	}
	return build.String()
}

func SHA1(text string) string {
	salt_text := text + plugins_pkg.Config.SecretKey
	ctx := sha1.New()
	ctx.Write([]byte(salt_text))
	return hex.EncodeToString(ctx.Sum(nil))
}
