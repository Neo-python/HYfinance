package order

import (
	"errors"
	"finance/models"
	"finance/models/order"
)

type OrderEditForm struct {
	OrderFormBase
	OrderIdBase
	Order     order.FinanceOrder
	ExtraData OrderFormExtraData
}

func (form *OrderEditForm) Valid() error {

	if err := models.DB.Find(&form.Order, form.OrderId).Error; err != nil {
		return errors.New("订单编号未找到")
	}

	if form.CheckProduct() == false {
		return errors.New("货物字段非法,请联系管理员.")
	}

	if err := form.PerfectArea(&form.ExtraData); err != nil {
		return err
	}

	// 验证通过,完善数据
	form.ExtraData.Receiver = form.GetReceiver()
	form.ExtraData.Sender = form.GetSender()
	return nil

}
