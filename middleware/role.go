package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func RoleMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetHeader("X-User-Role")
        if role == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user role"})
            c.Abort()
            return
        }
        c.Set("role", role)
        c.Next()
    }
}
