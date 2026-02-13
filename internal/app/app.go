package app

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/marisasha/kinolog/internal/handler"
	"github.com/marisasha/kinolog/internal/repository"
	"github.com/marisasha/kinolog/internal/service"
	httpserver "github.com/marisasha/kinolog/internal/transport/http"
)

type App struct {
	server   *httpserver.Server
	handlers *handler.Handler
	db       *sqlx.DB
}

func NewApp(cfg repository.Config) (*App, error) {
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(httpserver.Server)

	return &App{
		server:   server,
		handlers: handlers,
		db:       db,
	}, nil
}

func (a *App) Run(port string) error {
	return a.server.Run(port, a.handlers.InitRoutes())
}

func (a *App) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}
	return a.db.Close()
}
