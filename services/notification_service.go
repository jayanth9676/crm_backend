package services

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "crm_backend/models"
    "crm_backend/utils"
)

// CreateNotification creates a new notification in the database
func CreateNotification(notification *models.Notification) error {
    collection := utils.GetCollection("notifications")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, notification)
    return err
}

// GetNotifications retrieves notifications for a user
func GetNotifications(userID string) ([]models.Notification, error) {
    collection := utils.GetCollection("notifications")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var notifications []models.Notification
    cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
    if err != nil {
        return nil, err
    }
    if err = cursor.All(ctx, &notifications); err != nil {
        return nil, err
    }
    return notifications, nil
}

// MarkNotificationAsRead marks a notification as read
func MarkNotificationAsRead(id string) error {
    collection := utils.GetCollection("notifications")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID}
    update := bson.M{"$set": bson.M{"is_read": true}}

    _, err := collection.UpdateOne(ctx, filter, update)
    return err
}