// 通用接口
package api

import (
	"finance/validator/common"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func SMSsend(c *gin.Context) {
	var form common.SMSSend
	c.Bind(&form)
	result, err := govalidator.ValidateStruct(&form)
	fmt.Println(result, err)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{"message": err.Error()})
	}
	code := "6666"
	data := url.Values{"phone": {form.Phone}, "code": {code}, "template_id": {"446286"}}
	url := "http://127.0.0.1:8090/send_sms/code/"
	//resp, err := http.Get(url)
	resp, err := http.PostForm(url, data)
	fmt.Println(resp, err)
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
