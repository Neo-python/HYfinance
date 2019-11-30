package business

import (
	"finance/models"
	models_order "finance/models/order"
	plugins "finance/plugins/common"
	"finance/validator"
	"finance/validator/common"
	"finance/validator/order"
	"github.com/gin-gonic/gin"
)

// 新增订单
func AddOrder(context *gin.Context) {
	var form order.AddOrderForm
	context.ShouldBind(&form)
	// 基础验证
	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	// 自定义逻辑验证
	extra_data, err := form.Valid()
	if err != nil {
		plugins.ApiExport(context).Error(5400, err.Error())
		return
	}

	// 获取新增订单财务人信息
	finance, err := common.GetFinance(context)
	if err != nil {
		plugins.ApiExport(context).Error(4005, "用户未登录,请在登录后尝试.")
		return
	}

	order := models_order.FinanceOrder{
		Receiver:          extra_data.Receiver,
		ReceiverName:      extra_data.Receiver.Name,
		ReceiverPhone:     extra_data.Receiver.Phone,
		ReceiverAddress:   extra_data.Receiver.Address,
		ReceiverTel:       extra_data.Receiver.Tel,
		Sender:            extra_data.Sender,
		SenderCompanyName: extra_data.Sender.CompanyName,
		SenderPhone:       extra_data.Sender.Phone,
		SenderRemark:      extra_data.Sender.Remark,
		FinanceID:         finance.ID,
		ProvinceId:        extra_data.Province.ID,
		ProvinceName:      extra_data.Province.Name,
		CityId:            extra_data.City.ID,
		CityName:          extra_data.City.Name,
		AreaId:            extra_data.Area.ID,
		AreaName:          extra_data.Area.Name,
		TotalPrice:        extra_data.Price}

	models.DB.Save(&order)

	// 批量增加订单货物信息
	go order.AddDetails(form.Products)
	plugins.ApiExport(context).ApiExport()
	return
}
