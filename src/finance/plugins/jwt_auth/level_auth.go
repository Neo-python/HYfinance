package jwt_auth

import (
	"finance/plugins/common"
	"github.com/gin-gonic/gin"
)

// 等级权限验证中间件
func LevelAuth(level int) gin.HandlerFunc {
	return func(context *gin.Context) {
		checkToken(context)
		if context.IsAborted() == true {
			return
		}
		result, _ := context.Get("claims")
		claims, _ := result.(*CustomClaims)
		if claims.Level < level {
			common.ApiExport(context).Error(5011, "权限不足")
		}

	}
}
