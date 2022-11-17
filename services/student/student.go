package student

import (
	"errors"

	"github.com/victoorraphael/school-plus-BE/infra/contracts"
	"github.com/victoorraphael/school-plus-BE/infra/contracts/repository"
	"github.com/victoorraphael/school-plus-BE/infra/entities"
)

var (
	ErrMissingID = errors.New("missing ID")
)

type StudentService struct {
	students repository.IStudentRepo
}

func New(adapters *entities.Adapters) contracts.IService[entities.Student] {
	return &StudentService{
		students: repository.NewStudentRepo(adapters),
	}
}

func (ss *StudentService) Add(s entities.Student) (entities.Student, error) {
	var response entities.Student
	uid, err := ss.students.Add(s)
	if err != nil {
		return response, err
	}

	response.ID = uid
	return response, nil
}

func (ss *StudentService) Get(s entities.Student) (entities.Student, error) {
	response, err := ss.students.FindOne(s)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (ss *StudentService) List() ([]entities.Student, error) {
	panic("implement me")
}

func (ss *StudentService) Update(s entities.Student) error {
	panic("implement me")
}

func (ss *StudentService) Delete(s entities.Student) error {
	panic("implement me")
}
