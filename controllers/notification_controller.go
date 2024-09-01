package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "crm_backend/models"
    "crm_backend/utils"
)

func CreateNotification(c *gin.Context) {
    var notification models.Notification
    if err := c.ShouldBindJSON(&notification); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    notification.CreatedAt = time.Now().Format(time.RFC3339)
    if err := utils.CreateNotification(&notification); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create notification"})
        return
    }
    c.JSON(http.StatusCreated, notification)
}

// Implement GetNotifications, MarkNotificationAsRead similarly