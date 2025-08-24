package main

import (
	"otpLimiter/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api/v1")
	router.User(api)
	r.Run(":8080")
}
