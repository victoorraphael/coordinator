package entities

type Student struct {
	ClassroomID int `json:"classroom_id"`
	SchoolID    int `json:"school_id"`
	Person
}

// NewStudent returns a new Student as pointer
func NewStudent() *Student {
	return &Student{
		Person: Person{},
	}
}
