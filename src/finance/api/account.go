package api

import (
	"finance/models"
	plugins "finance/plugins/common"
	"finance/validator/account"
	//"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var finance = models.Finance{}
	finance.Name = registered.Name
	finance.Phone = registered.Phone
	finance.Password = registered.Password
	//finance := models.Finance{Name: registered.Name, Phone: registered.Phone, Password: registered.Password}
	models.DB.Create(finance)

	context.JSON(http.StatusOK, gin.H{"status": 1, "name": registered.Name})
}
