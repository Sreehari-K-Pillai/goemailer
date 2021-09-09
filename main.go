package main

import (
	"15_day_project/mailandotp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/otpmail", handlers.OtpMail)
	r.Run(":8080")

}
