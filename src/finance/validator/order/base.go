package order

import (
	"errors"
	"finance/models"
	"finance/models/area"
	"finance/models/order"
	models_order "finance/models/order"
	"finance/models/receiver"
	"finance/models/sender"
	"github.com/jinzhu/gorm"
)

type OrderIdBase struct {
	OrderId   uint `json:"order_id" form:"order_id" validate:"required" error_message:"订单编号~required:未选择,请选择后重试."`
	FormOrder models_order.FinanceOrder
}

func (form *OrderIdBase) Query() *gorm.DB {
	query := models.DB.Model(models_order.FinanceOrder{})

	query = query.Where("id=?", form.OrderId)

	return query
}

// 复用订单数据
func (form *OrderIdBase) Order() *models_order.FinanceOrder {
	if form.FormOrder.ID == 0 {
		form.FormOrder = *form.GetOrder()
	}

	return &form.FormOrder
}

func (form *OrderIdBase) GetOrder() *models_order.FinanceOrder {
	var order models_order.FinanceOrder
	models.DB.First(&order, form.OrderId)
	return &order
}

type OrderFormBase struct {
	// 收货人相关
	ReceiverName    string `json:"receiver_name" validate:"required,max=20" error_message:"收货人名~required:此字段必须填写;max:最大长度为20"`
	ReceiverPhone   string `json:"receiver_phone" validate:"required,max=11" error_message:"收货人手机号~required:此字段必须填写;max:最大长度为11"`
	ReceiverAddress string `json:"receiver_address" validate:"max=255" error_message:"收货人地址~max:最大长度为255"`
	ReceiverTel     string `json:"receiver_tel" validate:"max=13" error_message:"收货人电话~max:最大长度为13"`

	// 发货人相关
	SenderCompanyName string `json:"sender_company_name" validate:"max=10" error_message:"发货单位名称~max:最大长度为10"`
	SenderPhone       string `json:"sender_phone" validate:"required,max=11" error_message:"发货单位手机号~required:此字段必须填写;max:最大长度为11"`
	SenderRemark      string `json:"sender_remark" validate:"max=255" error_message:"发货人备注~max:最大长度为255"`

	// 收货人地址相关
	ProvinceId uint `json:"province_id" validate:"required" error_message:"省级信息~required:此字段必须填写"`
	CityId     uint `json:"city_id" validate:"required" error_message:"市级信息~required:此字段必须填写"`
	AreaId     uint `json:"area_id" validate:"required" error_message:"区级信息~required:此字段必须填写"`

	// 订单辅助信息
	Deliver            int                      `json:"deliver" validate:"required" error_message:"交付方式~required:此字段必须填写"`
	PaymentMethod      int                      `json:"payment_method" validate:"required" error_message:"付款方式~required:此字段必须填写"`
	ProductInformation []map[string]interface{} `json:"product_information" validate:"required" error_message:"货物信息~required:需要填写"`
	Products           []*order.Product
}

type QueryForm struct {
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
}

// 表单提交后,要交给视图处理的额外数据
type OrderFormExtraData struct {
	Receiver *receiver.FinanceReceiver
	Sender   *sender.FinanceSender
	Province *area.Area
	City     *area.Area
	Area     *area.Area
}

// 完善地区信息
func (form *OrderFormBase) PerfectArea(extra_data *OrderFormExtraData) error {
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

//
// 获取收货人
// 无匹配项时创建新收货人
// 更新地址与电话信息
func (form *OrderFormBase) GetReceiver() *receiver.FinanceReceiver {
	var receiver receiver.FinanceReceiver
	models.DB.First(&receiver, "phone=? AND name=?", form.ReceiverPhone, form.ReceiverName)

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
func (form *OrderFormBase) GetSender() *sender.FinanceSender {
	var sender sender.FinanceSender
	models.DB.First(&sender, "phone=? AND company_name=?", form.SenderPhone, form.SenderCompanyName)

	// 无匹配项时创建新发货人
	if sender.ID == 0 {
		sender.Phone = form.SenderPhone
		sender.CompanyName = form.SenderCompanyName

	}
	sender.Remark = form.SenderRemark
	models.DB.Save(&sender)
	return &sender
}

// 检查货物信息
func (form *OrderFormBase) CheckProduct() bool {
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

		unit, ok := item["unit"].(float64)
		if ok == false {
			return false
		}

		measure, ok := item["measure"].(float64)
		if ok == false {
			return false
		}
		form.Products = append(form.Products, &order.Product{Name: name, Quantity: int(quantity), Price: int(price), Unit: int(unit), Measure: int(measure)})
	}
	return true
}
