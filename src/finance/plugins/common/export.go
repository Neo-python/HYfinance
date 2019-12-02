package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
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
	context *gin.Context
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
func (export *Export) ApiExport() {
	if export.Data == nil {
		export.Data = map[string]interface{}{}
	}
	export.context.JSON(http.StatusOK, export)
	export.context.Abort()
}

// 列表页返回
func (export *Export) ListPageExport(items interface{}, page int, total int) {
	if export.Data == nil {
		export.Data = map[string]interface{}{}
	}
	export.SetData("items", items)
	export.SetData("page", page)
	export.SetData("total", total)
	export.context.JSON(http.StatusOK, export)
	export.context.Abort()

}

// 接口表单异常返回,错误码:1001
// 表单验证错误专用,其他类型错误已经无法兼容此方法逻辑
func (export *Export) FormError(err error) {
	error_export := ErrorExport{}
	error_export.ErrorFields = map[string]interface{}{} // map类型字段 必须初始化

	error_export.ErrorCode = 1001

	var error_message string
	var filed_name string
	//var error_fields string
	for _, err := range err.(validator.ValidationErrors) {
		err_field := err.Field()
		split_result := strings.Split(err_field, "~")
		filed_name = split_result[0]
		for _, item := range strings.Split(split_result[1], ";") {
			item_split := strings.Split(item, ":")
			key := item_split[0]
			error_message = item_split[1]
			if err.ActualTag() == key {
				//error_message = value
				//error_fields =
				error_export.ErrorFields[err.StructField()] = key
				break
			}

		}

	}

	error_export.Message = filed_name + error_message
	export.context.JSON(http.StatusOK, error_export)
	export.context.Abort()
}

func (export *Export) Error(error_code int, message string) {
	error_export := ErrorExport{}
	error_export.ErrorFields = map[string]interface{}{} // map类型字段 必须初始化

	error_export.ErrorCode = error_code
	error_export.Message = message

	export.context.JSON(http.StatusOK, error_export)
	export.context.Abort()
}

func ApiExport(context *gin.Context) *Export {
	export := new(Export)
	export.context = context
	return export
}
