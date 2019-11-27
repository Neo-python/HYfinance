package models

import (
	"finance/models/area"
	"finance/models/finance"
	"finance/models/order"
	"finance/models/receiver"
	"finance/models/sender"
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

	// area
	new_db.AutoMigrate(&area.Area{})

	// Finance
	new_db.AutoMigrate(&finance.Finance{})

	// FinanceOrder
	new_db.AutoMigrate(&order.FinanceOrder{})
	new_db.Model(&order.FinanceOrder{}).AddForeignKey("province_id", "area(id)", "no action", "no action")
	new_db.Model(&order.FinanceOrder{}).AddForeignKey("city_id", "area(id)", "no action", "no action")
	new_db.Model(&order.FinanceOrder{}).AddForeignKey("area_id", "area(id)", "no action", "no action")
	new_db.Model(&order.FinanceOrder{}).AddForeignKey("finance_id", "finance(id)", "no action", "no action")
	new_db.Model(&order.FinanceOrder{}).AddForeignKey("receiver_id", "finance_receiver(id)", "no action", "no action")
	new_db.Model(&order.FinanceOrder{}).AddForeignKey("sender_id", "finance_sender(id)", "no action", "no action")

	// FinanceOrderDetail
	new_db.AutoMigrate(&order.FinanceOrderDetail{})
	new_db.Model(&order.FinanceOrderDetail{}).AddForeignKey("order_id", "finance_order(id)", "no action", "no action")

	// FinanceReceiver
	new_db.AutoMigrate(&receiver.FinanceReceiver{})

	// FinanceSender
	new_db.AutoMigrate(&sender.FinanceSender{})

	DB = new_db

}
