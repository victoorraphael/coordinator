package student

import "context"

type Repository interface {
	Find(ctx context.Context, schoolID int, classroomsID ...int) ([]Student, error)
	FindUUID(ctx context.Context, uuid string) (Student, error)
	Create(ctx context.Context, student Student) (Student, error)
	Update(ctx context.Context, student Student) error
	Delete(ctx context.Context, uuid string) error
}

type Service interface {
	List(ctx context.Context, schoolID int, classroomID int) ([]Student, error)
	Get(ctx context.Context, uuid string) (Student, error)
	Create(ctx context.Context, student Student) (Student, error)
	Update(ctx context.Context, student Student) error
	Delete(ctx context.Context, uuid string) error
}
