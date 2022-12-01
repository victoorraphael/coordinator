package service

import "github.com/victoorraphael/school-plus-BE/internal/entities"

type IStudentSRV interface {
	Add(s entities.Student) (entities.Student, error)
	List() ([]entities.Student, error)
	Get(s entities.Student) (entities.Student, error)
}

func NewStudentService(adapters *entities.Adapters) IStudentSRV {
	return &studentsrv{adapters: adapters}
}

type studentsrv struct {
	adapters *entities.Adapters
}

func (srv *studentsrv) Add(s entities.Student) (entities.Student, error) {

}

func (srv *studentsrv) List() ([]entities.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (srv *studentsrv) Get(s entities.Student) (entities.Student, error) {
	//TODO implement me
	panic("implement me")
}
