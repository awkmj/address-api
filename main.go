package main

import (
	"github.com/awjmj/address-api/adapter/controller"
	"github.com/awjmj/address-api/infra/config"
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func main() {

	config.LoadEnv()
	app := gin.New()

	app.GET("/api/v1/address", GetAddress)
	app.GET("/api/health", GetHealth)
	app.Run(":3000")
}

func GetAddress(c *gin.Context) {

	addressController := controller.NewAddressController()
	addressController.Get(c.Writer, c.Request)
}

func GetHealth(c *gin.Context) {
	res := map[string]interface{}{}
	res["status"] = "UP"
	res["timestamp"] = time.Now()

	c.JSON(http.StatusOK, res)
}
