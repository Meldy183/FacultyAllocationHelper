package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-chi/chi/v5"
	app "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/application/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
	auth "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/domain/auth/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/cookies"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/database/postgres"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/http/handlers"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/http/middleware"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/jwt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	config, err := config.New()
	fmt.Printf("%+v\n", config)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	repo := postgres.NewRepo(ctx, config.Database)
	defer repo.CloseConn()
	jwtService := jwt.NewJWTService(*repo, config.JWT)
	cookiesService := cookies.NewCookiesService(config.Cookies)
	middlewareService := middleware.NewMiddlewareService(jwtService, cookiesService)
	domainService := auth.NewService(*repo)
	authService := app.NewAuthService(domainService, jwtService, cookiesService, *config)
	handlers := handlers.NewHandlers(authService)
	r := chi.NewRouter()
	r.Use(middlewareService.AuthMiddleware) // Example roles, adjust as needed
	r.Post("/auth/login", handlers.Login)
	r.Post("/auth/logout", handlers.Logout)
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/refresh", handlers.RefreshToken)
	srv := &http.Server{
		Addr:         config.HTTPserver.Host + ":" + strconv.Itoa(config.HTTPserver.Port),
		Handler:      r,
		ReadTimeout:  config.HTTPserver.ReadTimeout,
		WriteTimeout: config.HTTPserver.WriteTimeout,
		IdleTimeout:  config.HTTPserver.IdleTimeout,
	}
	log.Printf("RUNNIN'")
	log.Printf(srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v", err)
	}
}
