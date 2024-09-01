package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Analytics struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID      primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	InteractionCount int               `bson:"interaction_count" json:"interaction_count"`
	ConversionRate  float64            `bson:"conversion_rate" json:"conversion_rate"`
	LastUpdated     string             `bson:"last_updated" json:"last_updated"`
}