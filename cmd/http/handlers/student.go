package handlers

import (
	"github.com/victoorraphael/school-plus-BE/internal/student"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/victoorraphael/school-plus-BE/internal/contracts"
)

func StudentRoutes(e *echo.Echo, service contracts.IService[student.Student]) {
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

func StudentHandlerGetList(c echo.Context, s contracts.IService[student.Student]) error {
	stds, err := s.List()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, stds)
}

func StudentHandlerGet(c echo.Context, s contracts.IService[student.Student]) error {
	id := c.Param("id")
	uid := uuid.MustParse(id)
	studentQuery := student.New()
	studentQuery.UUID = uid
	std, err := s.Get(studentQuery)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, std)
}

func StudentHandlerPost(c echo.Context, s contracts.IService[student.Student]) error {
	var std student.Student
	if err := c.Bind(&std); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	data, err := s.Add(std)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
