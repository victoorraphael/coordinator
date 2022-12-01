package entities

type Student struct {
	Person
}

func New() Student {
	return Student{Person{
		Type: PersonStudent,
	}}
}
