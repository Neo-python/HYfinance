// 通用接口
package api

import (
	plugins_pkg "finance/plugins"
	plugins "finance/plugins/common"
	"finance/plugins/core_sms"
	"finance/plugins/redis"
	validator "finance/validator/common"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// 发送短信
// 记录验证码至redis,为后续校验做准备
func SMSSend(context *gin.Context) {
	// 处理表单
	var form validator.SMSSend
	context.Bind(&form)

	if _, err := govalidator.ValidateStruct(&form); err != nil {
		plugins.ApiExport(context).Error(1001, "表单验证失败", err)
		return
	}

	// 发送短信前的准备工作
	code := plugins.GenerateVerifyCode(4)
	key := fmt.Sprintf(plugins_pkg.Config.SMSCodeGenre[form.Genre], form.Phone)

	redis.Set(key, code, 30000)

	// 发送短信
	sms := core_sms.SMS{&core_sms.Phone{form.Phone}, &core_sms.Genre{form.Genre}}
	// 接口返回
	if result, err := sms.Send(code); err != nil {
		plugins.ApiExport(context).Error(5000, "短信接口调用失败", err)
		return
	} else {
		println(result)
	}

	plugins.ApiExport(context).ApiExport()
}
