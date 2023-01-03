package student

type Repository interface {
	Find(schoolID int, classroomID int) ([]Student, error)
	FindUUID(uuid string) (Student, error)
	Create(student Student) (Student, error)
	Update(student Student) error
	Delete(uuid string) error
}

type Service interface {
	List(schoolID int, classroomID int) ([]Student, error)
	Get(uuid string) (Student, error)
	Create(student Student) (Student, error)
	Update(student Student) error
	Delete(uuid string) error
}
