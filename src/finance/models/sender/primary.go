package sender

import "github.com/jinzhu/gorm"

type FinanceSender struct {
	gorm.Model
	CompanyName string `json:"company_name" gorm:"COMMENT:'单位名'"`
	Phone       string `json:"phone" gorm:"COMMENT:'手机号'"`
	Remark      string `json:"remark" gorm:"COMMENT:'发货单位备注'"`
}

type FinanceSenderJson struct {
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
	Remark      string `json:"remark"`
}
