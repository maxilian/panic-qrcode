package main

import (
	"panic-qrcode/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/generate-qr", controller.GenerateQR)
	r.POST("/multiple-qr", controller.GenerateMultipleQR)

	r.Run("0.0.0.0:8081")
}
