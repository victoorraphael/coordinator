package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	httphdl "github.com/victoorraphael/coordinator/cmd/http"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	debugMode := flag.Bool("debug", false, "run routes without authorization")
	flag.Parse()

	dbPool, err := database.NewPostgres(5)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	conn, err := dbPool.Acquire()
	if err != nil {
		log.Panic(err)
	}
	dbPool.Release(conn)

	log.Println("health: trying to ping database")
	if err := conn.Ping(); err != nil {
		log.Panic(err)
	}

	repo := repository.New(dbPool)
	s := services.New(repo)

	// setup server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler:           httphdl.Routes(s, *debugMode),
		WriteTimeout:      time.Second * 15,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		IdleTimeout:       time.Second * 60,
	}

	// start server and wait for OS signal
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_ = srv.Shutdown(ctx)
	log.Println("shutting down! 👋")
}
