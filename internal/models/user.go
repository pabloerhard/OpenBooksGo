package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"name,omitempty"`
	LastName          string             `json:"last_name,omitempty"`
	Email             string             `json:"email,omitempty"`
	EncryptedPassword string             `json:"password,omitempty"`
	ReadBooks         []string           `json:"books"`
	ReadingBooks      []string           `json:"reading_books"`
	WantToReadBooks   []string           `json:"want_to_read_books"`
}
