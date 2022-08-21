package address

import (
	"github.com/victoorraphael/school-plus-BE/infra/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Get(primitive.ObjectID) (entities.Address, error)
	Add(*entities.Address) (string, error)
	Update(entities.Address) error
	Delete(primitive.ObjectID) error
}
