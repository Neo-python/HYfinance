package order

import (
	"finance/models"
	models_order "finance/models/order"
	"finance/validator"
	"fmt"
	"github.com/jinzhu/gorm"
)

type OrderListForm struct {
	validator.ListPage
	// 收货人地址相关
	ProvinceId       uint   `json:"province_id" form:"province_id"`
	CityId           uint   `json:"city_id" form:"city_id"`
	AreaId           uint   `json:"area_id" form:"area_id"`
	SenderId         uint   `json:"sender_id" form:"sender_id"`
	ReceiverId       uint   `json:"receiver_id" form:"receiver_id"`
	ProductName      string `json:"product_name" form:"product_name"`
	StartDay         string `json:"start_day" form:"start_day"`
	EndDay           string `json:"end_day" form:"end_day"`
	AllocationStatus int    `json:"allocation_status" form:"allocation_status"`
}

func (form *OrderListForm) Query() *gorm.DB {
	query := models.DB.Model(models_order.FinanceOrder{})

	if form.AllocationStatus != 2 {
		query = query.Where("allocation_status=?", form.AllocationStatus)
	}

	if form.ReceiverId != 0 {
		query = query.Where("receiver_id=?", form.ReceiverId)
	}

	if form.SenderId != 0 {
		query = query.Where("sender_id=?", form.SenderId)
	}

	if form.ProvinceId != 0 {
		query = query.Where("province_id=?", form.ProvinceId)
	}
	if form.CityId != 0 {
		query = query.Where("city_id=?", form.CityId)
	}
	if form.AreaId != 0 {
		query = query.Where("area_id=?", form.AreaId)
	}

	if form.ProductName != "" {
		query = query.Joins("left join finance_order_detail on finance_order_detail.order_id=finance_order.id")
		query = query.Where(fmt.Sprintf("finance_order_detail.name like '%%%s%%'", form.ProductName))
	}

	if form.StartDay != "" {
		query = query.Where("created_at > ?", form.StartDay)
	}

	if form.EndDay != "" {
		query = query.Where("created_at < ?", form.EndDay)
	}

	return query
}
