package entities_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"testing"
)

func TestUser_EncryptPassword(t *testing.T) {
	cases := []struct {
		Name     string
		Password string
		Expected string
	}{
		{
			Name:     "empty password",
			Password: "",
			Expected: "",
		},
		{
			Name:     "password length < 8",
			Password: "short",
			Expected: "",
		},
		{
			Name:     "password valid",
			Password: "validpassword",
			Expected: "somevalidpassword",
		},
	}

	u := entities.User{}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			pass, _ := u.EncryptPassword(c.Password)
			assert.GreaterOrEqual(t, len(pass), len(c.Expected))
		})
	}
}
