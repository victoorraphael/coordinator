package domain

type Student struct {
	Classroom Classroom
	School    School
	Person
}

// NewStudent returns a new Student as pointer
func NewStudent() *Student {
	return &Student{
		Classroom: Classroom{},
		School:    School{},
		Person:    Person{},
	}
}
