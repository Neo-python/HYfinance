package receiver

import "github.com/jinzhu/gorm"

type FinanceReceiver struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

type FinanceReceiverJson struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}
