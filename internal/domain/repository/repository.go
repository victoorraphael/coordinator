package repository

import (
	"github.com/victoorraphael/coordinator/pkg/database"
)

type Repo struct {
	Address   IAddressRepository
	Classroom IClassroomRepository
	Person    IPersonRepository
	School    ISchoolRepository
	Subject   ISubjectRepository
	User      IUserRepository
}

func New(pool database.DBPool) *Repo {
	return &Repo{
		Address:   NewAddressRepo(pool),
		Classroom: NewClassroomRepository(pool),
		Person:    NewPersonRepo(pool),
		School:    NewSchoolRespository(pool),
		Subject:   NewSubjectRepository(pool),
		User:      NewUserRepo(pool),
	}
}
