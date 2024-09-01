package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string             `bson:"name" json:"name"`
    ContactInfo string             `bson:"contact_info" json:"contact_info"`
    Company     string             `bson:"company,omitempty" json:"company,omitempty"`
    Status      string             `bson:"status" json:"status"`
    Notes       string             `bson:"notes,omitempty" json:"notes,omitempty"`
}