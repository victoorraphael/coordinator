package entities

import "github.com/google/uuid"

type School struct {
	ID      uuid.UUID `json:"id"`
	Address Address   `json:"address"`
	Classes []Class   `json:"classes"`
}
