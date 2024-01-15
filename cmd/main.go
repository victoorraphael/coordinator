package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golangsugar/chatty"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httphdl "github.com/victoorraphael/coordinator/cmd/http"
	"github.com/victoorraphael/coordinator/internal/domain/repository"
	"github.com/victoorraphael/coordinator/internal/domain/services"
	"github.com/victoorraphael/coordinator/pkg/database"
	"github.com/victoorraphael/coordinator/pkg/fixtures"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	debugMode bool
	seed      bool
)

func init() {
	_ = godotenv.Load(".env")
	flag.BoolVar(&debugMode, "debug", false, "run routes without authorization")
	flag.BoolVar(&seed, "seed", false, "seed database with dumb data")
	flag.Parse()

	chatty.SetSeverityLevelDebug()
	chatty.SetGlobalOutputFormatPlainText()

	if seed {
		chatty.Debug("seed database...")
		err := fixtures.
			Connect().
			Seed().
			Close()
		if err != nil {
			chatty.FatalErr(err)
		}
	}
}

func main() {
	dbPool, err := database.NewPostgres(5)
	if err != nil {
		chatty.FatalErr(err)
	}
	defer dbPool.Close()

	conn, err := dbPool.Acquire()
	if err != nil {
		chatty.FatalErr(err)
	}
	dbPool.Release(conn)

	chatty.Info("health: trying to ping database")
	if err := conn.Ping(); err != nil {
		chatty.FatalErr(err)
	}

	repo := repository.New(dbPool)
	s := services.New(repo)

	// setup server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%v", os.Getenv("PORT")),
		Handler:           httphdl.Routes(s, debugMode),
		WriteTimeout:      time.Second * 15,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		IdleTimeout:       time.Second * 60,
	}

	// start server and wait for OS signal
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			chatty.FatalErr(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	_ = srv.Shutdown(ctx)
	chatty.Info("shutting down! 👋")
}
