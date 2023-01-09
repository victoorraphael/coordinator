package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/victoorraphael/coordinator/internal/application/services"
	"github.com/victoorraphael/coordinator/internal/domain"
	"log"
	"net/http"
)

type StudentHandler struct{}

func (s *StudentHandler) Routes(e *echo.Echo) {
	std := e.Group("/students")

	std.GET("/", s.List)
	std.GET("/:id/", s.Get)
	std.POST("/", s.Create)
	std.DELETE("/:id/", s.Delete)
	std.PUT("/:id/", s.Update)
}

func (s *StudentHandler) List(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeHandler)
	defer cancel()
	list, _ := services.ListPerson(ctx, &studentRepository)
	return c.JSON(http.StatusOK, list)
}

func (s *StudentHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "get"})
}

func (s *StudentHandler) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeHandler)
	defer cancel()
	student := &domain.Student{}
	if err := c.Bind(student); err != nil {
		return c.String(http.StatusBadRequest, "invalid payload")
	}

	uid, err := services.CreatePerson(ctx, &studentRepository, student)
	if err != nil {
		log.Println("error: ", err.Error())
		return c.String(http.StatusInternalServerError, "internal error")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"uuid": uid})
}

func (s *StudentHandler) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "delete"})
}

func (s *StudentHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "update"})
}
