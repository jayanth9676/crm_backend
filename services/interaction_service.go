package services

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "crm_backend/models"
    "crm_backend/utils"
)

// CreateInteraction creates a new interaction in the database
func CreateInteraction(interaction *models.Interaction) error {
    collection := utils.GetCollection("interactions")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, interaction)
    return err
}

// GetInteractionByID retrieves an interaction by its ID
func GetInteractionByID(id string) (*models.Interaction, error) {
    collection := utils.GetCollection("interactions")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var interaction models.Interaction
    objID, _ := primitive.ObjectIDFromHex(id)
    err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&interaction)
    if err != nil {
        return nil, err
    }
    return &interaction, nil
}

// UpdateInteraction updates an existing interaction in the database
func UpdateInteraction(id string, updatedInteraction *models.Interaction) error {
    collection := utils.GetCollection("interactions")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID}
    update := bson.M{"$set": updatedInteraction}

    _, err := collection.UpdateOne(ctx, filter, update)
    return err
}

// DeleteInteraction deletes an interaction from the database
func DeleteInteraction(id string) error {
    collection := utils.GetCollection("interactions")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID}

    _, err := collection.DeleteOne(ctx, filter)
    return err
}