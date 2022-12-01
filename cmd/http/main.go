package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/victoorraphael/school-plus-BE/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/victoorraphael/school-plus-BE/cmd/http/handlers"
	_ "github.com/victoorraphael/school-plus-BE/internal/adapters"
	"github.com/victoorraphael/school-plus-BE/internal/connect"
)

type Status struct {
	System   bool
	Database bool
}

func main() {
	PORT := os.Getenv("PORT")

	adapters, _ := connect.Connect()

	e := echo.New()
	services := service.New(adapters)
	handlers.StudentRoutes(e, services.StudentSRV())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error {
		dbStatus := adapters.DB.Ping()
		res := Status{
			System:   true,
			Database: dbStatus,
		}
		return c.JSON(http.StatusOK, res)
	})

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}

	srv := e.Server
	srv.Addr = fmt.Sprintf(":%v", PORT)
	srv.WriteTimeout = time.Second * 15
	srv.ReadTimeout = time.Second * 15
	srv.IdleTimeout = time.Second * 60

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("routes registered...")
	fmt.Printf("%s\n", data)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_ = srv.Shutdown(ctx)
	log.Println("shutting down! 👋🏼")

	os.Exit(0)
}
