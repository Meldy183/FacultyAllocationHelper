package integration

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
	app "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/application/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
	auth "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/domain/auth/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/cookies"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/database/postgres"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/http/handlers"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/jwt"
)

func setupTestDB(t *testing.T, cfg config.Database, ctx context.Context) *pgxpool.Pool {
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "user"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "actual_database_password_here"
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	db := os.Getenv("POSTGRES_DB")
	if db == "" {
		db = "postgres"
	}
	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, "actual_database_password_here", host, db)
	t.Logf("Connection string: %s", connString)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Failed to parse database config: %v", err)
		t.Fatalf("failed to parse connection string: %v", err)
	}
	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)
	poolConfig.MaxConnIdleTime = cfg.ConnLifetime
	poolConfig.ConnConfig.ConnectTimeout = cfg.ConnTimeOut
	const maxAttempts = 10
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
		if err == nil {
			if err := postgres.InitTables(ctx, pool); err == nil {
				return pool
			}
			pool.Close()
		}
		t.Logf("Attempt %d/%d failed: %v. Retrying in 1s...", attempt, maxAttempts, err)
		time.Sleep(1 * time.Second)
	}
	t.Fatalf("Failed to connect to database after %d attempts", maxAttempts)
	return nil
}
func setupTestServer(t *testing.T, pl *pgxpool.Pool, config config.Config) *httptest.Server {
	repo := postgres.NewRepository(pl)
	jwtService := jwt.NewJWTService(*repo, config.JWT)
	cookiesService := cookies.NewCookiesService(config.Cookies)
	domainService := auth.NewService(*repo)
	authService := app.NewAuthService(domainService, jwtService, cookiesService, config)
	handlers := handlers.NewHandlers(authService)
	r := chi.NewRouter() // Example roles, adjust as needed
	r.Post("/auth/login", handlers.Login)
	r.Post("/auth/logout", handlers.Logout)
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/refresh", handlers.RefreshToken)
	return httptest.NewServer(r)
}
func setupTestConfig() *config.Config {
	testConfigPath := filepath.Join("..", "..", "config", "config.yaml") // Adjusted path
	if _, err := os.Stat(testConfigPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist at path: %s", testConfigPath)
	}
	var cfg config.Config
	if err := cleanenv.ReadConfig(testConfigPath, &cfg); err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	return &cfg
}
