package business

import (
	"finance/models"
	models_receiver "finance/models/receiver"
	models_sender "finance/models/sender"
	"finance/plugins/common"
	plugins "finance/plugins/common"
	"finance/validator"
	"finance/validator/client"
	"finance/validator/order"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询发货人
func QuerySender(context *gin.Context) {
	var form order.QueryForm
	context.ShouldBind(&form)

	var db = models.DB.Where("auto_fill=1").Limit(10)
	var senders []models_sender.FinanceSender
	sendersJson := make([]models_sender.FinanceSenderJson, 0)
	if form.Name != "" {
		db = db.Where("company_name LIKE ?", fmt.Sprintf("%%%s%%", form.Name))
	}

	if form.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", form.Phone))
	}

	db.Find(&senders)

	if db.Error != nil {
		common.ApiExport(context).ApiExport()
		return
	}

	for _, sender := range senders {
		sendersJson = append(sendersJson, models_sender.FinanceSenderJson{
			Id:          sender.ID,
			CompanyName: sender.CompanyName,
			Phone:       sender.Phone,
			Remark:      sender.Remark})
	}

	export := common.ApiExport(context)
	export.SetData("items", sendersJson)
	export.ApiExport()
}

// 查询收货人
func QueryReceiver(context *gin.Context) {
	var form order.QueryForm
	context.ShouldBind(&form)

	var query = models.DB.Where("auto_fill=1").Limit(10)
	var receivers []models_receiver.FinanceReceiver
	receiversJson := make([]map[string]interface{}, 0)
	if form.Name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", form.Name))
	}

	if form.Phone != "" {
		query = query.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", form.Phone))
	}

	query.Find(&receivers)

	if query.Error != nil {
		common.ApiExport(context).ApiExport()
		return
	}

	for _, receiver := range receivers {
		receiversJson = append(receiversJson, receiver.ToJson())
	}

	export := common.ApiExport(context)
	export.SetData("items", receiversJson)
	export.ApiExport()

}

// 收货人列表
func ReceiverList(context *gin.Context) {
	var form client.ReceiverList

	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	receivers := form.Query()
	receiversJson := make([]map[string]interface{}, 0)

	for _, item := range receivers {
		receiversJson = append(receiversJson, item.ToJson())
	}

	plugins.ApiExport(context).ListPageExport(receiversJson, form.Page, form.Total)
	return

}

// 收货人详情
func ReceiverInfo(context *gin.Context) {
	var form client.ReceiverInfo
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	receiver := form.GetReceiver()

	if receiver.ID == 0 {
		plugins.ApiExport(context).Error(5011, "收货人无法查询")
		return
	}

	export := plugins.ApiExport(context)

	export.SetData("receiver", receiver.ToJson())
	export.ApiExport()
	return
}

// 编辑收货人
func ReceiverEdit(context *gin.Context) {
	var form client.ReceiverEdit
	context.ShouldBindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	receiver := form.GetReceiver()

	if receiver.ID == 0 {
		plugins.ApiExport(context).Error(5011, "收货人无法查询")
		return
	}

	receiver.Name = form.Name
	receiver.IdCard = form.IdCard
	receiver.Phone = form.Phone
	receiver.AutoFill = form.AutoFill
	receiver.Tel = form.Tel
	receiver.Address = form.Address

	models.DB.Save(&receiver)

	plugins.ApiExport(context).ApiExport()
	return
}

// 发货人列表
func SenderList(context *gin.Context) {
	var form client.SenderList

	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	receivers := form.Query()
	receiversJson := make([]map[string]interface{}, 0)

	for _, item := range receivers {
		receiversJson = append(receiversJson, item.ToJson())
	}

	plugins.ApiExport(context).ListPageExport(receiversJson, form.Page, form.Total)
	return

}

// 发货人详情
func SenderInfo(context *gin.Context) {
	var form client.SenderInfoForm
	context.ShouldBind(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	sender := form.GetSender()

	if sender.ID == 0 {
		plugins.ApiExport(context).Error(5011, "发货人无法查询")
		return
	}

	export := plugins.ApiExport(context)

	export.SetData("sender", sender.ToJson())
	export.ApiExport()
	return
}

// 编辑发货人
func SenderEdit(context *gin.Context) {
	var form client.SenderEditForm
	context.ShouldBindJSON(&form)

	if err := validator.Valid.Struct(&form); err != nil {
		plugins.ApiExport(context).FormError(err)
		return
	}

	receiver := form.GetSender()

	if receiver.ID == 0 {
		plugins.ApiExport(context).Error(5011, "发货人无法查询")
		return
	}

	receiver.CompanyName = form.CompanyName
	receiver.IdCard = form.IdCard
	receiver.Phone = form.Phone
	receiver.AutoFill = form.AutoFill
	receiver.Remark = form.Remark

	models.DB.Save(&receiver)

	plugins.ApiExport(context).ApiExport()
	return
}
