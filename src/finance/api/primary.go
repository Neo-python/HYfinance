package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestApi struct {
	Name string `valid:"required"`
}

var Name string = "neo"

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
