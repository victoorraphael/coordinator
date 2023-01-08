package domain

type Professor struct {
	School School `json:"school"`
	Person
}
