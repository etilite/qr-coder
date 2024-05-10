package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/etilite/qr-coder/internal/app"
	httpserver "github.com/etilite/qr-coder/internal/delivery/http"
)

const (
	shutdownTime time.Duration = 10 * time.Second
)

func main() {
	cfg := app.NewConfigFromEnv()

	a := app.New(cfg)

	server := httpserver.NewServer(cfg.HTTPAddr)

	if err := a.Run(server); err != nil {
		slog.Error("unable to start app", "error", err)
		panic(err)
	}

	shutdown := make(chan bool)
	gracefulStop(shutdown)
	<-shutdown
	slog.Info("signal caught, shutting down")

	go aggressiveStop()

	if err := a.Stop(); err != nil {
		slog.Error("unable to stop app", "error", err)
		panic(err)
	}
}

func gracefulStop(shutdown chan<- bool) {
	c := make(chan os.Signal)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-c
		shutdown <- true
	}()
}

func aggressiveStop() {
	ticker := time.NewTicker(shutdownTime)
	<-ticker.C

	slog.Warn("application is aggressive stopped")
	os.Exit(0)
}
