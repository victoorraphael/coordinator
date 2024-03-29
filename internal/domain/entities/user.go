package entities

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/pkg/security"
)

type User struct {
	ID           int64  `db:"id"`
	PersonID     int64  `db:"person_id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password"`
}

type UserLoginView struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

func (u *User) EncryptPassword(password string) (string, error) {
	hash, err := security.HashPassword(password)
	return hash, err
}

func (u *User) Validate(password string) error {
	if u.Email == "" {
		return errors.New("email inválido")
	}
	err := checkmail.ValidateFormat(u.Email)
	if err != nil {
		chatty.Info("[user:Validate]: email inválido")
		return errors.New("credenciais inválidas")
	}

	ok := security.PasswordValid(u.PasswordHash, password)
	if !ok {
		chatty.Info("[user:Validate]: password inválido")
		return errors.New("credenciais inválidas")
	}

	return nil
}

func (u *UserLoginView) Validate() error {
	if u.Password == "" {
		return errors.New("password inválido")
	}

	if len(u.Password) < 8 {
		return errors.New("password deve ser pelo menos 8 caracteres")
	}

	if u.Email == "" {
		return errors.New("email inválido")
	}
	err := checkmail.ValidateFormat(u.Email)
	if err != nil {
		chatty.Info("email inválido")
		return errors.New("email inválido")
	}

	return nil
}
