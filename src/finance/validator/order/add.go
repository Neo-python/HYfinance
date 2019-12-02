package order

import "errors"

type AddOrderForm struct {
	OrderFormBase
	ExtraData OrderFormExtraData
}

// 自定义验证逻辑
func (form *AddOrderForm) Valid() error {

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
