package sender

import "github.com/jinzhu/gorm"

type Sender struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
	Remark      string `json:"remark"`
}
