package account

import (
	"finance/models"
	finance_model "finance/models/finance"
	plugins "finance/plugins/common"
	"finance/plugins/jwt_auth"
	"finance/validator"
	"finance/validator/account"
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

	finance := finance_model.Finance{Name: form.Name, Phone: form.Phone, Password: plugins.SHA1(form.Password)}
	if err := models.DB.Create(&finance).Error; err != nil {
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
}

// 登出
func SignOut(context *gin.Context) {

	result, status := context.Get("claims")

	if status != true {
		plugins.ApiExport(context).Error(4005, "当前状态:未登录")
		return
	}

	// 清理redis token.iat 缓存
	claims, status := result.(*jwt_auth.CustomClaims)
	claims.Clear()

	export := plugins.ApiExport(context)
	export.ApiExport()
}