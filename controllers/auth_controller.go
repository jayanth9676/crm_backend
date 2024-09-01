package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "your_project/models"
    "your_project/utils"
)

// Login handles user authentication
func Login(c *gin.Context) {
    var loginData struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // TODO: Retrieve user from database by email
    // For now, we'll use a mock user
    user := &models.User{
        ID:       "mock_user_id",
        Email:    loginData.Email,
        Password: "$2a$14$eSfaEYGZJXy5Gu.u5Jm/6.qJIJNr1Jx4kCkioQj8OMq7n3Q8Ldpuy", // hashed "password123"
    }

    if !utils.CheckPasswordHash(loginData.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}