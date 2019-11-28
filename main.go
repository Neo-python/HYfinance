package main

import (
	"finance/api"
	_ "finance/models"
	"finance/plugins"
	"finance/plugins/jwt_auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{Name: "a", Age: 1}
	fmt.Print(user.Name)
	router := gin.Default()
	fmt.Println(plugins.Config.SMSAppId)

	// 无需权限验证的接口
	open := router.Group("")
	{
		open.POST("registered", api.Registered)
		open.POST("edit_password", api.EditPassword)
		open.POST("/send_sms/code/", api.SMSSend)
		open.POST("sign_in", api.SignIn)
	}

	// 需要权限验证的接口
	auth := router.Group("", jwt_auth.JWTAuth())
	{
		auth.GET("test", api.Test)
		auth.GET("sign_out", api.SignOut)
		auth.GET("query_area", api.QueryArea)
	}

	router.Run()
}
