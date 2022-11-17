package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/victoorraphael/school-plus-BE/infra/contracts"
	"github.com/victoorraphael/school-plus-BE/infra/entities"

	"github.com/labstack/echo/v4"
)

func StudentRoutes(e *echo.Echo, service contracts.IService[entities.Student]) {
	std := e.Group("/students")

	std.Add("GET", "/", func(c echo.Context) error {
		return StudentHandlerGetList(c, service)
	})
	std.Add("GET", "/:id/", func(c echo.Context) error {
		return StudentHandlerGet(c, service)
	})
	std.Add("POST", "/", func(c echo.Context) error {
		return StudentHandlerPost(c, service)
	})
}

func StudentHandlerGetList(c echo.Context, s contracts.IService[entities.Student]) error {
	stds, err := s.List()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, stds)
}

func StudentHandlerGet(c echo.Context, s contracts.IService[entities.Student]) error {
	id := c.Param("id")
	uid := uuid.MustParse(id)
	std, err := s.Get(entities.Student{ID: uid})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, std)
}

func StudentHandlerPost(c echo.Context, s contracts.IService[entities.Student]) error {
	var std entities.Student
	if err := c.Bind(&std); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	data, err := s.Add(std)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
