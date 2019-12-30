package driver

import "github.com/jinzhu/gorm"

type FinanceDriver struct {
	gorm.Model
	Name        string `json:"name"`
	NumberPlate string `json:"number_plate"`
	Phone       string `json:"phone"`
}

func (driver *FinanceDriver) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":           driver.ID,
		"name":         driver.Name,
		"number_plate": driver.NumberPlate,
		"phone":        driver.Phone}
}
