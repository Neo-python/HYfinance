package business

import (
	"finance/models"
	models_order "finance/models/order"
	plugins "finance/plugins/common"
	"finance/validator"
	"finance/validator/common"
	forms "finance/validator/order"
	"github.com/gin-gonic/gin"
)

// 新增订单
func AddOrder(context *gin.Context) {
	var form forms.AddOrderForm
	context.ShouldBind(&form)
	// 基础验证
	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	// 自定义逻辑验证
	if err := form.Valid(); err != nil {
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
		Receiver:          form.ExtraData.Receiver,
		ReceiverName:      form.ExtraData.Receiver.Name,
		ReceiverPhone:     form.ExtraData.Receiver.Phone,
		ReceiverAddress:   form.ExtraData.Receiver.Address,
		ReceiverTel:       form.ExtraData.Receiver.Tel,
		Sender:            form.ExtraData.Sender,
		SenderCompanyName: form.ExtraData.Sender.CompanyName,
		SenderPhone:       form.ExtraData.Sender.Phone,
		SenderRemark:      form.ExtraData.Sender.Remark,
		FinanceID:         finance.ID,
		ProvinceId:        form.ExtraData.Province.ID,
		ProvinceName:      form.ExtraData.Province.Name,
		CityId:            form.ExtraData.City.ID,
		CityName:          form.ExtraData.City.Name,
		AreaId:            form.ExtraData.Area.ID,
		AreaName:          form.ExtraData.Area.Name,
		TotalPrice:        form.ExtraData.Price}

	// 保存修改
	models.DB.Save(&order)

	// 批量增加订单货物信息
	go order.AddDetails(form.Products)
	plugins.ApiExport(context).ApiExport()
	return
}

// 订单列表
func OrderList(context *gin.Context) {
	var form forms.OrderListForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	var orders []models_order.FinanceOrder
	var total int
	query := form.Query()

	query.Count(&total)
	query.Offset((form.Page - 1) * form.Limit).Limit(form.Limit).Find(&orders)

	plugins.ApiExport(context).ListPageExport(orders, form.Page, total)
}

// 订单详情
func OrderInfo(context *gin.Context) {
	var form forms.OrderInfo
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	var order models_order.FinanceOrder

	query := form.Query()
	query.Find(&order)
	export := plugins.ApiExport(context)
	export.SetData("order", order)
	export.ApiExport()
}

// 编辑订单
func OrderEdit(context *gin.Context) {
	var form forms.OrderEditForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	if err := form.Valid(); err != nil {
		plugins.ApiExport(context).Error(5011, err.Error())
		return
	}

	// 获取新增订单财务人信息
	finance, err := common.GetFinance(context)
	if err != nil {
		plugins.ApiExport(context).Error(4005, "用户未登录,请在登录后尝试.")
		return
	}

	form.Order.Receiver = form.ExtraData.Receiver
	form.Order.ReceiverName = form.ExtraData.Receiver.Name
	form.Order.ReceiverPhone = form.ExtraData.Receiver.Phone
	form.Order.ReceiverAddress = form.ExtraData.Receiver.Address
	form.Order.ReceiverTel = form.ExtraData.Receiver.Tel
	form.Order.Sender = form.ExtraData.Sender
	form.Order.SenderCompanyName = form.ExtraData.Sender.CompanyName
	form.Order.SenderPhone = form.ExtraData.Sender.Phone
	form.Order.SenderRemark = form.ExtraData.Sender.Remark
	form.Order.FinanceID = finance.ID
	form.Order.ProvinceId = form.ExtraData.Province.ID
	form.Order.ProvinceName = form.ExtraData.Province.Name
	form.Order.CityId = form.ExtraData.City.ID
	form.Order.CityName = form.ExtraData.City.Name
	form.Order.AreaId = form.ExtraData.Area.ID
	form.Order.AreaName = form.ExtraData.Area.Name

	// 保存修改
	models.DB.Save(&form.Order)

	// 先删除旧货物详情再添加
	form.Order.DeleteAllDetail()
	go form.Order.AddDetails(form.Products)

	plugins.ApiExport(context).ApiExport()
}

// 删除订单
func OrderDelete(context *gin.Context) {

}
