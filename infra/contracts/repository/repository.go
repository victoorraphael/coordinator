package repository

import "github.com/google/uuid"

type IReadRepository[T any] interface {
	List() ([]T, error)
	FindOne(T) (T, error)
}

type IWriteRepository[T any] interface {
	Add(*T) (uuid.UUID, error)
	Update(*T) error
	Delete(T) error
}
