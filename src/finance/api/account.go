package api

import (
	"finance/models"
	plugins "finance/plugins/common"
	"finance/plugins/common/structs_copy"
	"finance/validator/account"
	"github.com/gin-gonic/gin"
)

// 注册账号
func Registered(context *gin.Context) {

	var registered account.Registered
	context.BindJSON(&registered)

	//if _, err := govalidator.ValidateStruct(&registered); err != nil {
	//	// 表单验证失败,接口返回错误信息
	//	plugins.ApiExport.Error(context, err)
	//	return
	//}

	if _, err := registered.Valid(); err != nil {
		// 表单验证失败,接口返回错误信息
		plugins.ApiExport.Error(context, err)
		return
	}

	finance := models.Finance{Name: registered.Name, Phone: registered.Phone, Password: registered.Password}
	models.DB.Create(&finance)

	plugins.ApiExport.SetData("finance", structs_copy.Map(&finance))
	plugins.ApiExport.ApiExport(context)
}
