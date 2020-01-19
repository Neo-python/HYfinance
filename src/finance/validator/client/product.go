package client

import (
	"errors"
	"finance/models"
	models_receiver "finance/models/receiver"
	"finance/validator"
)

type ProductIdBase struct {
	ProductId uint `json:"product_id" form:"product_id" validate:"required" error_message:"产品编号~required:请填写后重试."`
}

func (form *ProductIdBase) GetProduct() *models_receiver.FinanceReceiverProduct {
	var product models_receiver.FinanceReceiverProduct

	models.DB.First(&product, form.ProductId)
	return &product
}

type ProductBase struct {
	Name  string `json:"name" form:"name" validate:"required" error_message:"产品名~required:请填写后重试."`
	Unit  int    `json:"unit" form:"unit"`
	Price int    `json:"price" form:"price"`
}

type AddProductForm struct {
	ReceiverIdBase
	ProductBase
}

// 添加产品验证
func (form *AddProductForm) Valid() error {

	record := 0
	models.DB.Model(models_receiver.FinanceReceiverProduct{}).Where("receiver_id=? and name=?", form.ReceiverId, form.Name).Count(&record)

	if record != 0 {
		return errors.New("产品名重复")
	}

	return nil
}

// 产品列表
type ProductListForm struct {
	validator.ListPage
	ReceiverIdBase
}

// 查询
func (form *ProductListForm) Query() []models_receiver.FinanceReceiverProduct {

	var list []models_receiver.FinanceReceiverProduct

	query := models.DB.Model(models_receiver.FinanceReceiverProduct{})
	query = query.Where("receiver_id=?", form.ReceiverId)

	query.Count(&form.Total)
	query = query.Offset((form.Page - 1) * form.Limit).Limit(form.Limit).Find(&list)

	return list
}

// 产品详情
type ProductInfoForm struct {
	ProductIdBase
}

// 产品编辑
type ProductEditForm struct {
	ProductIdBase
	ProductBase
}

// 删除产品
type ProductDeleteForm struct {
	ProductIdBase
}
