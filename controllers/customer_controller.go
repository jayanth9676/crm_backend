package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "crm_backend/models"
    "crm_backend/utils"
)

func CreateCustomer(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Save customer to the database
    if err := utils.CreateCustomer(&customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create customer"})
        return
    }
    c.JSON(http.StatusCreated, customer)
}

// Implement GetCustomer, UpdateCustomer, and DeleteCustomer functions similarly