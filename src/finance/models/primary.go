package models

import (
	"finance/models/finance"
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var DB *gorm.DB

func init() {
	new_db, err := gorm.Open("mysql", "root:000000@tcp(127.0.0.1:3306)/HY?loc=Local&parseTime=true")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	new_db.SingularTable(true)
	new_db.AutoMigrate(&finance.Finance{})
	DB = new_db

}
