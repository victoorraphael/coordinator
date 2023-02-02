package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"log"
	"net/http"
)

type AddressHandler struct {
	addr services.IAddressService
}

func NewAddressHandler(s *services.Services) *AddressHandler {
	return &AddressHandler{s.Address}
}

func (a *AddressHandler) Find(c *gin.Context) {
	list, err := a.addr.FetchAll(c)
	if err != nil {
		log.Println("falha ao buscar endereços: err:", err)
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

	err := a.addr.Create(c, &addr)
	if err != nil {
		c.String(http.StatusInternalServerError, "não foi possível criar o endereço")
		return
	}

	c.JSON(http.StatusOK, addr)
}
