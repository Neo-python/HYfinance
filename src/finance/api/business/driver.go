package business

import (
	"finance/models"
	models_driver "finance/models/driver"
	plugins "finance/plugins/common"
	"finance/validator"
	forms "finance/validator/driver"
	"github.com/gin-gonic/gin"
	"strings"
)

// 添加驾驶员
func AddDriver(context *gin.Context) {
	var form forms.DriverAddForm
	context.ShouldBindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	driver := models_driver.FinanceDriver{
		Name:        form.Name,
		NumberPlate: form.NumberPlate,
		Phone:       form.Phone}

	// 保存修改
	models.DB.Save(&driver)

	plugins.ApiExport(context).ApiExport()
	return
}

// 驾驶员详情
func DriverInfo(context *gin.Context) {
	var form forms.DriverInfoForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	driver := form.GetDriver()

	if driver.ID == 0 {
		plugins.ApiExport(context).Error(5011, "驾驶员编号错误")
		return
	}
	export := plugins.ApiExport(context)

	export.SetData("driver", driver.ToJson())
	export.ApiExport()
	return
}

// 编辑驾驶员
func DriverEdit(context *gin.Context) {
	var form forms.DriverEditForm
	context.ShouldBindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	driver := form.GetDriver()

	if driver.ID == 0 {
		plugins.ApiExport(context).Error(5011, "驾驶员编号错误")
		return
	}

	driver.Name = form.Name
	driver.NumberPlate = form.NumberPlate
	driver.Phone = form.Phone

	models.DB.Save(&driver)

	plugins.ApiExport(context).ApiExport()
	return

}

// 删除驾驶员
func DeleteDriver(context *gin.Context) {
	var form forms.DriverDeleteForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	driver := form.GetDriver()

	if driver.ID == 0 {
		plugins.ApiExport(context).Error(5011, "驾驶员编号错误")
		return
	}

	driver.Delete()

	plugins.ApiExport(context).ApiExport()
	return

}

// 驾驶员列表
func DriverList(context *gin.Context) {
	var form forms.DriverListForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	drivers := form.GetDrivers()

	driversJson := make([]map[string]interface{}, 0)

	for _, item := range drivers {
		driversJson = append(driversJson, item.ToJson())
	}

	plugins.ApiExport(context).ListPageExport(driversJson, form.Page, form.Total)
	return
}

// 添加驾驶员车次
func AddDriverTrips(context *gin.Context) {
	var form forms.TripsAddForm
	context.ShouldBindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	// 自定义逻辑验证
	if err := form.Valid(); err != nil {
		plugins.ApiExport(context).Error(1001, err.Error())
		return
	}

	trips := models_driver.FinanceDriverTrips{
		ProvinceId:   form.ProvinceId,
		ProvinceName: form.ProvinceName,
		Date:         form.ValidDate,
		DriverId:     form.DriverId,
		Remark:       form.Remark}

	// 保存修改
	models.DB.Save(&trips)

	plugins.ApiExport(context).ApiExport()
	return
}

// 驾驶员车次列表
func DriverTripsList(context *gin.Context) {
	var form forms.TripsListForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	tripss := form.GetTrips()

	tripss_json := make([]map[string]interface{}, 0)

	for _, item := range tripss {
		tripss_json = append(tripss_json, item.ToJson())
	}

	plugins.ApiExport(context).ListPageExport(tripss_json, form.Page, form.Total)
	return

}

// 驾驶员车次详情
func DriverTripsInfo(context *gin.Context) {
	var form forms.TripsInfoForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	trips := form.Trips()

	if trips.ID == 0 {
		plugins.ApiExport(context).Error(5011, "车次编号错误")
		return
	}

	export := plugins.ApiExport(context)

	export.SetData("trips", trips.ToJson())
	export.ApiExport()
	return
}

// 编辑驾驶员车次
func DriverTripsEdit(context *gin.Context) {
	var form forms.TripsEditForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	// 自定义逻辑验证
	if err := form.Valid(); err != nil {
		plugins.ApiExport(context).Error(1001, err.Error())
		return
	}

	trips := form.Trips()

	if trips.ID == 0 {
		plugins.ApiExport(context).Error(5011, "车次编号错误")
		return
	}

	trips.Remark = form.Remark
	trips.Date = form.ValidDate
	trips.ProvinceName = form.ProvinceName
	trips.ProvinceId = form.ProvinceId

	models.DB.Save(&trips)

	plugins.ApiExport(context).ApiExport()
	return

}

// 删除车次
func DeleteDriverTrips(context *gin.Context) {
	var form forms.TripsDeleteForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	trips := form.Trips()

	if trips.ID == 0 {
		plugins.ApiExport(context).Error(5011, "车次编号错误")
		return
	}

	trips.DeleteSelf()

	plugins.ApiExport(context).ApiExport()
	return
}

// 驾驶员车次订单列表
func DriverTripsOrderList(context *gin.Context) {
	var form forms.TripsOrderListForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	// 自定义逻辑验证
	if err := form.Valid(); err != nil {
		plugins.ApiExport(context).Error(5011, err.Error())
		return
	}

	trips := form.Trips()
	totalOrderCount := 0
	completeOrderCount := 0
	var totalExpectedAmount float64
	var totalActualAmount float64

	trips.GetDetails()
	detailsJson := make([]map[string]interface{}, 0)

	for _, item := range trips.Details {
		detailsJson = append(detailsJson, item.ToJson(context))

		totalOrderCount += 1
		if item.ExpectedAmount == item.ActualAmount && item.ExpectedAmount != 0 {
			completeOrderCount += 1
		}
		totalActualAmount += item.ActualAmount
		totalExpectedAmount += item.ExpectedAmount
	}

	export := plugins.ApiExport(context)
	export.SetData("items", detailsJson)
	export.SetData("total_actual_amount", totalActualAmount)
	export.SetData("total_expected_amount", totalExpectedAmount)
	export.SetData("total_order_count", totalOrderCount)
	export.SetData("complete_order_count", completeOrderCount)
	export.ApiExport()
	return
}

// 驾驶员车次添加订单
func DriverTripsAddOrder(context *gin.Context) {
	var form forms.AddTripsOrderForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	if err := form.Valid(); err != nil {
		plugins.ApiExport(context).Error(1001, err.Error())
		return
	}

	order := form.Order()

	trips_order := models_driver.FinanceDriverTripsDetails{
		TripsId:        form.TripsId,
		OrderId:        form.OrderId,
		ExpectedAmount: order.ExpectedAmount,
		ActualAmount:   order.ActualAmount}

	// 创建分配记录
	if err := models.DB.Save(&trips_order).Error; err != nil {
		// 错误内容包含Duplicate entry双重输入错误
		if index := strings.Index(err.Error(), "Duplicate entry"); index >= 0 {
			plugins.ApiExport(context).Error(1001, "此订单已分配,勿重复操作.")
			return
		}
		plugins.ApiExport(context).Error(1001, err.Error())
		return
	}

	// 修改订单分配状态
	order.EditAllocationStatus(1)

	plugins.ApiExport(context).ApiExport()
	return
}

// 驾驶员车次删除订单
func DriverTripsDeleteOrder(context *gin.Context) {
	var form forms.DeleteTripsOrderForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}
	details := form.GetTripsDetails()

	if details.ID == 0 {
		plugins.ApiExport(context).Error(5011, "车次订单未找到.")
		return
	}

	details.DeleteSelf()

	plugins.ApiExport(context).ApiExport()
	return

}

// 驾驶员车次修改金额
func DriverTripsEditOrderAmount(context *gin.Context) {
	var form forms.EditTripsOrderAmountForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	if err := form.Valid(); err != nil {
		plugins.ApiExport(context).Error(5011, err.Error())
		return
	}

	details := form.TripsDetails()

	details.ExpectedAmount = form.ExpectedAmount
	details.ActualAmount = form.ActualAmount

	models.DB.Save(&details)

	plugins.ApiExport(context).ApiExport()
	return

}
