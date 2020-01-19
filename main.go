package main

import (
	"finance/api/account"
	"finance/api/business"
	"finance/api/common"
	_ "finance/models"
	_ "finance/models/init"
	"finance/plugins"
	"finance/plugins/jwt_auth"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	if plugins.Config.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(jwt_auth.Cors())
	BusinessApiRegistered(router)
	CommonApiRegistered(router)
	AccountApiRegistered(router)
	DriverApiRegistered(router)
	router.Run("127.0.0.1:8095")
}

// 注册业务api
func BusinessApiRegistered(engine *gin.Engine) {
	//open := engine.Group("/business/")
	relative_path := "/business"
	auth := engine.Group(relative_path, jwt_auth.JWTAuth())

	{
		auth.GET("/order/list/", business.OrderList)
		auth.GET("/order/info/", business.OrderInfo)
		auth.GET("/query_receiver/", business.QueryReceiver)
		auth.GET("/query_sender/", business.QuerySender)

	}
	{
		auth.POST("/order/add/", business.AddOrder)
		auth.POST("/order/edit/", business.OrderEdit)

	}

	{
		auth.DELETE("/order/delete/", business.OrderDelete)
	}

	level_auth := engine.Group(relative_path, jwt_auth.LevelAuth(2))

	{
		level_auth.GET("/order/info/amount/", business.OrderAmount)
		level_auth.GET("/sender/list/", business.SenderList)
		level_auth.GET("/sender/info/", business.SenderInfo)
		level_auth.GET("/receiver/list/", business.ReceiverList)
		level_auth.GET("/receiver/info/", business.ReceiverInfo)
		level_auth.GET("/receiver/product/list/", business.ProductList)
		level_auth.GET("/receiver/product/info/", business.ProductInfo)
		level_auth.GET("/receiver/product/query/", business.ProductQuery)
	}
	{
		level_auth.POST("/order/amount/edit/", business.OrderAmountEdit)
		level_auth.POST("/sender/edit/", business.SenderEdit)
		level_auth.POST("/receiver/edit/", business.ReceiverEdit)
		level_auth.POST("/receiver/product/add/", business.ProductAdd)
		level_auth.POST("/receiver/product/edit/", business.ProductEdit)
	}
	{
		level_auth.DELETE("/receiver/product/delete/", business.ProductDelete)
	}

}

// 注册通用api
func CommonApiRegistered(engine *gin.Engine) {
	open := engine.Group("")

	{
		open.POST("/send_sms/code/", common.SMSSend)
	}

	auth := engine.Group("", jwt_auth.JWTAuth())
	{
		auth.GET("/query_area/", common.QueryArea)
	}

}

// 注册账号api
func AccountApiRegistered(engine *gin.Engine) {
	open := engine.Group("/account")

	{
		open.POST("/registered/", account.Registered)
		open.POST("/edit_password/", account.EditPassword)
		open.POST("/sign_in/", account.SignIn)
	}

	auth := engine.Group("/account/", jwt_auth.JWTAuth())
	{
		auth.GET("/sign_out/", account.SignOut)
		auth.GET("/factory/get_token/", account.GetFactoryToken)
	}
}

// 注册驾驶员api
func DriverApiRegistered(engine *gin.Engine) {
	level_auth := engine.Group("/driver", jwt_auth.LevelAuth(2))

	{
		level_auth.POST("/add/", business.AddDriver)
		level_auth.POST("/edit/", business.DriverEdit)
		level_auth.POST("/trips/add/", business.AddDriverTrips)
		level_auth.POST("/trips/edit/", business.DriverTripsEdit)
		level_auth.POST("/trips/order/add/", business.DriverTripsAddOrder)
		level_auth.POST("/trips/order/amount/edit/", business.DriverTripsEditOrderAmount)
	}

	{
		level_auth.GET("/info/", business.DriverInfo)
		level_auth.GET("/list/", business.DriverList)
		level_auth.GET("/trips/info/", business.DriverTripsInfo)
		level_auth.GET("/trips/list/", business.DriverTripsList)
		level_auth.GET("/trips/order/list/", business.DriverTripsOrderList)
	}
	{
		level_auth.DELETE("/delete/", business.DeleteDriver)
		level_auth.DELETE("trips/delete/", business.DeleteDriverTrips)
		level_auth.DELETE("trips/order/delete/", business.DriverTripsDeleteOrder)
	}
}
