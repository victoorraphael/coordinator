package person

import (
	"github.com/victoorraphael/coordinator/internal/address"
)

type Builder struct {
	*Person
}

// NewPersonBuilder returns a builder to person
func NewPersonBuilder() *Builder {
	return &Builder{&Person{}}
}

func (b *Builder) As(personType Type) *Builder {
	b.Type = personType
	return b
}

func (b *Builder) LivesAt(address address.Address) *Builder {
	b.Address = address
	return b
}

func (b *Builder) Build() Person {
	return *b.Person
}
