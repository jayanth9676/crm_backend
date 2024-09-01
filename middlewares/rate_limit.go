package middlewares

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
    "net/http"
)

func RateLimiter(limit rate.Limit, burst int) gin.HandlerFunc {
    limiter := rate.NewLimiter(limit, burst)
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            c.Abort()
            return
        }
        c.Next()
    }
}