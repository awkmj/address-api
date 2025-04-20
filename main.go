package main

import (
	"github.com/awjmj/address-api/adapter/controller"
	"github.com/awjmj/address-api/infra/config"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	app := gin.New()

	app.GET("/address", GetAddress)
	app.Run(":3000")
}

func GetAddress(c *gin.Context) {

	addressController := controller.NewAddressController()
	addressController.Get(c.Writer, c.Request)
}
