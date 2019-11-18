package api

import (
	"finance/validator/account"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册账号
func Registered(c *gin.Context) {
	var registered account.Registered
	c.Bind(&registered)
	result, err := govalidator.ValidateStruct(&registered)
	fmt.Println(result, err)
	fmt.Println(registered)
	fmt.Println("ok")
	c.JSON(http.StatusOK, gin.H{"status": 1, "name": registered.Name})
}
