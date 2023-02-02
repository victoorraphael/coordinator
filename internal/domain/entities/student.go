package entities

type Student struct {
	ClassroomID int `db:"classroom_id"`
	SchoolID    int `db:"school_id"`
	Person
}

// NewStudent returns a new Student as pointer
func NewStudent() *Student {
	return &Student{
		Person: Person{},
	}
}
