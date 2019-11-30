package business

import (
	"finance/models"
	models_receiver "finance/models/receiver"
	models_sender "finance/models/sender"
	"finance/plugins/common"
	"finance/validator/order"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询发货人
func QuerySender(context *gin.Context) {
	var form order.QueryForm
	context.ShouldBind(&form)

	var db = models.DB.Limit(10)
	var senders []models_sender.FinanceSender
	var senders_json []models_sender.FinanceSenderJson
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
		senders_json = append(senders_json, models_sender.FinanceSenderJson{
			CompanyName: sender.CompanyName,
			Phone:       sender.Phone,
			Remark:      sender.Remark})
	}

	export := common.ApiExport(context)
	export.SetData("senders", senders_json)
	export.ApiExport()
}

// 查询收货人
func QueryReceiver(context *gin.Context) {
	var form order.QueryForm
	context.ShouldBind(&form)

	var db = models.DB.Limit(10)
	var receivers []models_receiver.FinanceReceiver
	var receivers_json []models_receiver.FinanceReceiverJson
	if form.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", form.Name))
	}

	if form.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", form.Phone))
	}

	db.Find(&receivers)

	if db.Error != nil {
		common.ApiExport(context).ApiExport()
		return
	}

	for _, receiver := range receivers {
		receivers_json = append(receivers_json, models_receiver.FinanceReceiverJson{
			Name:    receiver.Name,
			Phone:   receiver.Phone,
			Address: receiver.Address,
			Tel:     receiver.Tel})
	}

	export := common.ApiExport(context)
	export.SetData("receivers", receivers_json)
	export.ApiExport()

}
