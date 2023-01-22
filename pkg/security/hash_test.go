package security_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/victoorraphael/coordinator/pkg/security"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHashPassword(t *testing.T) {
	cases := []struct {
		Name     string
		Password string
		Expected string
	}{
		{Name: "empty password", Password: "", Expected: ""},
		{Name: "invalid length", Password: "super", Expected: ""},
		{Name: "valid password", Password: "topsecret", Expected: "somelength"},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res, _ := security.HashPassword(test.Password)
			assert.GreaterOrEqual(t, len(res), len(test.Expected))
		})
	}
}

func TestPasswordValid(t *testing.T) {
	password := "supersecretpassword"
	right, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	wrongHash, _ := bcrypt.GenerateFromPassword([]byte("wronghash"), bcrypt.DefaultCost)
	cases := []struct {
		Name     string
		Password string
		Hash     string
		Match    bool
	}{
		{
			Name:     "wrong password",
			Password: password,
			Hash:     string(wrongHash),
			Match:    false,
		}, {
			Name:     "right password",
			Password: password,
			Hash:     string(right),
			Match:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ok := security.PasswordValid(c.Hash, c.Password)
			assert.Equal(t, c.Match, ok)
		})
	}
}
