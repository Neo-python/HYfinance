package sender

import "github.com/jinzhu/gorm"

type FinanceSender struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
	Remark      string `json:"remark"`
}

type FinanceSenderJson struct {
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
	Remark      string `json:"remark"`
}
