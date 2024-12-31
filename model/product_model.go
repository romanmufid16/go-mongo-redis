package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Price    int64              `json:"price,omitempty" bson:"price,omitempty"`
	Category string             `json:"category,omitempty" bson:"category,omitempty"`
}
