package main

import (
	"context"
	"fmt"
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
	dbPool, err := database.NewPostgres(5)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	repo := repository.New(dbPool)
	s := services.New(repo)

	// setup server
	srv := &http.Server{}
	srv.Addr = fmt.Sprintf(":%v", os.Getenv("PORT"))
	srv.Handler = httphdl.Routes(s)
	srv.WriteTimeout = time.Second * 15
	srv.ReadTimeout = time.Second * 15
	srv.ReadHeaderTimeout = time.Second * 15
	srv.IdleTimeout = time.Second * 60

	// start server and wait for os signal
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
