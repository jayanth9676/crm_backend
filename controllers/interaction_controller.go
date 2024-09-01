package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "crm_backend/models"
    "crm_backend/utils"
)

func RaiseTicket(c *gin.Context) {
    var interaction models.Interaction
    if err := c.ShouldBindJSON(&interaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    interaction.Type = "ticket"
    interaction.Status = "unresolved"
    if err := utils.CreateInteraction(&interaction); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create ticket"})
        return
    }
    c.JSON(http.StatusCreated, interaction)
}

// Implement ResolveTicket, ScheduleMeeting, and GetInteractionHistory functions similarly