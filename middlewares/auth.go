package middlewares

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func RoleBasedAuth(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole := c.GetString("role")
        for _, role := range roles {
            if userRole == role {
                c.Next()
                return
            }
        }
        c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
        c.Abort()
    }
}