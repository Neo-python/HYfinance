package jwt_auth

import (
	plugins "finance/plugins/common"
	"github.com/gin-gonic/gin"
)

// 等级权限验证中间件
func LevelAuth(level int) gin.HandlerFunc {
	return func(context *gin.Context) {
		checkToken(context)
		if context.IsAborted() == true {
			return
		}

		claims, err := GetClaims(context)
		if err != nil {
			plugins.ApiExport(context).Error(4005, "用户未登录")
		}
		if claims.Level < level {
			plugins.ApiExport(context).Error(5011, "权限不足")
		}

	}
}
