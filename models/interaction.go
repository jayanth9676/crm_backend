package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Interaction struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    CustomerID  primitive.ObjectID `bson:"customer_id" json:"customer_id"`
    Type        string             `bson:"type" json:"type"` // e.g., "ticket", "meeting"
    Status      string             `bson:"status" json:"status"`
    ScheduledAt string             `bson:"scheduled_at,omitempty" json:"scheduled_at,omitempty"`
    CreatedAt   string             `bson:"created_at" json:"created_at"`
    Notes       string             `bson:"notes,omitempty" json:"notes,omitempty"`
}