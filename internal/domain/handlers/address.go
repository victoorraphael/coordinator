package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"net/http"
)

type AddressHandler struct {
	srv *services.Services
}

func NewAddressHandler(s *services.Services) *AddressHandler {
	return &AddressHandler{s}
}

func (a *AddressHandler) Find(c *gin.Context) {
	list, err := a.srv.Address.FetchAll(c)
	if err != nil {
		chatty.Errorf("falha ao buscar endereços: err: %v", err)
		c.String(http.StatusInternalServerError, "não foi possível buscar endereços")
		return
	}
	c.JSON(http.StatusOK, list)
}

func (a *AddressHandler) Create(c *gin.Context) {
	var addr entities.Address
	if err := c.Bind(&addr); err != nil {
		c.String(http.StatusBadRequest, "campos inválidos")
		return
	}

	err := a.srv.Address.Create(c, &addr)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Errorf("não foi possível criar o endereço: %w", err).Error())
		return
	}

	c.JSON(http.StatusOK, addr)
}
