package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/victoorraphael/school-plus-BE/internal/repositories/student/mongo"
	"net/http"
	"os"
)

type Status struct {
	System  bool
	MongoDB bool
}

func main() {
	PORT := os.Getenv("PORT")
	MONGO := os.Getenv("MONGO_URI")

	repo, _ := mongo.New(MONGO)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error {
		dbStatus := repo.Ping()
		res := Status{
			System:  true,
			MongoDB: dbStatus,
		}
		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
}
