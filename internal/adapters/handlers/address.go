package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/victoorraphael/coordinator/internal/application/services"
	"github.com/victoorraphael/coordinator/internal/domain"
	"net/http"
)

type AddressHandler struct{}

func (a *AddressHandler) Routes(e *echo.Echo) {
	addr := e.Group("/address")

	addr.GET("/:id", a.Find)
	addr.POST("/", a.Create)
}

func (a *AddressHandler) Find(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "not implemented")
}

func (a *AddressHandler) Create(c echo.Context) error {
	var addr domain.Address
	if err := c.Bind(&addr); err != nil {
		return c.String(http.StatusBadRequest, "invalid payload")
	}
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeHandler)
	defer cancel()
	err := services.CreateAddress(ctx, addressRepository, &addr)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal error")
	}

	return c.JSON(http.StatusCreated, addr)
}
