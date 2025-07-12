package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"go.uber.org/zap"
)

var _ language.Repository = (*LanguageRepo)(nil)

type LanguageRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewLanguageRepo(pool *pgxpool.Pool, logger *zap.Logger) *LanguageRepo {
	return &LanguageRepo{pool: pool, logger: logger}
}

const (
	queryGetLangByCode = `SELECT code, language_name FROM language WHERE code = $1`
	queryGetAllLang    = `SELECT code, language_name FROM language`
	logGetAll          = "GetAll"
	logGetByCode       = "GetByCode"
)

func (r *LanguageRepo) GetAll(ctx context.Context) ([]*language.Language, error) {
	rows, err := r.pool.Query(ctx, queryGetAllLang)
	if err != nil {
		r.logger.Error("failed to query all languages",
			zap.String("layer", logLayer),
			zap.String("function", logGetAll),
			zap.Error(err),
		)
		return nil, fmt.Errorf("get all languages: %w", err)
	}
	defer rows.Close()
	var languages []*language.Language
	for rows.Next() {
		if rows.Err() != nil {
			r.logger.Error("failed to query all languages",
				zap.String("layer", logLayer),
				zap.String("function", logGetAll),
				zap.Error(rows.Err()),
			)
			return nil, fmt.Errorf("get all languages: %w", rows.Err())
		}
		var lang language.Language
		err := rows.Scan(
			&lang.LanguageCode,
			&lang.LanguageName)
		if err != nil {
			r.logger.Error("failed to query all languages",
				zap.String("layer", logLayer),
				zap.String("function", logGetAll),
				zap.Error(err),
			)
			return nil, fmt.Errorf("get all languages: %w", err)
		}
		languages = append(languages, &lang)
	}
	r.logger.Info("all languages returned",
		zap.String("layer", logLayer),
		zap.String("function", logGetAll),
		zap.Int("count", len(languages)),
	)
	return languages, nil
}
func (r *LanguageRepo) GetByCode(ctx context.Context, code string) (*language.Language, error) {
	row := r.pool.QueryRow(ctx, queryGetLangByCode, code)
	var lang language.Language
	err := row.Scan(&lang.LanguageCode, &lang.LanguageName)
	if err != nil {
		r.logger.Error("GetLanguageByCode",
			zap.String("layer", logLayer),
			zap.String("function", logGetByCode),
			zap.String("code", code),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetLanguageByCode: %w", err)
	}
	r.logger.Info("GetLanguageByCode",
		zap.String("layer", logLayer),
		zap.String("function", logGetByCode),
		zap.String("code", code),
	)
	return &lang, nil
}
