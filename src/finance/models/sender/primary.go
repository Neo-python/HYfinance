package sender

import "github.com/jinzhu/gorm"

type FinanceSender struct {
	gorm.Model
	CompanyName string `json:"company_name" gorm:"COMMENT:'单位名'"`
	Phone       string `json:"phone" gorm:"COMMENT:'手机号'"`
	Remark      string `json:"remark" gorm:"COMMENT:'发货单位备注'"`
	AutoFill    int    `json:"auto_fill" gorm:"COMMENT:'自动填充';DEFAULT:1"`
	IdCard      string `json:"id_card" gorm:"COMMENT:'身份证'"`
}

type FinanceSenderJson struct {
	Id          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
	Remark      string `json:"remark"`
}

func (receiver *FinanceSender) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":           receiver.ID,
		"company_name": receiver.CompanyName,
		"remark":       receiver.Remark,
		"auto_fill":    receiver.AutoFill,
		"id_card":      receiver.IdCard,
		"phone":        receiver.Phone}

}
