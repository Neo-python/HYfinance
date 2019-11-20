package main

import (
	"finance/api"
	"finance/plugins"
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
	router.GET("test", api.Registered)
	router.GET("core_sms", api.SMSSend)
	router.Run()
}
