package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/victoorraphael/coordinator/internal/adapters/repository"
	"github.com/victoorraphael/coordinator/internal/domain"
	"net/http"
	"time"
)

var (
	studentRepository = repository.Student{}
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	list, _ := studentRepository.Find(ctx, 0, 0)
	return c.JSON(http.StatusOK, map[string][]domain.Student{"message": list})
}

func (s *StudentHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "get"})
}

func (s *StudentHandler) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "create"})
}

func (s *StudentHandler) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "delete"})
}

func (s *StudentHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "update"})
}
