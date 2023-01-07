package handlers

import (
	"context"
	"github.com/victoorraphael/coordinator/internal/student"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func StudentRoutes(e *echo.Echo, service student.Service) {
	std := e.Group("/students")

	std.Add("GET", "/", func(c echo.Context) error { return StudentHandlerGetList(c, service) })
	std.Add("GET", "/:id/", func(c echo.Context) error { return StudentHandlerGet(c, service) })
	std.Add("POST", "/", func(c echo.Context) error { return StudentHandlerPost(c, service) })
	std.Add("DELETE", "/:id/", func(c echo.Context) error { return StudentHandlerDelete(c, service) })
	std.Add("PUT", "/:id/", func(c echo.Context) error { return StudentHandlerUpdate(c, service) })
}

func StudentHandlerGetList(c echo.Context, s student.Service) error {
	//stds, err := s.List(context.Background())
	//if err != nil {
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}

	return c.String(http.StatusOK, "not implemented")
}

func StudentHandlerGet(c echo.Context, s student.Service) error {
	//id := c.Param("id")
	//if id == "" {
	//	return c.String(http.StatusBadRequest, "id should not be empty!")
	//}
	//log.Println("info: trying to find student with uuid:", id)
	//uid := uuid.MustParse(id)
	//studentQuery := student.NewStudent()
	//studentQuery.UUID = uid
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//std, err := s.Get(ctx, studentQuery)
	//if err != nil {
	//	return c.String(http.StatusBadRequest, err.Error())
	//}

	return c.String(http.StatusOK, "not implemented")
}

func StudentHandlerPost(c echo.Context, s student.Service) error {
	var std student.Student
	if err := c.Bind(&std); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	data, err := s.Create(ctx, std)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func StudentHandlerDelete(c echo.Context, s student.Service) error {
	//id := c.Param("id")
	//if id == "" {
	//	return c.String(http.StatusBadRequest, "id should not be empty!")
	//}
	//uid := uuid.MustParse(id)
	//studentQuery := student.NewStudent()
	//studentQuery.UUID = uid
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//
	//err := s.Delete(ctx, studentQuery)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return c.String(http.StatusBadRequest, fmt.Sprintf("student with id: %s does not exists", uid))
	//	}
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}
	return c.String(http.StatusOK, "not implemented")
}

func StudentHandlerUpdate(c echo.Context, s student.Service) error {
	//id := c.Param("id")
	//if id == "" {
	//	return c.String(http.StatusBadRequest, "id should not be empty!")
	//}
	//
	//var std student.Student
	//if err := c.Bind(&std); err != nil {
	//	return c.String(http.StatusBadRequest, err.Error())
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//
	//err := s.Update(ctx, std)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return c.String(http.StatusBadRequest, fmt.Sprintf("student with id: %s does not exists", id))
	//	}
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}

	return c.String(http.StatusOK, "not implemented")
}
