package order

import (
	"errors"
	"finance/models"
	"finance/models/area"
	"finance/models/order"
	"finance/models/receiver"
	"finance/models/sender"
	"fmt"
)

type AddOrderForm struct {
	// 收货人相关
	ReceiverName    string `json:"receiver_name" validate:"required" error_message:"收货人名~required:此字段必须填写"`
	ReceiverPhone   string `json:"receiver_phone" validate:"required,max=11" error_message:"收货人手机号~required:此字段必须填写;max:最大长度为11"`
	ReceiverAddress string `json:"receiver_address"`
	ReceiverTel     string `json:"receiver_tel"`

	// 发货人相关
	SenderCompanyName string `json:"sender_company_name" validate:"max=10" error_message:"发货单位名称~max:最大长度为10"`
	SenderPhone       string `json:"sender_phone" validate:"required,max=11" error_message:"发货单位手机号~required:此字段必须填写;max:最大长度为11"`
	SenderRemark      string `json:"sender_remark"`

	// 收货人地址相关
	ProvinceId uint `json:"province_id" validate:"required" error_message:"省级信息~required:此字段必须填写"`
	CityId     uint `json:"city_id" validate:"required" error_message:"市级信息~required:此字段必须填写"`
	AreaId     uint `json:"area_id" validate:"required" error_message:"区级信息~required:此字段必须填写"`

	// 订单辅助信息
	Deliver            int                      `json:"deliver"`
	PaymentMethod      int                      `json:"payment_method"`
	TotalPrice         float64                  `json:"total_price"`
	ProductInformation []map[string]interface{} `json:"product_information" validate:"required" error_message:"货物信息~required:需要填写"`
	Products           []*order.Product
}

// 表单提交后,要交给视图处理的额外数据
type AddOrderFormData struct {
	Receiver *receiver.FinanceReceiver
	Sender   *sender.FinanceSender
	Province *area.Area
	City     *area.Area
	Area     *area.Area
	Price    float64
}

// 自定义验证逻辑
func (form *AddOrderForm) Valid() (AddOrderFormData, error) {

	extra_data := AddOrderFormData{}

	if form.CheckProduct() == false {
		return extra_data, errors.New("货物字段非法,请联系管理员.")
	}

	if err := form.PerfectArea(&extra_data); err != nil {
		return extra_data, err
	}
	// 验证通过,完善数据
	extra_data.Receiver = form.GetReceiver()
	extra_data.Sender = form.GetSender()
	return extra_data, nil
}

// 验证产品信息
func (form *AddOrderForm) ValidProductInformation() {

}

//
// 获取收货人
// 无匹配项时创建新收货人
// 更新地址与电话信息
func (form *AddOrderForm) GetReceiver() *receiver.FinanceReceiver {
	var receiver receiver.FinanceReceiver
	models.DB.First(&receiver, "receiver_phone=? AND receiver_name=?", form.ReceiverPhone, form.ReceiverName)

	// 无匹配项时创建新收货人
	if receiver.ID == 0 {
		receiver.Phone = form.ReceiverPhone
		receiver.Name = form.ReceiverName

	}
	//更新地址与电话信息
	receiver.Address = form.ReceiverAddress
	receiver.Tel = form.ReceiverTel
	models.DB.Save(&receiver)
	return &receiver
}

// 获取发货人
func (form *AddOrderForm) GetSender() *sender.FinanceSender {
	var sender sender.FinanceSender
	models.DB.First(&sender, "sender_phone=? AND sender_company_name=?", form.SenderPhone, form.SenderCompanyName)

	// 无匹配项时创建新发货人
	if sender.ID == 0 {
		sender.Phone = form.SenderPhone
		sender.CompanyName = form.SenderCompanyName

	}
	models.DB.Save(&sender)
	return &sender
}

// 完善地区信息
func (form *AddOrderForm) PerfectArea(extra_data *AddOrderFormData) error {
	var province area.Area
	var city area.Area
	var area area.Area

	models.DB.First(&province, form.ProvinceId)
	models.DB.First(&city, form.CityId)
	models.DB.First(&area, form.AreaId)

	if (province.ID == 0 || city.ID == 0 || area.ID == 0) || (area.SuperiorId != city.ID || city.SuperiorId != province.ID) {
		return errors.New("地区编号错误!")
	} else {
		extra_data.Province = &province
		extra_data.City = &city
		extra_data.Area = &area
	}
	return nil
}

// 检查货物信息
func (form *AddOrderForm) CheckProduct() bool {
	fmt.Println(form.ProductInformation, "CheckProduct")
	// item 单条货物数据
	for _, item := range form.ProductInformation {

		name, ok := item["name"].(string)
		if ok == false {
			return false
		}
		quantity, ok := item["quantity"].(float64)
		if ok == false {
			return false
		}
		price, ok := item["price"].(float64)
		if ok == false {
			return false
		}
		form.Products = append(form.Products, &order.Product{Name: name, Quantity: int(quantity), Price: int(price)})
	}
	return true
}
