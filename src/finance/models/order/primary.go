package order

import (
	"finance/models"
	"finance/models/receiver"
	"finance/models/sender"
	"github.com/jinzhu/gorm"
)

type FinanceOrder struct {
	gorm.Model
	ReceiverName      string `json:"receiver_name" gorm:"COMMENT:'收货人'"`
	ReceiverPhone     string `json:"receiver_phone" gorm:"COMMENT:'收货人手机'"`
	ReceiverAddress   string `json:"receiver_address" gorm:"COMMENT:'收货人地址'"`
	ReceiverTel       string `json:"receiver_tel" gorm:"COMMENT:'收货人电话'"`
	SenderCompanyName string `json:"sender_company_name" gorm:"COMMENT:'发货单位'"`
	SenderPhone       string `json:"sender_phone" gorm:"COMMENT:'发货单位手机'"`
	SenderRemark      string `json:"sender_remark" gorm:"COMMENT:'发货单位备注'"`

	FinanceID uint `json:"finance_id" gorm:"COMMENT:'财务id'"`

	Receiver   receiver.FinanceReceiver `gorm:"ForeignKey:ReceiverId"`
	ReceiverId uint                     `json:"receiver_id" gorm:"COMMENT:'收货人id'"`

	Sender   sender.FinanceSender `gorm:"ForeignKey:SenderId"`
	SenderId uint                 `json:"sender_id" gorm:"COMMENT:'发货人id'"`

	//Province     area.Area `gorm:"ForeignKey:ProvinceId"`
	ProvinceId   uint   `json:"province_id"`
	ProvinceName string `json:"province_name"`

	//City     area.Area `gorm:"ForeignKey:CityId"`
	CityId   uint   `json:"city_id"`
	CityName string `json:"city_name"`

	//Area     area.Area `gorm:"ForeignKey:AreaId"`
	AreaId   uint   `json:"area_id"`
	AreaName string `json:"area_name"`

	Deliver          int                  `json:"deliver" gorm:"COMMENT:'提货方式: 1:自提 2:送到'"`
	PaymentMethod    int                  `json:"payment_method" gorm:"COMMENT:'付款方式: 1:现付 2:到付 3:汇款'"`
	ExpectedAmount   float64              `json:"_" gorm:"COMMENT:'预期收费'"`
	ActualAmount     float64              `json:"-" gorm:"COMMENT:'实际付款'"`
	Details          []FinanceOrderDetail `gorm:"ForeignKey:OrderId"`
	AllocationStatus int                  `gorm:"DEFAULT:0;COMMENT:'分配状态:0:未分配 1:已分配'"`
}

// 批量插入订单详情
func (order *FinanceOrder) AddDetails(products []*Product) {
	for _, product := range products {
		detail := FinanceOrderDetail{
			OrderId:  order.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
			Unit:     product.Unit,
			Price:    product.Price}
		models.DB.Save(&detail)
	}
}

// 调整订单分配状态
func (order *FinanceOrder) EditAllocationStatus(status int) {
	order.AllocationStatus = status
	models.DB.Save(&order)
}

func (order *FinanceOrder) DeleteAllDetail() {
	models.DB.Delete(FinanceOrderDetail{}, "order_id=?", order.ID)
}

// 序列化
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
			"payment_method":      order.PaymentMethod,
			"allocation_status":   order.AllocationStatus},
		"product_information": order.Details,
		"create_time":         order.CreatedAt,
		"update_time":         order.UpdatedAt}
}

// 加载订单详情
func (order *FinanceOrder) QueryDetails() {
	models.DB.Model(&order).Association("Details").Find(&order.Details)
}
