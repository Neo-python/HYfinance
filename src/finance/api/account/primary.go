package account

import (
	"encoding/json"
	"finance/models"
	plugins "finance/plugins/common"
	"finance/plugins/jwt_auth"
	"finance/plugins/redis"
	"finance/validator"
	"finance/validator/account"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 注册账号
func Registered(context *gin.Context) {

	var form account.RegisteredForm
	context.BindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		// 表单验证失败,接口返回错误信息
		plugins.ApiExport(context).FormError(err)
		return
	}

	if _, err := form.Valid(); err != nil {
		// 表单验证失败,接口返回错误信息
		plugins.ApiExport(context).Error(5011, err.Error())
		return
	}

	finance := form.GetFinance()
	finance.Name = form.Name
	finance.Password = plugins.SHA1(form.Password)

	if err := models.DB.Save(&finance).Error; err != nil {
		plugins.ApiExport(context).Error(5010, "手机号已被注册")
		return
	}

	export := plugins.ApiExport(context)
	export.SetData("finance", finance.ToJson())
	export.ApiExport()

	// 注册完成后执行清理工作
	defer form.Complete(form.RedisCodeKey("registered", form.Phone))
}

// 修改密码
func EditPassword(context *gin.Context) {
	var form account.EditPasswordForm
	context.BindJSON(&form)

	if err := validator.Valid.Struct(form); err != nil {
		// 表单验证失败,接口返回错误信息
		plugins.ApiExport(context).FormError(err)
		return
	}

	finance, err := form.Valid()
	if err != nil {
		// 表单验证失败,接口返回错误信息
		plugins.ApiExport(context).Error(5011, err.Error())
		return
	}

	finance.Password = plugins.SHA1(form.Password)
	models.DB.Save(&finance)
	export := plugins.ApiExport(context)
	export.ApiExport()

	// 删除旧token记录
	redis.Delete(fmt.Sprintf("FinanceIat_%s", form.Phone))

	// 操作完成后执行清理工作
	defer form.Complete(form.RedisCodeKey("edit_password", form.Phone))

}

// 登录
func SignIn(context *gin.Context) {
	var form account.SignInForm
	context.BindJSON(&form)

	if err := validator.Valid.Struct(form); err != nil {
		// 表单验证失败,接口返回错误信息
		plugins.ApiExport(context).FormError(err)
		return
	}

	finance, err := form.Valid()
	if err != nil {
		plugins.ApiExport(context).Error(5011, err.Error())
		return
	}

	export := plugins.ApiExport(context)
	export.SetData("token", finance.Token())
	export.SetData("finance", finance.ToJson())
	export.ApiExport()
	return
}

// 登出
func SignOut(context *gin.Context) {

	claims, err := jwt_auth.GetClaims(context)
	if err != nil {
		plugins.ApiExport(context).Error(4005, err.Error())
		return
	}
	claims.Clear()

	export := plugins.ApiExport(context)
	export.ApiExport()
}

// 获取绑定的厂家token
func GetFactoryToken(context *gin.Context) {
	claims, err := jwt_auth.GetClaims(context)
	if err != nil {
		plugins.ApiExport(context).Error(4005, err.Error())
		return
	}
	token, err := plugins.CoreGetFactoryToken(claims.FactoryUuid)

	detail_map := make(map[string]interface{})
	json.Unmarshal([]byte(token), &detail_map)

	if detail_map["error_code"].(float64) != 0 {
		plugins.ApiExport(context).Error(5011, detail_map["message"].(string))
		return
	}

	export := plugins.ApiExport(context)
	export.SetData("factory_token", detail_map["data"].(string))
	export.ApiExport()
	return

}
