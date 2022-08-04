package handlers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/labstack/echo/v4"
	stdRepo "github.com/victoorraphael/school-plus-BE/internal/repositories/student"
	"github.com/victoorraphael/school-plus-BE/services/student"
)

func StudentRoutes(e *echo.Echo, service student.Service) {
	//studentRoutes := e.Router()
	e.Add("GET", "/student/:id", func(c echo.Context) error {
		return StudentHandlerGet(c, service)
	})
	e.Add("POST", "/student", func(c echo.Context) error {
		return StudentHandlerPost(c, service)
	})
}

func StudentHandlerGet(c echo.Context, s student.Service) error {
	id := c.Param("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	std, err := s.Get(uid)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, std)
}

func StudentHandlerPost(c echo.Context, s student.Service) error {
	var std stdRepo.Student
	if err := c.Bind(&std); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	data, err := s.Add(std)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data["message"] = "successfully created"
	return c.JSON(http.StatusOK, data)
}
