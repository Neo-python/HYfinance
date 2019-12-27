package order

import (
	"finance/models"
	"finance/models/receiver"
	"finance/models/sender"
	"github.com/jinzhu/gorm"
)

type FinanceOrder struct {
	gorm.Model
	ReceiverName      string `json:"receiver_name"`
	ReceiverPhone     string `json:"receiver_phone"`
	ReceiverAddress   string `json:"receiver_address"`
	ReceiverTel       string `json:"receiver_tel"`
	SenderCompanyName string `json:"sender_company_name"`
	SenderPhone       string `json:"sender_phone"`
	SenderRemark      string `json:"sender_remark"`

	FinanceID uint `json:"finance_id"`

	Receiver   receiver.FinanceReceiver `gorm:"ForeignKey:ID"`
	ReceiverId uint                     `json:"receiver_id"`

	Sender   sender.FinanceSender `gorm:"ForeignKey:SenderId"`
	SenderId uint                 `json:"sender_id"`

	//Province     area.Area `gorm:"ForeignKey:ProvinceId"`
	ProvinceId   uint   `json:"province_id"`
	ProvinceName string `json:"province_name"`

	//City     area.Area `gorm:"ForeignKey:CityId"`
	CityId   uint   `json:"city_id"`
	CityName string `json:"city_name"`

	//Area     area.Area `gorm:"ForeignKey:AreaId"`
	AreaId   uint   `json:"area_id"`
	AreaName string `json:"area_name"`

	Deliver        int                  `json:"deliver"`
	PaymentMethod  int                  `json:"payment_method"`
	ExpectedAmount float64              `json:"_"`
	ActualAmount   float64              `json:"-"`
	Details        []FinanceOrderDetail `gorm:"ForeignKey:OrderId"`
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

func (order *FinanceOrder) DeleteAllDetail() {
	models.DB.Delete(FinanceOrderDetail{}, "order_id=?", order.ID)
}

func (order *FinanceOrder) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"base_info": map[string]interface{}{
			"receiver_name":       order.ReceiverName,
			"receiver_phone":      order.ReceiverPhone,
			"receiver_address":    order.ReceiverAddress,
			"receiver_tel":        order.ReceiverTel,
			"sender_company_name": order.SenderCompanyName,
			"sender_phone":        order.SenderPhone,
			"sender_remark":       order.SenderRemark,
			"province_id":         order.ProvinceId,
			"province_name":       order.ProvinceName,
			"city_id":             order.CityId,
			"city_name":           order.CityName,
			"area_id":             order.AreaId,
			"area_name":           order.AreaName,
			"deliver":             order.Deliver,
			"payment_method":      order.PaymentMethod},
		"product_information": order.Details,
		"create_time":         order.CreatedAt,
		"update_time":         order.UpdatedAt}
}
