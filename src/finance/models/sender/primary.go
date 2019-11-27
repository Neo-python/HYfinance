package sender

import "github.com/jinzhu/gorm"

type SenderModel struct {
	SenderCompanyName string `json:"sender_company_name"`
	SenderPhone       string `json:"sender_phone"`
	SenderRemark      string `json:"sender_remark"`
}

type FinanceSender struct {
	gorm.Model
	SenderModel
}
