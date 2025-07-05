package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

func NewPostgresPool(ctx context.Context, cfg config.Database) (*pgxpool.Pool, error) {
	const op = "postgresql connection"
	log := logctx.Logger(ctx)
	log.Info("Connecting to PostgreSQL",
		zap.String("host", cfg.Host),
		zap.String("port", cfg.Port),
		zap.String("database", cfg.DatabaseName),
		zap.String("sslmode", cfg.SSLMode),
	)
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.SSLMode)
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Error("Error parsing PostgreSQL connection string", zap.Error(err))
		return nil, err
	}
	log.Info("Connected to PostgreSQL", zap.String("connectionString", connectionString))
	poolConfig.MaxConns = int32(cfg.MaxOpenConnections)
	poolConfig.MinConns = int32(cfg.MaxIdleConnections)
	poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Error("Error connecting to PostgreSQL", zap.Error(err))
		return nil, err
	}
	log.Info("config sent successfully, end connection func")
	return pool, err
}

func InitSchema(ctx context.Context, pool *pgxpool.Pool) error {
	log := logctx.Logger(ctx)
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Error("Error acquiring connection", zap.Error(err))
		return err
	}
	log.Info("connected to PostgreSQL")
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
		log.Error("Error creating user_table", zap.Error(err))
		return err
	}
	log.Info("created user_table")
	query = `CREATE TABLE IF NOT EXISTS language (
    	code VARCHAR(20) PRIMARY KEY,
    	language_name VARCHAR(255) UNIQUE NOT NULL
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		log.Error("Error creating language_table", zap.Error(err))
		return err
	}
	log.Info("created language_table")
	query = `CREATE TABLE IF NOT EXISTS institute (
    	institute_id SERIAL PRIMARY KEY,
    	name VARCHAR(255) UNIQUE NOT NULL
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		log.Error("Error creating institute_table", zap.Error(err))
		return err
	}
	log.Info("created institute_table")
	query = `CREATE TABLE IF NOT EXISTS lab (
    	lab_id SERIAL PRIMARY KEY,
    	name VARCHAR(255) UNIQUE NOT NULL,
    	institute_id INT NOT NULL,
    	FOREIGN KEY (institute_id) REFERENCES institute(institute_id)
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		log.Error("Error creating lab_table", zap.Error(err))
		return err
	}
	log.Info("created lab_table")
	query = `CREATE TABLE IF NOT EXISTS user_language (
    	user_language_id SERIAL PRIMARY KEY,
    	profile_id INT NOT NULL,
    	language_code VARCHAR(255) NOT NULL,
    	FOREIGN KEY (profile_id) REFERENCES user_profile(profile_id),
    	FOREIGN KEY (language_code) REFERENCES language(language_code)
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		log.Error("Error creating user_language_table", zap.Error(err))
		return err
	}
	log.Info("created user_language_table")
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

		log.Error("Error creating user_institute_table", zap.Error(err))
		return err
	}
	log.Info("created course_table")
	query = `CREATE TABLE IF NOT EXISTS user_course_instance (
		user_course_id SERIAL PRIMARY KEY,
		profile_id INT NOT NULL,
		instance_id INT NOT NULL,
		FOREIGN KEY (profile_id) REFERENCES user_profile(profile_id),
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {

		log.Error("Error creating user_course_table", zap.Error(err))
		return err
	}
	tx, err := pool.Begin(ctx)
	if err != nil {

		log.Error("Error starting transaction", zap.Error(err))
		return err
	}
	log.Info("started transaction")
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `
		INSERT INTO language (language_code, language_name)
		VALUES ('en', 'English'), ('ru', 'Russian')
		ON CONFLICT (language_code) DO NOTHING;
	`)
	if err != nil {
		log.Error("Error adding language", zap.Error(err))
		return fmt.Errorf("failed to insert languages: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {

		log.Error("Error committing transaction", zap.Error(err))
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	//TODO FSRO when will be more obvious what to do
	log.Info("committed transaction")
	return nil
}
