package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"net/http"
)

type AuthHandler struct {
	srv *services.Services
}

func NewAuthHandler(s *services.Services) *AuthHandler {
	return &AuthHandler{s}
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req entities.UserLoginView
	if err := c.ShouldBind(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := a.srv.Auth.Login(c, req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (a *AuthHandler) SignIn(c *gin.Context) {
	var req entities.UserLoginView
	if err := c.ShouldBind(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := a.srv.Auth.SignIn(c, req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

}
