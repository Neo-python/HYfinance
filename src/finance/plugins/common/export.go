package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ErrorCodeStruct struct {
	ErrorCode int `json:"error_code"`
}
type MessageStruct struct {
	Message string `json:"message"`
}

type ErrorFieldsStruct struct {
	ErrorFields map[string]interface{} `json:"error_fields"`
}

type Export struct {
	ErrorCodeStruct
	Data map[string]interface{} `json:"data"`
	MessageStruct
}

type ErrorExport struct {
	ErrorCodeStruct
	MessageStruct
	ErrorFieldsStruct
}

// 装填数据
func (export *Export) SetData(key string, value interface{}) {
	if export.Data == nil {
		export.Data = map[string]interface{}{}
	}

	export.Data[key] = value
}

// 接口正常返回
func (export *Export) ApiExport(context *gin.Context) {
	if export.Data == nil {
		export.Data = map[string]interface{}{}
	}
	context.JSON(http.StatusOK, export)
	context.Abort()
}

// 接口表单异常返回
func (export *Export) Error(context *gin.Context, err error) {
	error_export := ErrorExport{}
	error_export.ErrorFields = map[string]interface{}{} // map类型字段 必须初始化

	error_export.ErrorCode = 500
	error_export.Message = "表单未能通过接口验证"

	// 收集错误信息
	fileds := strings.Split(err.Error(), ";")
	for _, item := range fileds {
		errorInfo := strings.Split(item, ":")
		if len(errorInfo) > 1 {
			error_export.ErrorFields[errorInfo[0]] = errorInfo[1]
		} else {
			error_export.ErrorFields[errorInfo[0]] = errorInfo[0]
		}

	}
	context.JSON(http.StatusOK, error_export)
	context.Abort()
}

var ApiExport Export

//func main() {
//	export := Export{}
//	var err error
//	export.Error(err)
//	fmt.Println("ok")
//}
