package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logger"
	"go.uber.org/zap"
)

func NewPostgresPool(ctx context.Context, cfg config.Database) (*pgxpool.Pool, error) {
	l := logger.GetLoggerFromCtx(ctx)
	l.Info("Connecting to database")
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode)
	poolcfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		l.Error("Error parsing connection string", zap.Error(err))
		return nil, err
	}
	poolcfg.MaxConns = int32(cfg.MaxOpenConns)
	poolcfg.MinConns = int32(cfg.MaxIdleConns)
	poolcfg.MaxConnLifetime = cfg.ConnMaxLifetime
	pool, err := pgxpool.NewWithConfig(ctx, poolcfg)
	if err != nil {
		l.Error("Error connecting to PostgreSQL", zap.Error(err))
		return nil, err
	}
	l.Info("config sent successfully, end connection func")
	return pool, err
}
