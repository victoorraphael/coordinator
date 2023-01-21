package repository

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/database"
)

type IUserRepository interface {
	FindEmail(ctx context.Context, email string) (entities.User, error)
	Add(ctx context.Context, user *entities.User) error
}

type user struct {
	pool database.DBPool
}

func NewUserRepo(pool database.DBPool) IUserRepository {
	return &user{pool}
}

func (u *user) FindEmail(ctx context.Context, email string) (entities.User, error) {
	conn, err := u.pool.Acquire()
	if err != nil {
		return entities.User{}, err
	}
	defer u.pool.Release(conn)

	var res entities.User
	_, errSelect := conn.Select("users.id, users.email, users.password, persons.uuid").
		From("users").
		Join("persons", "persons.email = users.email").
		Where("users.email = ?", email).
		LoadContext(ctx, &res)

	return res, errSelect
}

func (u *user) Add(ctx context.Context, user *entities.User) error {
	conn, err := u.pool.Acquire()
	if err != nil {
		return err
	}
	defer u.pool.Release(conn)

	_, errIns := conn.InsertInto("users").
		Pair("email", user.Email).
		Pair("password", user.PasswordHash).
		ExecContext(ctx)

	return errIns
}
