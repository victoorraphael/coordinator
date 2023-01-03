package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/store"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/victoorraphael/coordinator/cmd/http/handlers"
	_ "github.com/victoorraphael/coordinator/internal/adapters"
	"github.com/victoorraphael/coordinator/internal/connect"
	"github.com/victoorraphael/coordinator/internal/service"
)

type Status struct {
	System   bool
	Database bool
}

func main() {
	PORT := os.Getenv("PORT")

	adapters, _ := connect.Connect()
	defer func() {
		log.Println("closing connection with database ⌛")
		_ = adapters.DB.GetDatabase().Close()
	}()

	e := echo.New()
	s := store.New(adapters)
	services := service.New(s)
	handlers.StudentRoutes(e, services.StudentSRV())

	//e.Use(middleware.Logger())
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
