package init

import (
	"finance/models"
	"finance/models/area"
	"finance/models/driver"
	"finance/models/finance"
	"finance/models/order"
	"finance/models/receiver"
	"finance/models/sender"
)

func init() {

	// area
	models.DB.AutoMigrate(&area.Area{})

	// Finance
	models.DB.AutoMigrate(&finance.Finance{})

	// FinanceOrder
	models.DB.AutoMigrate(&order.FinanceOrder{})

	//new_db.Model(&order.FinanceOrder{}).AddForeignKey("province_id", "area(id)", "no action", "no action")
	//new_db.Model(&order.FinanceOrder{}).AddForeignKey("city_id", "area(id)", "no action", "no action")
	//new_db.Model(&order.FinanceOrder{}).AddForeignKey("area_id", "area(id)", "no action", "no action")
	models.DB.Model(&order.FinanceOrder{}).AddForeignKey("finance_id", "finance(id)", "no action", "no action")
	models.DB.Model(&order.FinanceOrder{}).AddForeignKey("receiver_id", "finance_receiver(id)", "no action", "no action")
	models.DB.Model(&order.FinanceOrder{}).AddForeignKey("sender_id", "finance_sender(id)", "no action", "no action")

	// FinanceOrderDetail
	models.DB.AutoMigrate(&order.FinanceOrderDetail{})
	models.DB.Model(&order.FinanceOrderDetail{}).AddForeignKey("order_id", "finance_order(id)", "no action", "no action")

	// FinanceReceiver
	models.DB.AutoMigrate(&receiver.FinanceReceiver{})

	// FinanceSender
	models.DB.AutoMigrate(&sender.FinanceSender{})

	// FinanceDriver
	models.DB.AutoMigrate(&driver.FinanceDriver{})
	models.DB.AutoMigrate(&driver.FinanceDriverTrips{})
	models.DB.AutoMigrate(&driver.FinanceDriverTripsDetails{})
	models.DB.Model(&driver.FinanceDriverTrips{}).AddForeignKey("driver_id", "finance_driver(id)", "no action", "no action")
	models.DB.Model(&driver.FinanceDriverTripsDetails{}).AddForeignKey("trips_id", "finance_driver_trips(id)", "no action", "no action")
	models.DB.Model(&driver.FinanceDriverTripsDetails{}).AddForeignKey("order_id", "finance_order(id)", "no action", "no action")
}
