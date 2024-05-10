package app

import (
	"log/slog"
)

type service interface {
	Run() error
	Stop() error
}

type App struct {
	logger   *slog.Logger
	services []service
}

func New(cfg Config) *App {
	return &App{
		logger: slog.Default(),
	}
}

func (app *App) Run(services ...service) error {
	for _, srv := range services {
		if err := srv.Run(); err != nil {
			return err
		}
		app.services = append(app.services, srv)
	}
	app.logger.Info("app started")

	return nil
}

func (app *App) Stop() error {
	for _, srv := range app.services {
		if err := srv.Stop(); err != nil {
			return err
		}
	}
	app.logger.Info("app stopped")

	return nil
}
