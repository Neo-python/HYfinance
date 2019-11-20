package core_sms

import (
	"net/http"
	"net/url"
)

// 注册短信模板编号
var SMSTemplateId map[string]string = map[string]string{"registered": "477438"}

type PhoneBase struct {
	Phone string `json:"phone"`
}

type GenreBase struct {
	Genre string `json:"genre"`
}

type SMS struct {
	*PhoneBase
	*GenreBase
}

type EditCode struct {
	Code string `json:"code"`
}

// 发送短信
func (sms *SMS) Send(code string) (string, error) {
	core_url := "http://127.0.0.1:8090/send_sms/code/"
	template_id := SMSTemplateId[sms.Genre]
	data := url.Values{"phone": {sms.Phone}, "code": {code}, "template_id": {template_id}}
	_, err := http.PostForm(core_url, data)
	if err != nil {
		return "短信发送失败", err
	} else {
		return "短信发送成功", nil
	}
}
