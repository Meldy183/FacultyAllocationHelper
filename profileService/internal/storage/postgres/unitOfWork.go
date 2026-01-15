package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/transaction"
	"go.uber.org/zap"
)

var _ transaction.UnitOfWork = (*PostgresUnitOfWork)(nil)
var _ transaction.Transaction = (*PostgresTransaction)(nil)

type PostgresUnitOfWork struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewPostgresUnitOfWork(pool *pgxpool.Pool, logger *zap.Logger) *PostgresUnitOfWork {
	return &PostgresUnitOfWork{pool: pool, logger: logger}
}

func (u *PostgresUnitOfWork) Do(ctx context.Context, fn func(tx transaction.Transaction) error) error {
	transact, err := u.pool.Begin(ctx)
	if err != nil {
		return err
	}

	tx := &PostgresTransaction{tx: transact}
	err = fn(tx)
	if err != nil {
		rbErr := tx.Rollback(ctx)
		if rbErr != nil {
			u.logger.Error("failed to rollback transaction",
				zap.String("layer", "unitOfWork"),
				zap.Error(rbErr),
			)
		}
		return err
	}

	return tx.Commit(ctx)
}

type PostgresTransaction struct {
	tx pgx.Tx
}

func (t *PostgresTransaction) Commit(ctx context.Context) error   { return t.tx.Commit(ctx) }
func (t *PostgresTransaction) Rollback(ctx context.Context) error { return t.tx.Rollback(ctx) }
