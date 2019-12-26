package common

import (
	"errors"
	"finance/models"
	"finance/models/finance"
	"finance/plugins/jwt_auth"
	"github.com/gin-gonic/gin"
)

// 获取用户模型
func GetFinance(context *gin.Context) (*finance.Finance, error) {
	var finance finance.Finance
	result, status := context.Get("claims")

	if status != true {
		return &finance, errors.New("未能获取用户")
	}

	claims, status := result.(*jwt_auth.CustomClaims)

	if status != true {
		return &finance, errors.New("用户数据异常,请重新登录!")
	}

	// 操作数据库
	if err := models.DB.First(&finance, claims.ID).RecordNotFound(); err {
		return &finance, nil
	} else {
		return &finance, models.DB.Error
	}

}
