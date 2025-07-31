package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
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
	queryGetLangByCode     = `SELECT code, language_name FROM language WHERE code = $1`
	queryGetAllLang        = `SELECT code FROM language`
	queryGetLangCodeByName = `SELECT code FROM language WHERE language_name = $1`
)

func (r *LanguageRepo) GetAllLanguages(ctx context.Context) ([]string, error) {
	rows, err := r.pool.Query(ctx, queryGetAllLang)
	if err != nil {
		r.logger.Error("failed to query all languages",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllLanguages),
			zap.Error(err),
		)
		return nil, fmt.Errorf("get all languages: %w", err)
	}
	defer rows.Close()
	var languages []string
	for rows.Next() {
		if rows.Err() != nil {
			r.logger.Error("failed to query all languages",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllLanguages),
				zap.Error(rows.Err()),
			)
			return nil, fmt.Errorf("get all languages: %w", rows.Err())
		}
		var lang string
		err := rows.Scan(
			&lang,
		)
		if err != nil {
			r.logger.Error("failed to query all languages",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllLanguages),
				zap.Error(err),
			)
			return nil, fmt.Errorf("get all languages: %w", err)
		}
		languages = append(languages, lang)
	}
	r.logger.Info("all languages returned successfully",
		zap.Strings("languages", languages),
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllLanguages),
		zap.Int("count", len(languages)),
	)
	return languages, nil
}
func (r *LanguageRepo) GetCodeByLanguageName(ctx context.Context, name string) (*string, error) {
	row := r.pool.QueryRow(ctx, queryGetLangCodeByName, name)
	var languageID string
	err := row.Scan(&languageID)
	if err != nil {
		r.logger.Error("GetLanguageByCode",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCodeByLanguageName),
			zap.String("name", name),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetLanguageByCode: %w", err)
	}
	r.logger.Info("GetLanguageByCode successfully retrieved",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetLanguageByCode),
		zap.String("name", name),
	)
	return &languageID, nil
}
func (r *LanguageRepo) GetLanguageByCode(ctx context.Context, code string) (*language.Language, error) {
	row := r.pool.QueryRow(ctx, queryGetLangByCode, code)
	var lang language.Language
	err := row.Scan(&lang.LanguageCode, &lang.LanguageName)
	if err != nil {
		r.logger.Error("GetLanguageByCode",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetLanguageByCode),
			zap.String("code", code),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetLanguageByCode: %w", err)
	}
	r.logger.Info("GetLanguageByCode successfully retrieved",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetLanguageByCode),
		zap.String("code", code),
	)
	return &lang, nil
}
