package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/victoorraphael/coordinator/internal/adapters/handlers"
	"github.com/victoorraphael/coordinator/internal/adapters/postgres"
	"github.com/victoorraphael/coordinator/internal/adapters/repository"
	"github.com/victoorraphael/coordinator/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Status struct {
	System   bool
	Database bool
}

func main() {
	dbPool, err := postgres.NewAdapter(5)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	repo := repository.New(dbPool)
	s := services.New(repo)

	{
		//connect handlers and register routes
		hand := handlers.NewHandlerAdapter()
		hand.Connect(e, s)
	}

	PORT := os.Getenv("PORT")

	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ping", func(c echo.Context) error {
		conn, _ := dbPool.Acquire()
		defer dbPool.Release(conn)

		res := Status{
			System:   true,
			Database: conn.Ping() == nil,
		}
		return c.JSON(http.StatusOK, res)
	})

	// start server
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

	{
		//print all registered routes
		data, err := json.MarshalIndent(e.Routes(), "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println("routes registered...")
		fmt.Printf("%s\n", data)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_ = srv.Shutdown(ctx)
	log.Println("shutting down! 👋")

	os.Exit(0)
}
