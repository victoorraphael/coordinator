package handlers

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/entities"
	"github.com/victoorraphael/coordinator/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func StudentRoutes(e *echo.Echo, service service.IStudentSRV) {
	std := e.Group("/students")

	std.Add("GET", "/", func(c echo.Context) error { return StudentHandlerGetList(c, service) })
	std.Add("GET", "/:id/", func(c echo.Context) error { return StudentHandlerGet(c, service) })
	std.Add("POST", "/", func(c echo.Context) error { return StudentHandlerPost(c, service) })
}

func StudentHandlerGetList(c echo.Context, s service.IStudentSRV) error {
	stds, err := s.List(context.Background())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, stds)
}

func StudentHandlerGet(c echo.Context, s service.IStudentSRV) error {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "id should not be empty!")
	}
	log.Println("info: trying to find student with uuid:", id)
	uid := uuid.MustParse(id)
	studentQuery := entities.NewStudent()
	studentQuery.UUID = uid
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	std, err := s.Get(ctx, studentQuery)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, std)
}

func StudentHandlerPost(c echo.Context, s service.IStudentSRV) error {
	var std entities.Student
	if err := c.Bind(&std); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	data, err := s.Add(ctx, std)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
