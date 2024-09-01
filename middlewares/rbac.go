package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "your_project/utils"
)

func RoleCheck(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("user_id") // Assume user_id is set in the context after authentication
        user, err := utils.GetUserByID(userID)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        for _, role := range allowedRoles {
            if user.Role == role {
                c.Next()
                return
            }
        }

        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
        c.Abort()
    }
}