package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golangsugar/chatty"
	"github.com/victoorraphael/coordinator/internal/domain/entities"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/errs"
	"net/http"
)

type addressHandler struct {
	srv *services.Services
}

func RegisterAddressRoutes(srv *services.Services, router *gin.RouterGroup) {
	hdl := &addressHandler{srv}
	addressGroup := router.Group("address")
	addressGroup.
		GET("", hdl.Find).
		GET("/:uuid", hdl.FindUUID).
		POST("", hdl.Create)
}

func (a *addressHandler) Find(c *gin.Context) {
	list, err := a.srv.Address.FetchAll(c)
	if err != nil {
		chatty.Errorf("falha ao buscar endereços: err: %v", err)
		c.String(http.StatusInternalServerError, "não foi possível buscar endereços")
		return
	}
	c.JSON(http.StatusOK, list)
}

func (a *addressHandler) Create(c *gin.Context) {
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

func (a *addressHandler) FindUUID(c *gin.Context) {
	key := c.Param("uuid")
	if key == "" {
		chatty.Error("falha ao buscar endereco: empty uuid")
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrFieldViolation, "uuid nao pode ser vazio").Error())
		return
	}

	addr, err := a.srv.Address.Find(c, entities.Address{UUID: key})
	if err != nil {
		chatty.Errorf("falha ao buscar endereco: %v", err)
		c.String(http.StatusBadRequest, errs.WrapError(errs.ErrInternalError, "falha ao buscar endereco").Error())
		return
	}

	if addr.ID == 0 {
		c.String(http.StatusNotFound, errs.WrapError(errs.ErrNotFound, "endereco nao existe").Error())
		return
	}

	c.JSON(http.StatusOK, addr)
}
