package order

import (
	"finance/models/receiver"
	"finance/models/sender"
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	receiver.Receiver
	sender.Sender
	ProvinceId    uint    `json:"province_id"`
	ProvinceName  string  `json:"province_name"`
	CityId        uint    `json:"city_id"`
	CityName      string  `json:"city_name"`
	AreaId        uint    `json:"area_id"`
	AreaName      string  `json:"area_name"`
	Deliver       int     `json:"deliver"`
	PaymentMethod int     `json:"payment_method"`
	TotalPrice    float64 `json:"total_price"`
	FinanceID     uint    `json:"finance_id"`
}
