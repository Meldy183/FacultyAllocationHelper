package app

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/courses"
	userprofile2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/filters"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/parse"
	router "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/http"
	"go.uber.org/zap"
)

type App struct {
	facultyHandler *userprofile2.Handler
	courseHandler  *courses.Handler
	filtersHandler *filters.Handler
	parsingHandler *parse.Handler

	router chi.Router
	server *http.Server

	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewApp(
	pool *pgxpool.Pool,
	logger *zap.Logger,
	facultyHandler *userprofile2.Handler,
	courseHandler *courses.Handler,
	filtersHandler *filters.Handler,
	parsingHandler *parse.Handler,
) *App {
	router := router.NewRouter(
		facultyHandler,
		courseHandler,
		filtersHandler,
		parsingHandler,
	)
	return &App{
		pool:           pool,
		logger:         logger,
		router:         router,
		facultyHandler: facultyHandler,
		courseHandler:  courseHandler,
		filtersHandler: filtersHandler,
		parsingHandler: parsingHandler,
	}
}
func (a *App) Run(ctx context.Context,
	host string, port string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
	idleTimeout time.Duration) error {
	a.server = &http.Server{
		Addr:         host + ":" + port,
		Handler:      a.router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
	a.logger.Info("Started Server",
		zap.String("address", a.server.Addr),
	)
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
func (a *App) Close() error {
	a.logger.Info("Closing application resources...")
	if a.pool != nil {
		a.pool.Close()
	}

	return nil
}
