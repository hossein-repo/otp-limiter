package router

import (
	"github.com/gin-gonic/gin"
	"otpLimiter/middleware"
)

func User(router *gin.RouterGroup) {
	router.POST("/send-otp", middleware.OtpLimiter(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OTP sent"})
	})
}
