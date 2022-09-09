package contracts

type IService[T any] interface {
	Add()
	Get()
	List()
	Update()
	Delete()
}
