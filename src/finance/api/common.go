// 通用接口
package api

import (
	plugins "finance/plugins/common"
	"finance/plugins/core_sms"
	"finance/plugins/redis"
	validator "finance/validator/common"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 发送短信
func SMSSend(c *gin.Context) {
	// 处理表单
	var form validator.SMSSend
	c.Bind(&form)
	_, err := govalidator.ValidateStruct(&form)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{"message": err.Error()})
	}

	// 发送短信前的准备工作
	code := plugins.GenerateVerifyCode(4)
	redis.Set(fmt.Sprint("Registered_%", form.Phone), code, 1)

	// 发送短信
	sms := core_sms.SMS{&core_sms.PhoneBase{form.Phone}, &core_sms.GenreBase{form.Genre}}
	resp, err := sms.Send(code)
	export := plugins.Export{}
	if err != nil {
		// 修改接口错误码与提示信息
		export.ErrorCode = http.StatusNotAcceptable
		export.Message = resp
	}
	time.Sleep(time.Second + 1)
	fmt.Println(redis.Get(fmt.Sprint("Registered_%", form.Phone)))
	// 接口返回
	c.JSON(http.StatusOK, export.ApiExport())
}
