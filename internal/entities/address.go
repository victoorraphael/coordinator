package entities

import "github.com/google/uuid"

type Address struct {
	ID     uuid.UUID `json:"id" bson:"_id"`
	Street string    `json:"street" bson:"street"`
	City   string    `json:"city" bson:"city"`
	Zip    string    `json:"zip" bson:"zip"`
	Number int       `json:"number" bson:"number"`
}
