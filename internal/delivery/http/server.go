package http

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
}

func NewServer(addr string) *Server {
	return &Server{
		http: &http.Server{
			Addr:    addr,
			Handler: newMux(),
		},
	}
}

func (s *Server) Run() error {
	slog.Info("http-server: starting web server", "address", s.http.Addr)

	go func() {
		if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("http-server: failed to serve", "error", err)
		}

		slog.Info("http-server: stopped gracefully")
	}()

	return nil
}

func (s *Server) Stop() error {
	shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
	defer done()

	slog.Info("http-server: shutting down")
	if err := s.http.Shutdown(shutdownCtx); err != nil {
		slog.Error("http-server: shutdown failed", "error", err)
		return err
	}

	return nil
}
