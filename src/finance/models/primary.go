package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/jinzhu/gorm/dialects/mysql"

type Finance struct {
	gorm.Model
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `structs:",remove"`
}

func (finance *Finance) Privacy_fields() []string {
	return []string{"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Password"}
}

func (finance *Finance) ToJson() interface{} {
	if result, err := json.Marshal(finance); err != nil {
		fmt.Println(result, err)
		return err.Error()
	} else {
		fmt.Println(string(result))
		return string(result)
	}
}

var DB *gorm.DB

func init() {
	new_db, err := gorm.Open("mysql", "root:000000@tcp(127.0.0.1:3306)/HY?loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	new_db.SingularTable(true)
	new_db.AutoMigrate(&Finance{})
	DB = new_db

}

//func GetDB() *gorm.DB {
//	return DB.New()
//}
