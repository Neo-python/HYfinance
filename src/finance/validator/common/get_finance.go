package common

import (
	"finance/models"
	"finance/models/finance"
	"finance/plugins/jwt_auth"
	"github.com/gin-gonic/gin"
)

// 获取用户模型
func GetFinance(context *gin.Context) (*finance.Finance, error) {
	var finance finance.Finance

	claims, err := jwt_auth.GetClaims(context)

	if err != nil {
		return &finance, err
	}

	// 操作数据库
	if err := models.DB.First(&finance, claims.ID).RecordNotFound(); err {
		return &finance, nil
	} else {
		return &finance, models.DB.Error
	}

}
