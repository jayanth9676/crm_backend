package utils

import (
    "context"
    "errors"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "your_project/models"
)

// CreateCustomer creates a new customer in the database
func CreateCustomer(customer *models.Customer) error {
    collection := GetCollection("customers")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, customer)
    return err
}

// GetCustomerByID retrieves a customer by their ID
func GetCustomerByID(id string) (*models.Customer, error) {
    collection := GetCollection("customers")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var customer models.Customer
    objID, _ := primitive.ObjectIDFromHex(id)
    err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&customer)
    if err != nil {
        return nil, err
    }
    return &customer, nil
}

// UpdateCustomer updates an existing customer in the database
func UpdateCustomer(id string, updatedCustomer *models.Customer) error {
    collection := GetCollection("customers")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID}
    update := bson.M{"$set": updatedCustomer}

    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return errors.New("no customer found with the given ID")
    }
    return nil
}

// DeleteCustomer deletes a customer from the database
func DeleteCustomer(id string) error {
    collection := GetCollection("customers")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID}

    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return errors.New("no customer found with the given ID")
    }
    return nil
}