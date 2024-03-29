package services

import (
	"context"
	"errors"
	"github.com/gocraft/dbr/v2"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/pkg/jwttoken"
)

type IAuthenticationService interface {
	SignIn(ctx context.Context, user entities.UserLoginView) error
	Login(ctx context.Context, user entities.UserLoginView) (entities.UserLoginResponse, error)
	Logout()
}

type authentication struct {
	repo *repository.Repo
}

func NewAuthenticationService(repo *repository.Repo) IAuthenticationService {
	return &authentication{repo}
}

func (a *authentication) SignIn(ctx context.Context, u entities.UserLoginView) error {
	if err := u.Validate(); err != nil {
		return err
	}

	existentUser, err := a.repo.User.FindEmail(ctx, u.Email)
	if err != nil && !errors.Is(err, dbr.ErrNotFound) {
		return err
	}

	if existentUser.ID > 0 {
		return errors.New("email já cadastrado")
	}

	user := entities.User{Email: u.Email}
	passHash, err := user.EncryptPassword(u.Password)
	if err != nil {
		return err
	}
	user.PasswordHash = passHash

	if err := a.repo.User.Add(ctx, &user); err != nil {
		return err
	}

	return nil
}

func (a *authentication) Login(ctx context.Context, data entities.UserLoginView) (entities.UserLoginResponse, error) {
	var resp entities.UserLoginResponse
	if err := data.Validate(); err != nil {
		return resp, err
	}
	user, err := a.repo.User.FindEmail(ctx, data.Email)
	if err != nil {
		return resp, err
	}

	if user.Email == "" {
		return resp, errors.New("usuário não existe")
	}

	if err := user.Validate(data.Password); err != nil {
		return resp, err
	}

	p, err := a.repo.Person.FindID(ctx, user.ID)
	if err != nil {
		return resp, err
	}
	token, err := jwttoken.CreateToken(user.ID, p.UUID)
	if err != nil {
		return resp, err
	}

	resp.Exp = token.ExpiredToken
	resp.Token = token.Token
	return resp, nil
}

func (a *authentication) Logout() {}
