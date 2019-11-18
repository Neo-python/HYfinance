package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/jinzhu/gorm/dialects/mysql"

type (
	User struct {
		gorm.Model
		Name string `gorm:"not_null"`
	}

	SafeUser struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	}
)

var DB *gorm.DB

func init() {
	new_db, err := gorm.Open("mysql", "root:000000@tcp(127.0.0.1:3306)/HY")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	new_db.AutoMigrate(&User{})
	new_db.SingularTable(true)
	DB = new_db

}

func GetDB() *gorm.DB {
	return DB.New()
}
