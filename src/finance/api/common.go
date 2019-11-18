// 通用接口
package api

import (
	"finance/validator/common"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SMSsend(c *gin.Context) {
	var form common.SMSSend
	c.Bind(&form)
	result, err := govalidator.ValidateStruct(&form)
	fmt.Println(result, err)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
