package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"your_project/models"
	"your_project/utils"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user to the database
	if err := utils.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Implement GetUser, UpdateUser, and DeleteUser functions similarly