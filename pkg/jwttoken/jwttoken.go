package jwttoken

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

type Details struct {
	Token        string
	ExpiredToken int64
}

type Claims struct {
	UserID     int64
	UserUUID   string
	Authorized bool
}

var (
	ErrInvalidUserData = errors.New("invalid user data")
)

func CreateToken(userID int64, userUUID string) (*Details, error) {
	if userID == 0 || userUUID == "" {
		return nil, ErrInvalidUserData
	}

	td := &Details{}
	td.ExpiredToken = time.Now().Add(time.Minute * 30).Unix()
	var err error
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["user_uuid"] = userUUID
	claims["exp"] = td.ExpiredToken

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	td.Token, err = at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong signature method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return token, err
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func ExtractTokenMetadata(r *http.Request) (*Claims, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authorized, ok := claims["authorized"].(bool)
		if !ok {
			return nil, err
		}

		userID := int64(claims["user_id"].(float64))
		userUUIDStr := claims["user_uuid"].(string)

		return &Claims{
			UserID:     userID,
			UserUUID:   userUUIDStr,
			Authorized: authorized,
		}, nil
	}

	return nil, errors.New("invalid token")
}
