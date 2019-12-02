package main

import (
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

	BusinessApiRegistered(router)
	CommonApiRegistered(router)
	AccountApiRegistered(router)
	router.Run("127.0.0.1:8095")
}

// 注册业务api
func BusinessApiRegistered(engine *gin.Engine) {
	//open := engine.Group("/business/")
	auth := engine.Group("/business", jwt_auth.JWTAuth())

	{
		auth.GET("/order/list", business.OrderList)
		auth.GET("/query_receiver", business.QueryReceiver)
		auth.GET("/query_sender", business.QuerySender)
	}
	{
		auth.POST("/order/add", business.AddOrder)
	}

}

// 注册通用api
func CommonApiRegistered(engine *gin.Engine) {
	open := engine.Group("")

	{
		open.POST("/send_sms/code", common.SMSSend)
	}

	auth := engine.Group("", jwt_auth.JWTAuth())
	{
		auth.GET("/query_area", common.QueryArea)
	}

}

// 注册账号api
func AccountApiRegistered(engine *gin.Engine) {
	open := engine.Group("/account")

	{
		open.POST("/registered", account.Registered)
		open.POST("/edit_password", account.EditPassword)
		open.POST("/sign_in", account.SignIn)
	}

	auth := engine.Group("/account", jwt_auth.JWTAuth())
	{
		auth.GET("/sign_out", account.SignOut)
	}
}
