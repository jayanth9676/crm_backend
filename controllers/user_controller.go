package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"crm_backend/models"
	"crm_backend/utils"
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

// Implement GetUser, UpdateUser, DeleteUser similarly