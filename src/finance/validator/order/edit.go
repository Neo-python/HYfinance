package order

import (
	"errors"
)

type OrderEditForm struct {
	OrderFormBase
	OrderIdBase
	ExtraData OrderFormExtraData
}

func (form *OrderEditForm) Valid() error {
	form.Order()
	if form.FormOrder.ID == 0 {
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

// 订单金额修改
type OrderAmountEditForm struct {
	OrderInfo
	ExpectedAmount float64 `json:"expected_amount"`
	ActualAmount   float64 `json:"actual_amount"`
}
