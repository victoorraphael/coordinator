package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/victoorraphael/coordinator/internal/services"
	"log"
	"net/http"
)

type AddressHandler struct{}

func (a *AddressHandler) Routes(e *echo.Echo, services *services.Services) {
	addr := e.Group("/address")

	addr.GET("/", a.Find(services))
	//addr.GET("/:id", a.Find(services))
	addr.POST("/", a.Create(services))
}

func (a *AddressHandler) Find(srv *services.Services) func(c echo.Context) error {
	return func(c echo.Context) error {
		list, err := srv.Address.FetchAll()
		if err != nil {
			log.Println("falha ao buscar endereços: err:", err)
			return c.String(http.StatusInternalServerError, "não foi possível buscar endereços")
		}
		return c.JSON(http.StatusOK, list)
	}
}

func (a *AddressHandler) Create(srv *services.Services) func(ctx echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "not implemented")
	}
}
