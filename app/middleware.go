package app

import (
	"github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(1, 3) // 1 request per second with burst of 3


func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-Role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CustomerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-Role")
		if role != "customer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Customer access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}


func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests"})
			return
		}
		c.Next()
	}
}
