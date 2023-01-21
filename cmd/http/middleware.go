package main

import (
	"github.com/gin-gonic/gin"
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
