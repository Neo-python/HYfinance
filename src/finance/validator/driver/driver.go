package driver

import (
	"finance/models"
	models_driver "finance/models/driver"
	"finance/validator"
)

type DriverIdBase struct {
	DriverId uint `json:"driver_id" form:"driver_id" validate:"required" error_message:"驾驶员编号~required:请填写后重试."`
}

func (form *DriverIdBase) GetDriver() *models_driver.FinanceDriver {
	var driver models_driver.FinanceDriver
	models.DB.First(&driver, form.DriverId)
	return &driver

}

type DriverAddForm struct {
	Name        string `json:"name" form:"name" validate:"required,max=20" error_message:"驾驶员姓名~required:此字段必须填写;max:最大长度为20"`
	NumberPlate string `json:"number_plate" form:"number_plate" validate:"required,max=20" error_message:"驾驶员车牌号~required:此字段必须填写;max:最大长度为20"` //`form:"number_plate" validate:"required,max=20" error_message:"驾驶员车牌号~required:此字段必须填写;max:最大长度为20"`
	Phone       string `json:"phone" form:"phone" validate:"required,max=12" error_message:"驾驶员手机号~required:此字段必须填写;max:最大长度为12"`
}

type DriverListForm struct {
	validator.ListPage
}

func (form *DriverListForm) GetDrivers() []models_driver.FinanceDriver {
	var drivers []models_driver.FinanceDriver
	models.DB.Model(models_driver.FinanceDriver{}).Count(&form.Total)
	models.DB.Offset((form.Page - 1) * form.Limit).Limit(form.Limit).Find(&drivers)
	return drivers

}

type DriverInfoForm struct {
	DriverIdBase
}

type DriverEditForm struct {
	DriverIdBase
	DriverAddForm
}

type DriverDeleteForm struct {
	DriverIdBase
}
