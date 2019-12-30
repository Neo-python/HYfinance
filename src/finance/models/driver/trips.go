package driver

import (
	"finance/models"
	"finance/models/order"
	"github.com/jinzhu/gorm"
	"time"
)

type FinanceDriverTrips struct {
	gorm.Model
	ProvinceId   uint                        `json:"province_id"`
	ProvinceName string                      `json:"province_name" gorm:"COMMENT:'目的地'"`
	Date         time.Time                   `json:"date" gorm:"COMMENT:'出发日期'"`
	Driver       FinanceDriver               `gorm:"ForeignKey:DriverId"`
	DriverId     uint                        `json:"driver_id"`
	Remark       string                      `json:"remark" form:"remark" gorm:"COMMENT:'车次备注'"`
	Details      []FinanceDriverTripsDetails `gorm:"ForeignKey:TripsId"`
}

func (trips *FinanceDriverTrips) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":            trips.ID,
		"province_id":   trips.ProvinceId,
		"province_name": trips.ProvinceName,
		"date":          trips.Date,
		"driver_id":     trips.DriverId,
		"remark":        trips.Remark,
	}
}
func (trips *FinanceDriverTrips) GetDetails() {
	models.DB.Model(&trips).Association("Details").Find(&trips.Details)
}

// 删除车次,处理订单分配状态
func (trips *FinanceDriverTrips) DeleteSelf() {
	trips.GetDetails()
	for _, item := range trips.Details {
		item.DeleteSelf()
	}
	models.DB.Unscoped().Delete(&trips)
}

type FinanceDriverTripsDetails struct {
	gorm.Model
	TripsId        uint               `json:"trips_id" gorm:"unique_index:trips_order_only"`
	Trips          FinanceDriverTrips `gorm:"ForeignKey:TripsId"`
	OrderId        uint               `json:"order_id" gorm:"unique_index:trips_order_only"`
	Order          order.FinanceOrder `gorm:"ForeignKey:OrderId"`
	ExpectedAmount float64            `json:"_" gorm:"COMMENT:'预期收费'"`
	ActualAmount   float64            `json:"-" gorm:"COMMENT:'实际付款'"`
}

// 获取对应订单
func (details *FinanceDriverTripsDetails) GetOrder() {
	models.DB.Model(&details).Association("Order").Find(&details.Order)
}

// 自身序列化
func (details *FinanceDriverTripsDetails) SelfToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":              details.ID,
		"trips_id":        details.TripsId,
		"order_id":        details.OrderId,
		"expected_amount": details.ExpectedAmount,
		"actual_amount":   details.ActualAmount}
}

// 车次分配订单详情序列化
func (details *FinanceDriverTripsDetails) ToJson() map[string]interface{} {
	details.GetOrder()
	return map[string]interface{}{
		"record": details.SelfToJson(),
		"order":  details.Order.ToJson()["base_info"]}
}

// 释放订单
func (details *FinanceDriverTripsDetails) FreedOrder() {
	details.GetOrder()
	details.Order.EditAllocationStatus(0)
}

// 删除自身
func (details *FinanceDriverTripsDetails) DeleteSelf() {
	details.FreedOrder()
	models.DB.Unscoped().Delete(&details)
}
