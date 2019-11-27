package area

import "github.com/jinzhu/gorm"

type Area struct {
	gorm.Model
	SuperiorId int    `json:"superior_id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
}
