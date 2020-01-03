// 核心接口调用
package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CoreGetFactoryToken(factory_uuid string) (string, error) {
	url := fmt.Sprintf("http://127.0.0.1:8090/factory/get_token/?factory_uuid=%s", factory_uuid)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
