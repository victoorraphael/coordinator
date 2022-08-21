package memory

import (
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/victoorraphael/school-plus-BE/internal/repositories/student"
)

type MemoryRepository struct {
	students map[primitive.ObjectID]student.Student
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		students: make(map[primitive.ObjectID]student.Student),
	}
}

func (store *MemoryRepository) Get(id primitive.ObjectID) (student.Student, error) {
	if s, ok := store.students[id]; ok {
		return s, nil
	}

	return student.Student{}, student.ErrStudentNotFound
}

func (store *MemoryRepository) Add(s student.Student) error {
	if store.students == nil {
		store.Lock()
		store.students = make(map[primitive.ObjectID]student.Student)
		store.Unlock()
	}

	if _, ok := store.students[s.ID]; ok {
		return student.ErrFailedAddStudent
	}

	store.Lock()
	store.students[s.ID] = s
	store.Unlock()

	return nil
}

func (store *MemoryRepository) Update(s student.Student) error {
	if _, ok := store.students[s.ID]; !ok {
		return student.ErrStudentNotFound
	}

	store.students[s.ID] = s
	return nil
}

func (store *MemoryRepository) Delete(id primitive.ObjectID) error {
	if _, ok := store.students[id]; !ok {
		return student.ErrStudentNotFound
	}

	delete(store.students, id)
	return nil
}
