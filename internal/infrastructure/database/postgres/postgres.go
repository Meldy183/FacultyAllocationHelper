package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
)

func New(context context.Context, cfg config.Database) (*pgxpool.Pool, error) {
	log.Printf("Connecting to PostgreSQL database at %s:%s", cfg.User, cfg.Password)
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}
	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)
	poolConfig.MaxConnIdleTime = cfg.ConnLifetime
	poolConfig.ConnConfig.ConnectTimeout = cfg.ConnTimeOut
	pool, err := pgxpool.NewWithConfig(context, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}
	defer pool.Close()
	log.Printf("HORRAY WE CONNECTED")
	return pool, nil
}

func InitTables(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255)PRIMARY KEY,
		email VARCHAR(50) UNIQUE NOT NULL,
		password_hash BYTEA NOT NULL,
		role_id INT NOT NULL,
		Foreign Key (role_id) REFERENCES roles(id),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`
	conn.Exec(ctx, query)
	query = `
	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		action VARCHAR(50) NOT NULL,
		timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		Foreign Key (user_id) REFERENCES users(id)
		);`
	conn.Exec(ctx, query)
	query = `
	CREATE TABLE IF NOT EXISTS roles (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) UNIQUE NOT NULL,
		);`
	conn.Exec(ctx, query)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	roleCreatQuery := `INSERT INTO roles (name) VALUES ($1) ON CONFLICT (name) DO NOTHING`
	roles := []string{"Super_Admin", "Institute_Repr", "Support", "Education_Dept", "Faculty"}
	for _, role := range roles {
		_, err := tx.Exec(ctx, roleCreatQuery, role)
		if err != nil {
			return fmt.Errorf("failed to insert role %s: %w", role, err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
