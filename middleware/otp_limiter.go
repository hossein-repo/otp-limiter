package middleware

import (
	"net"
	"net/http"
	"time"

	"otpLimiter/limiter"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter() gin.HandlerFunc {
	ipLimiter := limiter.NewIPRateLimiter(rate.Every(5*time.Second), 1)
	return func(c *gin.Context) {
		ip := getIP(c.Request.RemoteAddr)
		limiter := ipLimiter.GetLimiter(ip)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please wait.",
			})
			return
		}
		c.Next()
	}
}

func getIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return ip
}
