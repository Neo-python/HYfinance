package core_sms

import (
	"errors"
	"net/http"
	"net/url"
)

// 注册短信模板编号
var SMSTemplateId map[string]string = map[string]string{"registered": "477438", "edit_password": "482880"}

type Phone struct {
	Phone string `json:"phone"`
}

type Genre struct {
	Genre string `json:"genre"`
}

type SMS struct {
	*Phone
	*Genre
}

type EditCode struct {
	Code string `json:"code"`
}

// 发送短信
func (sms *SMS) Send(code string) (string, error) {
	core_url := "http://127.0.0.1:8090/send_sms/code/"
	template_id := SMSTemplateId[sms.Genre.Genre]
	data := url.Values{"phone": {sms.Phone.Phone}, "code": {code}, "template_id": {template_id}}
	resp, err := http.PostForm(core_url, data)

	if err != nil {
		return "短信发送失败", err
	} else if resp.StatusCode != 200 {
		return "短信接口请求失败", errors.New("短信接口请求失败")
	} else {
		return "短信发送成功", nil
	}
}
