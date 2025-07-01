package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
)

func NewPostgresPool(ctx context.Context, cfg config.Database) (*pgxpool.Pool, error) {
	const op = "postgresql connection"
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.SSLMode)
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = int32(cfg.MaxOpenConnections)
	poolConfig.MinConns = int32(cfg.MaxIdleConnections)
	poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}
	return pool, err
}

func InitSchema(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	query := `CREATE TABLE IF NOT EXISTS user_profile (
    	profile_id SERIAL PRIMARY KEY,
    	user_id VARCHAR(255) UNIQUE NOT NULL,
    	email VARCHAR(50) UNIQUE NOT NULL,
    	position VARCHAR(255) NOT NULL,
    	english_name VARCHAR(255) UNIQUE NOT NULL,
    	russian_name VARCHAR(255) UNIQUE NOT NULL,
    	telegram_alias VARCHAR(255) UNIQUE NOT NULL,
    	employment_type VARCHAR(255) UNIQUE,
    	degree BOOL NOT NULL,
    	mode VARCHAR(255) NOT NULL,
    	start_date DATE,
    	end_date DATE,
    	maxload INTEGER
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS language (
    	language_code VARCHAR(20) PRIMARY KEY,
    	language_name VARCHAR(255) UNIQUE NOT NULL
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS institute (
    	institute_id SERIAL PRIMARY KEY,
    	name VARCHAR(255) UNIQUE NOT NULL
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS lab (
    	lab_id SERIAL PRIMARY KEY,
    	name VARCHAR(255) UNIQUE NOT NULL,
    	institute_id INT NOT NULL,
    	FOREIGN KEY (institute_id) REFERENCES institute(institute_id)
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS user_language (
    	user_language_id SERIAL PRIMARY KEY,
    	profile_id INT NOT NULL,
    	language_code VARCHAR(255) NOT NULL,
    	FOREIGN KEY (profile_id) REFERENCES user_profile(profile_id),
    	FOREIGN KEY (language_code) REFERENCES language(language_code)
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS user_institute (
    	user_institute_id SERIAL PRIMARY KEY,
    	profile_id INT NOT NULL,
    	institute_id INT NOT NULL,
    	is_repr BOOL NOT NULL,
    	FOREIGN KEY (profile_id) REFERENCES user_profile(profile_id),
    	FOREIGN KEY (institute_id) REFERENCES institute(institute_id)
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `CREATE TABLE IF NOT EXISTS user_course (
		user_course_id SERIAL PRIMARY KEY,
		profile_id INT NOT NULL,
		instance_id INT NOT NULL,
		FOREIGN KEY (profile_id) REFERENCES user_profile(profile_id),
		FOREIGN KEY (instance_id) REFERENCES course_instance(instance_id)
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		return err
	}
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `
		INSERT INTO language (language_code, language_name)
		VALUES ('en', 'English'), ('ru', 'Russian')
		ON CONFLICT (language_code) DO NOTHING;
	`)
	if err != nil {
		return fmt.Errorf("failed to insert languages: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	//TODO FSRO when will be more obvious what to do
	return nil
}
