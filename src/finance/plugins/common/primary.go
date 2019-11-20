package common

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Export struct {
	ErrorCode int                    `json:"error_code"`
	Data      map[string]interface{} `json:"data"`
	Message   string                 `json:"message"`
}

func (export *Export) ApiExport() *Export {
	if export.Data == nil {
		export.Data = map[string]interface{}{}
	}
	return export
}

// 生成一个指定位数的数据码
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
