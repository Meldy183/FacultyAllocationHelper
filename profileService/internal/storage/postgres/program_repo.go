package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/program"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ program.Repository = (*ProgramRepo)(nil)

type ProgramRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewProgramRepo(pool *pgxpool.Pool, logger *zap.Logger) *ProgramRepo {
	return &ProgramRepo{pool: pool, logger: logger}
}

const (
	queryProgramByID    = `SELECT name FROM program WHERE code = $1`
	queryGetAllPrograms = `SELECT program_id, name FROM program`
)

func (r *ProgramRepo) GetProgramNameByID(ctx context.Context, id int) ([]*string, error) {
	row := r.pool.QueryRow(ctx, queryGetAllPrograms, id)
	var program program.Program
	err := row.Scan(&program.Name)
	if err != nil {
		r.logger.Error("GetProgramNameByID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProgramNameByID),
			zap.Int("id", id),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetLanguageByCode: %w", err)
	}
	r.logger.Info("GetLanguageByCode",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetLanguageByCode),
		zap.Int("id", id),
	)
	return &program, nil
}
func (r *ProgramRepo) GetAllPrograms(ctx context.Context, id int) (*string, error) {
	rows, err := r.pool.Query(ctx, queryProgramByID)
	if err != nil {
		r.logger.Error("failed to query all languages",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllLanguages),
			zap.Error(err),
		)
		return nil, fmt.Errorf("get all languages: %w", err)
	}
	defer rows.Close()
	var languages []*language.Language
	for rows.Next() {
		if rows.Err() != nil {
			r.logger.Error("failed to query all languages",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllLanguages),
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
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllLanguages),
				zap.Error(err),
			)
			return nil, fmt.Errorf("get all languages: %w", err)
		}
		languages = append(languages, &lang)
	}
	r.logger.Info("all languages returned",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllLanguages),
		zap.Int("count", len(languages)),
	)
	return languages, nil
}
