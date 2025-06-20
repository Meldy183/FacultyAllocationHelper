package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	app "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/application/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
	auth "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/domain/auth/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/cookies"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/database/postgres"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/http/handlers"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/jwt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config, err := config.New()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	repo := postgres.NewRepo(ctx, config)
	defer repo.CloseConn()
	jwtService := jwt.NewJWTService(*repo, config.JWT)
	cookiesService := cookies.NewCookiesService(config.Cookies)
	domainService := auth.NewService(*repo)
	authService := app.NewAuthService(domainService, jwtService, cookiesService, *config)
	handlers := handlers.NewHandlers(authService)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // Allow all headers
		AllowCredentials: true,          // Allow credentials
	})
	r := chi.NewRouter() // Example roles, adjust as needed
	r.Post("/auth/login", handlers.Login)
	r.Post("/auth/logout", handlers.Logout)
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/refresh", handlers.RefreshToken)
	handler := c.Handler(r)
	srv := &http.Server{
		Addr:         config.HTTPserver.Host + ":" + strconv.Itoa(config.HTTPserver.Port),
		Handler:      handler,
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
