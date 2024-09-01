package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
    Type      string             `bson:"type" json:"type"`
    Title     string             `bson:"title" json:"title"`
    Message   string             `bson:"message" json:"message"`
    IsRead    bool               `bson:"is_read" json:"is_read"`
    CreatedAt string             `bson:"created_at" json:"created_at"`
}