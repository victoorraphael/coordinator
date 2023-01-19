package entities

type Professor struct {
	School School `json:"school"`
	Person
}
