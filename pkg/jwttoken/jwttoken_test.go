package jwttoken_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/victoorraphael/coordinator/pkg/jwttoken"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := os.Setenv("JWT_SECRET", "topsecret")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestCreateToken(t *testing.T) {
	cases := []struct {
		Name      string
		UserID    int64
		UserUUID  string
		ErrExpect error
	}{
		{
			Name:      "correct data",
			UserID:    1,
			UserUUID:  uuid.New().String(),
			ErrExpect: nil,
		}, {
			Name:      "missing user id",
			UserID:    0,
			UserUUID:  uuid.New().String(),
			ErrExpect: jwttoken.ErrInvalidUserData,
		}, {
			Name:      "missing user uuid",
			UserID:    1,
			UserUUID:  "",
			ErrExpect: jwttoken.ErrInvalidUserData,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			dt, err := jwttoken.CreateToken(c.UserID, c.UserUUID)
			assert.Equal(t, c.ErrExpect, err)
			if c.ErrExpect != nil {
				assert.Nil(t, dt)
			} else {
				assert.NotNil(t, dt)
				assert.NotEmpty(t, dt.Token)
				assert.Greater(t, dt.ExpiredToken, int64(0))
			}
		})
	}
}
