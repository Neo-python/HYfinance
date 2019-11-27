package receiver

import "github.com/jinzhu/gorm"

type ReceiverModel struct {
	ReceiverName    string `json:"receiver_name"`
	ReceiverPhone   string `json:"receiver_phone"`
	ReceiverAddress string `json:"receiver_address"`
	ReceiverTel     string `json:"receiver_tel"`
}

type FinanceReceiver struct {
	gorm.Model
	ReceiverModel
}
