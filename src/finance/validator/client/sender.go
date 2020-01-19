package client

import (
	"finance/models"
	models_sender "finance/models/sender"
	"finance/validator"
	"fmt"
)

type SenderIdBase struct {
	SenderId uint `json:"sender_id" form:"sender_id" validate:"required" error_message:"发货人编号~required:请填写后重试."`
}

type SenderBase struct {
	CompanyName string `json:"company_name" form:"company_name"`
	Phone       string `json:"phone" form:"phone"`
	Remark      string `json:"remark" form:"remark"`
	AutoFill    int    `json:"auto_fill" form:"auto_fill"`
	IdCard      string `json:"id_card" form:"id_card"`
}

// 获取收货人
func (form *SenderIdBase) GetSender() *models_sender.FinanceSender {
	var sender models_sender.FinanceSender
	models.DB.First(&sender, form.SenderId)
	return &sender
}

type SenderList struct {
	validator.ListPage
	CompanyName string `json:"company_name" form:"company_name"`
	Phone       string `json:"phone" form:"phone"`
}

// 查询
func (form *SenderList) Query() []models_sender.FinanceSender {
	var receivers []models_sender.FinanceSender
	query := models.DB.Model(models_sender.FinanceSender{})

	if form.CompanyName != "" {
		query = query.Where("company_name LIKE ?", fmt.Sprintf("%%%s%%", form.CompanyName))

	}

	if form.Phone != "" {
		query = query.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", form.Phone))

	}
	query.Count(&form.Total)
	query.Offset((form.Page - 1) * form.Limit).Limit(form.Limit).Find(&receivers)
	return receivers

}

type SenderInfoForm struct {
	SenderIdBase
}

type SenderEditForm struct {
	SenderIdBase
	SenderBase
}
