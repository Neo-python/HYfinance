package driver

import (
	"errors"
	"finance/models"
	"finance/models/area"
	model_driver "finance/models/driver"
	"finance/validator"
	"finance/validator/order"
	"time"
)

type TripsDetailsIdBase struct {
	TripsOrderId     uint `json:"trips_order_id" form:"trips_order_id" validate:"required" error_message:"车次订单分配编号~required:请填写后重试."`
	FormTripsDetails model_driver.FinanceDriverTripsDetails
}

// 复用车次订单分配信息
func (form *TripsDetailsIdBase) TripsDetails() *model_driver.FinanceDriverTripsDetails {
	if form.FormTripsDetails.ID == 0 {
		form.FormTripsDetails = *form.GetTripsDetails()
	}
	return &form.FormTripsDetails
}

// 检查订单分配信息是否存在
func (form *TripsDetailsIdBase) CheckTripsDetails() error {
	if form.TripsDetails().ID == 0 {
		return errors.New("车次订单分配号不存在.")
	}
	return nil
}

func (form *TripsDetailsIdBase) GetTripsDetails() *model_driver.FinanceDriverTripsDetails {
	var trips_details model_driver.FinanceDriverTripsDetails
	models.DB.First(&trips_details, form.TripsOrderId)
	return &trips_details
}

type TripsIdBase struct {
	TripsId   uint `form:"trips_id" json:"trips_id" validate:"required" error_message:"车次编号~required:请填写后重试."`
	FormTrips model_driver.FinanceDriverTrips
}

// 复用车次数据
func (form *TripsIdBase) Trips() *model_driver.FinanceDriverTrips {
	if form.FormTrips.ID == 0 {
		form.FormTrips = *form.GetTrips()
	}

	return &form.FormTrips
}

// 检查车次数据是否存在
func (form *TripsIdBase) CheckTrips() error {
	if form.Trips().ID == 0 {
		return errors.New("车次不存在")
	}
	return nil
}

func (form *TripsIdBase) GetTrips() *model_driver.FinanceDriverTrips {
	var trips model_driver.FinanceDriverTrips
	models.DB.First(&trips, form.TripsId)
	return &trips
}

type TripsBase struct {
	ProvinceId   uint `json:"province_id" form:"province_id" validate:"required" error_message:"目的地(省级)编号~required:请填写后重试."`
	ProvinceName string
	Date         string `json:"date" form:"date" gorm:"COMMENT:'出发日期'" validate:"required" error_message:"出发日期~required:请填写后重试."`
	Remark       string `json:"remark" form:"remark"`
	ValidDate    time.Time
}

type TripsAddForm struct {
	TripsBase
	DriverIdBase
}

func (form *TripsAddForm) Valid() error {
	date, err := time.ParseInLocation("2006-01-02 15:04:05", form.Date, time.Local)
	form.ValidDate = date
	if err != nil {
		return err
	}
	var province area.Area
	models.DB.First(&province, form.ProvinceId)

	if province.ID == 0 || province.Level != 1 {
		return errors.New("目的地省级编号错误")
	} else {
		form.ProvinceName = province.Name
	}

	driver := form.GetDriver()

	if driver.ID == 0 {
		return errors.New("驾驶员编号错误.")
	}

	return nil
}

// 车次详情
type TripsInfoForm struct {
	TripsIdBase
}

// 车次编辑
type TripsEditForm struct {
	TripsIdBase
	TripsBase
}

func (form *TripsEditForm) Valid() error {
	date, err := time.ParseInLocation("2006-01-02 15:04:05", form.Date, time.Local)
	form.ValidDate = date
	if err != nil {
		return err
	}
	var province area.Area
	models.DB.First(&province, form.ProvinceId)

	if province.ID == 0 || province.Level != 1 {
		return errors.New("目的地省级编号错误")
	} else {
		form.ProvinceName = province.Name
	}

	return nil
}

// 车次列表
type TripsListForm struct {
	DriverId uint `form:"driver_id"`
	validator.ListPage
}

func (form *TripsListForm) GetTrips() []model_driver.FinanceDriverTrips {
	var tripss []model_driver.FinanceDriverTrips
	query := models.DB.Model(model_driver.FinanceDriverTrips{})
	if form.DriverId != 0 {
		query = query.Where("driver_id=?", form.DriverId)
	}
	query.Count(&form.Total)
	query = query.Offset((form.Page - 1) * form.Limit).Limit(form.Limit).Find(&tripss)

	return tripss

}

// 删除车次
type TripsDeleteForm struct {
	TripsIdBase
}

// 车次订单列表
type TripsOrderListForm struct {
	TripsIdBase
}

// 表单自定义验证逻辑
func (form *TripsOrderListForm) Valid() error {
	if err := form.CheckTrips(); err != nil {
		return err
	}
	return nil
}

// 车次添加订单
type AddTripsOrderForm struct {
	TripsIdBase
	order.OrderIdBase
}

func (form *AddTripsOrderForm) Valid() error {
	order := form.Order()
	if order.ID == 0 {
		return errors.New("订单编号错误")
	}

	if order.AllocationStatus == 1 {
		return errors.New("订单已处于分配状态")
	}
	trips := form.Trips()

	if trips.ID == 0 {
		return errors.New("车次编号错误")
	}
	return nil
}

// 车次删除订单
type DeleteTripsOrderForm struct {
	TripsDetailsIdBase
}

// 修改车次订单金额
type EditTripsOrderAmountForm struct {
	TripsDetailsIdBase
	ExpectedAmount float64 `json:"expected_amount"`
	ActualAmount   float64 `json:"actual_amount"`
}

// 自定义表单验证
func (form *EditTripsOrderAmountForm) Valid() error {
	if err := form.CheckTripsDetails(); err != nil {
		return err
	}
	return nil
}
