package api

import (
	"finance/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestApi struct {
	Name string `valid:"required"`
}

var Name string = "neo"

func Test(c *gin.Context) {
	fmt.Println("api.test")
	fmt.Println(c.Param("name"))
	fmt.Println(c.Request.Form)
	fmt.Println(c.Request.Body)
	fmt.Println(c.Request.ParseForm())
	fmt.Println(c.Request.FormValue("name"))
	fmt.Println(c.Request.MultipartForm)
	user := models.User{Name: "gao"}
	models.DB.Create(&user)
	safe_user := models.SafeUser{Id: user.ID, Name: user.Name}
	c.JSON(http.StatusOK, safe_user)
}
