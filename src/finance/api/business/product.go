package business

import (
	"finance/models"
	"finance/models/receiver"
	plugins "finance/plugins/common"
	"finance/validator"
	"finance/validator/client"
	"github.com/gin-gonic/gin"
)

// 货物列表
func ProductList(context *gin.Context) {

	var form client.ProductListForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	list := form.Query()

	productJson := make([]map[string]interface{}, 0)

	for _, item := range list {
		productJson = append(productJson, item.ToJson())
	}

	plugins.ApiExport(context).ListPageExport(productJson, form.Page, form.Total)
	return
}

// 添加货物
func ProductAdd(context *gin.Context) {

	var form client.AddProductForm
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

	product := receiver.FinanceReceiverProduct{Name: form.Name, Unit: form.Unit, Price: form.Price, ReceiverId: form.ReceiverId}

	models.DB.Save(&product)

	export := plugins.ApiExport(context)
	export.SetData("product_id", product.ID)
	export.ApiExport()
	return

}

// 货物详情
func ProductInfo(context *gin.Context) {
	var form client.ProductInfoForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	product := form.GetProduct()

	if product.ID == 0 {
		plugins.ApiExport(context).Error(5011, "产品编号错误")
		return
	}

	export := plugins.ApiExport(context)
	export.SetData("product", product.ToJson())
	export.ApiExport()
	return

}

// 货物查询
func ProductQuery(context *gin.Context) {
	var form client.ProductQueryForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	products := form.Query()

	productJson := make([]map[string]interface{}, 0)

	for _, item := range products {
		productJson = append(productJson, item.ToJson())
	}

	export := plugins.ApiExport(context)
	export.SetData("items", productJson)
	export.ApiExport()
	return

}

// 编辑货物
func ProductEdit(context *gin.Context) {

	var form client.ProductEditForm
	context.ShouldBindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	product := form.GetProduct()

	if product.ID == 0 {
		plugins.ApiExport(context).Error(5011, "产品编号错误")
		return
	}

	product.Price = form.Price
	product.Unit = form.Unit

	models.DB.Save(&product)

	plugins.ApiExport(context).ApiExport()
	return
}

// 删除货物
func ProductDelete(context *gin.Context) {

	var form client.ProductEditForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	product := form.GetProduct()

	if product.ID == 0 {
		plugins.ApiExport(context).Error(5011, "产品编号错误")
		return
	}

	models.DB.Unscoped().Delete(&product)

	plugins.ApiExport(context).ApiExport()
	return
}
