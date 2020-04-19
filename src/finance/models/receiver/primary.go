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
	AutoFill  int        `json:"auto_fill" gorm:"COMMENT:'自动填充状态';DEFAULT:1"`
	IdCard    string     `json:"id_card" gorm:"COMMENT:'身份证'"`
}

type FinanceReceiverJson struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

type FinanceReceiverProduct struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name       string `json:"name" gorm:"COMMENT:'名称';index:name_index"`
	Unit       int    `json:"unit" gorm:"COMMENT:'单位'"`
	Price      int    `json:"price" gorm:"COMMENT:'单价'"`
	ReceiverId uint   `json:"receiver_id" gorm:"COMMENT:'收货人编号'"`
}

func (product *FinanceReceiverProduct) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"unit":        product.Unit,
		"price":       product.Price,
		"receiver_id": product.ReceiverId}
}

func (receiver *FinanceReceiver) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":        receiver.ID,
		"name":      receiver.Name,
		"address":   receiver.Address,
		"auto_fill": receiver.AutoFill,
		"id_card":   receiver.IdCard,
		"tel":       receiver.Tel,
		"phone":     receiver.Phone}

}
