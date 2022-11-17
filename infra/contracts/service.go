package contracts

type IService[T any] interface {
	Add(T) (T, error)
	Get(T) (T, error)
	List() ([]T, error)
	Update(T) error
	Delete(T) error
}
