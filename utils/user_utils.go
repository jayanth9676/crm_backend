package utils

import (
    "context"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "your_project/models"
)

// CreateUser creates a new user in the database
func CreateUser(user *models.User) error {
    collection := GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, user)
    return err
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id string) (*models.User, error) {
    collection := GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user models.User
    objID, _ := primitive.ObjectIDFromHex(id)
    err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(id string, updatedUser *models.User) error {
    collection := GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID}
    update := bson.M{"$set": updatedUser}

    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return errors.New("no user found with the given ID")
    }
    return nil
}

// DeleteUser deletes a user from the database
func DeleteUser(id string) error {
    collection := GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID}

    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return errors.New("no user found with the given ID")
    }
    return nil
}