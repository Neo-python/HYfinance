package driver

import (
	"finance/models"
	"github.com/jinzhu/gorm"
)

type FinanceDriver struct {
	gorm.Model
	Name        string `json:"name" gorm:"COMMENT:'名称'"`
	NumberPlate string `json:"number_plate" gorm:"COMMENT:'车牌号'"`
	Phone       string `json:"phone" gorm:"COMMENT:'手机号'"`
}

func (driver *FinanceDriver) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":           driver.ID,
		"name":         driver.Name,
		"number_plate": driver.NumberPlate,
		"phone":        driver.Phone}
}

// 删除驾驶员
func (driver *FinanceDriver) Delete() {
	var tripss []FinanceDriverTrips
	models.DB.Where("driver_id=?", driver.ID).Find(&tripss)
	models.DB.Delete(&tripss)
	models.DB.Delete(&driver)
}
