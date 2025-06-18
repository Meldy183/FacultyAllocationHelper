package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
)

func New(context context.Context, cfg config.Database) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Failed to parse database config: %v", err)
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}
	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)
	poolConfig.MaxConnIdleTime = cfg.ConnLifetime
	poolConfig.ConnConfig.ConnectTimeout = cfg.ConnTimeOut
	pool, err := pgxpool.NewWithConfig(context, poolConfig)
	if err != nil {
		log.Fatalf("Failed to create pool: %v", err)
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}
	log.Printf("HORRAY WE CONNECTED")
	return pool, nil
}

func InitTables(ctx context.Context, pool *pgxpool.Pool) error {
	log.Println("are you alive?")
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Printf("Failed to acquire connection: %v", err)
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()
	query := `
	CREATE TABLE IF NOT EXISTS roles (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) UNIQUE NOT NULL
    );`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		log.Printf("you failed in third command: %v", err)
		return err
	}
	query = `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		email VARCHAR(50) UNIQUE NOT NULL,
		password_hash BYTEA NOT NULL,
		role_id INTEGER NOT NULL,
		FOREIGN KEY (role_id) REFERENCES roles(id),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		log.Println("you failed in first command")
		return err
	}
	query = `
	CREATE TABLE IF NOT EXISTS logs (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR(255) NOT NULL,
		action VARCHAR(50) NOT NULL,
		timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
		);`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		log.Println("you failed in second command")
		return err
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)
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
