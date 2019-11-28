// 通用接口
package api

import (
	"finance/models"
	models_area "finance/models/area"
	plugins_pkg "finance/plugins"
	plugins "finance/plugins/common"
	"finance/plugins/core_sms"
	"finance/plugins/redis"
	"finance/validator"
	forms "finance/validator/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 发送短信
// 记录验证码至redis,为后续校验做准备
func SMSSend(context *gin.Context) {
	// 处理表单
	var form forms.SMSSend
	context.BindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
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
		plugins.ApiExport(context).Error(5002, "短信接口调用失败")
		return
	} else {
		println(result)
	}

	plugins.ApiExport(context).ApiExport()
}

// 地区信息查询
func QueryArea(context *gin.Context) {
	var form forms.QueryAreaForm
	context.Bind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}
	form.Valid()

	var areas []models_area.Area
	models.DB.Where("superior_id=?", form.SuperiorId).Find(&areas)

	export := plugins.ApiExport(context)

	export.SetData("areas", areas)
	export.ApiExport()
}
