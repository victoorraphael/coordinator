package domain

type School struct {
	ID      int64
	Name    string
	Address Address
	Classes []Classroom
}
