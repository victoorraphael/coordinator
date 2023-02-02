package httphdl

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/pkg/jwttoken"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwttoken.TokenValid(c.Request)
		if err != nil {
			c.String(http.StatusForbidden, "unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := jwttoken.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.String(http.StatusForbidden, "unauthorized")
			c.Abort()
			return
		}
		c.Set("session", entities.Session{
			UserID:   data.UserID,
			UserUUID: data.UserUUID,
		})

		c.Next()
	}
}
