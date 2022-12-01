package entities

type Student struct {
	Person
}

func NewStudent() Student {
	return Student{Person{
		Type: PersonStudent,
	}}
}
