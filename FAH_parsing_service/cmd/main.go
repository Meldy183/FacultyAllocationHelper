package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/infrastructure/http/handlers"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/infrastructure/http/middleware"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	config, err := config.New()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	ctx, err = logger.NewLogger(config.Logger, ctx)
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}
	l := logger.GetFromContext(ctx)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://87.228.102.156"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow all headers
		AllowCredentials: true,          // Allow credentials
	})
	r := chi.NewRouter()
	r.Use(middleware.LoggerMiddleware(l))
	r.Post("parsing/parse", handlers.Parse)
	handler := c.Handler(r)
	srv := &http.Server{
		Addr:         config.HTTPserver.Host + ":" + strconv.Itoa(config.HTTPserver.Port),
		Handler:      handler,
		ReadTimeout:  config.HTTPserver.ReadTimeout,
		WriteTimeout: config.HTTPserver.WriteTimeout,
		IdleTimeout:  config.HTTPserver.IdleTimeout,
	}
	l.Info(ctx, "Parsing service running", zap.String("Adress", srv.Addr))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal(ctx, "Error starting server", zap.Error(err))
		}
	}()
	<-ctx.Done()
	l.Info(ctx, "Shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		l.Error(ctx, "Server forced to shutdown", zap.Error(err))
	}

}
