package entities

import "github.com/google/uuid"

type Person struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Address Address   `json:"address"`
}
