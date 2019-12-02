package order

import (
	"finance/models"
	models_order "finance/models/order"
	"github.com/jinzhu/gorm"
)

type OrderInfo struct {
	OrderIdBase
}

func (form *OrderInfo) Query() *gorm.DB {
	query := models.DB.Model(models_order.FinanceOrder{})

	query = query.Where("id=?", form.OrderId)

	return query
}
