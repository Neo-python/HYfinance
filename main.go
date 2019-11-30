package main

import (
	"finance/api"
	"finance/api/account"
	"finance/api/business"
	"finance/api/common"
	_ "finance/models"
	_ "finance/models/init"
	"finance/plugins/jwt_auth"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := gin.Default()

	// 无需权限验证的接口
	open := router.Group("")
	{
		open.POST("registered", account.Registered)
		open.POST("edit_password", account.EditPassword)
		open.POST("/send_sms/code/", common.SMSSend)
		open.POST("sign_in", account.SignIn)
	}

	// 需要权限验证的接口
	auth := router.Group("", jwt_auth.JWTAuth())
	{
		auth.GET("test", api.Test)
		auth.GET("sign_out", account.SignOut)
		auth.GET("query_area", common.QueryArea)
		auth.POST("add_order", business.AddOrder)
		auth.GET("query_receiver", business.QueryReceiver)
		auth.GET("query_sender", business.QuerySender)
	}

	router.Run("127.0.0.1:8095")
}
