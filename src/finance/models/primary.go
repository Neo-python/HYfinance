package models

import (
	"finance/plugins"
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var DB *gorm.DB

func init() {
	new_db, err := gorm.Open("mysql", plugins.Config.MysqlUrl)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}
	new_db.SingularTable(true)

	DB = new_db

}
