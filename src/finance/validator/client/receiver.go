package client

import (
	"finance/models"
	models_receiver "finance/models/receiver"
	"finance/validator"
	"fmt"
)

type ReceiverIdBase struct {
	ReceiverId uint `json:"receiver_id" form:"receiver_id" validate:"required" error_message:"收货人编号~required:请填写后重试."`
}

// 获取收货人
func (form *ReceiverIdBase) GetReceiver() *models_receiver.FinanceReceiver {
	var receiver models_receiver.FinanceReceiver
	models.DB.First(&receiver, form.ReceiverId)
	return &receiver
}

type ReceiverList struct {
	validator.ListPage
	Name  string `json:"name" form:"name"`
	Phone string `json:"phone" form:"phone"`
}

// 查询
func (form *ReceiverList) Query() []models_receiver.FinanceReceiver {
	var receivers []models_receiver.FinanceReceiver
	query := models.DB.Model(models_receiver.FinanceReceiver{})

	if form.Name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", form.Name))

	}

	if form.Phone != "" {
		query = query.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", form.Phone))

	}
	query.Count(&form.Total)
	query.Offset((form.Page - 1) * form.Limit).Limit(form.Limit).Find(&receivers)
	return receivers

}

type ReceiverInfo struct {
	ReceiverIdBase
}

type ReceiverEdit struct {
	ReceiverIdBase
	Name     string `json:"name" form:"name" validate:"required" error_message:"收货人名~required:请填写后重试."`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Tel      string `json:"tel" form:"tel"`
	AutoFill int    `json:"auto_fill" form:"auto_fill"`
	IdCard   string `json:"id_card" form:"id_card"`
}
