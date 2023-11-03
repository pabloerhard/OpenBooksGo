package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	GoogleId string             `json:"google_id,omitempty"`
}
