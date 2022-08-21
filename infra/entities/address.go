package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Street string             `json:"street" bson:"street"`
	City   string             `json:"city" bson:"city"`
	Zip    string             `json:"zip" bson:"zip"`
	Number int                `json:"number" bson:"number"`
}
