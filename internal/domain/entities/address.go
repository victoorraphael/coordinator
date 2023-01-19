package entities

type Address struct {
	ID     int64  `json:"id,omitempty"`
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	Zip    string `json:"zip,omitempty"`
	Number int64  `json:"number,omitempty"`
}
