package order

import (
	"finance/models"
	"finance/models/receiver"
	"finance/models/sender"
	"github.com/jinzhu/gorm"
)

type FinanceOrder struct {
	gorm.Model
	*receiver.ReceiverModel
	*sender.SenderModel

	FinanceID uint `json:"finance_id"`

	Receiver   *receiver.FinanceReceiver `gorm:"ForeignKey:ReceiverId"`
	ReceiverId uint                      `json:"receiver_id"`

	Sender   *sender.FinanceSender `gorm:"ForeignKey:SenderId"`
	SenderId uint                  `json:"sender_id"`

	//Province     area.Area `gorm:"ForeignKey:ProvinceId"`
	ProvinceId   uint   `json:"province_id"`
	ProvinceName string `json:"province_name"`

	//City     area.Area `gorm:"ForeignKey:CityId"`
	CityId   uint   `json:"city_id"`
	CityName string `json:"city_name"`

	//Area     area.Area `gorm:"ForeignKey:AreaId"`
	AreaId   uint   `json:"area_id"`
	AreaName string `json:"area_name"`

	Deliver       int     `json:"deliver"`
	PaymentMethod int     `json:"payment_method"`
	TotalPrice    float64 `json:"total_price"`
}

// 批量插入订单详情
func (order *FinanceOrder) AddDetails(products []*Product) {
	for _, product := range products {
		detail := FinanceOrderDetail{
			OrderId:  order.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
			Price:    product.Price}
		models.DB.Save(&detail)
	}
}
