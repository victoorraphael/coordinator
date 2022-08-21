package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Email   string             `json:"email" bson:"email"`
	Phone   string             `json:"phone" bson:"phone"`
	Address Address            `json:"address" bson:"address"`
}
