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
		//plugins.ApiExport(context).Error(4005, "用户未登录,请在登录后尝试.")
		plugins.ApiExport(context).Error(4005, err.Error())
		return
	}

	order := models_order.FinanceOrder{
		Receiver:          *form.ExtraData.Receiver,
		ReceiverName:      form.ExtraData.Receiver.Name,
		ReceiverPhone:     form.ExtraData.Receiver.Phone,
		ReceiverAddress:   form.ExtraData.Receiver.Address,
		ReceiverTel:       form.ExtraData.Receiver.Tel,
		Sender:            *form.ExtraData.Sender,
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
		Deliver:           form.Deliver,
		PaymentMethod:     form.PaymentMethod}

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
	query := form.Query()

	query.Count(&form.Total)
	query.Offset((form.Page - 1) * form.Limit).Limit(form.Limit).Find(&orders)

	orders_json := []map[string]interface{}{}
	for _, item := range orders {
		orders_json = append(orders_json, item.ToJson())
	}

	plugins.ApiExport(context).ListPageExport(orders_json, form.Page, form.Total)
}

// 订单详情
func OrderInfo(context *gin.Context) {
	var form forms.OrderInfo
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	order := form.Order()
	order.QueryDetails()
	if order.ID != 0 {
		export := plugins.ApiExport(context)
		export.SetData("order", order.ToJson())
		export.ApiExport()
	} else {
		plugins.ApiExport(context).Error(5011, "订单未找到")
	}

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
	order := form.Order()
	order.Receiver = *form.ExtraData.Receiver
	order.ReceiverName = form.ExtraData.Receiver.Name
	order.ReceiverPhone = form.ExtraData.Receiver.Phone
	order.ReceiverAddress = form.ExtraData.Receiver.Address
	order.ReceiverTel = form.ExtraData.Receiver.Tel
	order.Sender = *form.ExtraData.Sender
	order.SenderCompanyName = form.ExtraData.Sender.CompanyName
	order.SenderPhone = form.ExtraData.Sender.Phone
	order.SenderRemark = form.ExtraData.Sender.Remark
	order.FinanceID = finance.ID
	order.ProvinceId = form.ExtraData.Province.ID
	order.ProvinceName = form.ExtraData.Province.Name
	order.CityId = form.ExtraData.City.ID
	order.CityName = form.ExtraData.City.Name
	order.AreaId = form.ExtraData.Area.ID
	order.AreaName = form.ExtraData.Area.Name

	// 保存修改
	models.DB.Save(&order)

	// 先删除旧货物详情再添加
	order.DeleteAllDetail()
	go order.AddDetails(form.Products)

	plugins.ApiExport(context).ApiExport()
}

// 删除订单
func OrderDelete(context *gin.Context) {

	var form forms.OrderDeleteForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	order := form.Order()

	if order.ID == 0 {
		plugins.ApiExport(context).Error(5011, "订单编号未找到")
		return
	}

	order.DeleteAllDetail()
	models.DB.Unscoped().Delete(order)

	plugins.ApiExport(context).ApiExport()

}

// 查看订单预期收费与实际收费
func OrderAmount(context *gin.Context) {
	var form forms.OrderInfo
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	order := form.Order()
	if order.ID != 0 {
		export := plugins.ApiExport(context)
		export.SetData("expected_amount", order.ExpectedAmount)
		export.SetData("actual_amount", order.ActualAmount)
		export.ApiExport()
	} else {
		plugins.ApiExport(context).Error(5011, "订单未找到")
	}
}

// 修改订单金额
func OrderAmountEdit(context *gin.Context) {
	var form forms.OrderAmountEditForm
	context.BindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	order := form.Order()
	if order.ID != 0 {
		order.ExpectedAmount = form.ExpectedAmount
		order.ActualAmount = form.ActualAmount
		// 保存修改
		models.DB.Save(&order)
		export := plugins.ApiExport(context)
		export.ApiExport()
	} else {
		plugins.ApiExport(context).Error(5011, "订单未找到")
	}
}
