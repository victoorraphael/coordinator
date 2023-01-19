package entities

type School struct {
	ID      int64       `json:"id,omitempty"`
	Name    string      `json:"name,omitempty"`
	Address Address     `json:"address"`
	Classes []Classroom `json:"classes,omitempty"`
}
