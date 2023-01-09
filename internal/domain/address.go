package domain

type Address struct {
	ID     int64  `json:"id,omitempty"`
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	Zip    string `json:"zip,omitempty"`
	Number int    `json:"number,omitempty"`
}
