package receiver

import (
	"time"
)

type FinanceReceiver struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `json:"name" gorm:"COMMENT:'名称'"`
	Phone     string     `json:"phone" gorm:"COMMENT:'手机号'"`
	Address   string     `json:"address" gorm:"COMMENT:'收货地址'"`
	Tel       string     `json:"tel" gorm:"COMMENT:'电话号'"`
}

type FinanceReceiverJson struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}
