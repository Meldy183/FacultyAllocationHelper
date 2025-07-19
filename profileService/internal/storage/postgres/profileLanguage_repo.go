package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileLanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileLanguage.Repository = (*UserLanguageRepo)(nil)

type UserLanguageRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserLanguageRepo(pool *pgxpool.Pool, logger *zap.Logger) *UserLanguageRepo {
	return &UserLanguageRepo{pool: pool, logger: logger}
}

const (
	queryAdd              = `INSERT INTO user_language (user_language_id, profile_id, code) VALUES ($1, $2, $3)`
	queryGetUserLanguages = `SELECT code FROM user_language WHERE profile_id = $1`
)

func (r *UserLanguageRepo) AddUserLanguage(ctx context.Context, userLanguage *profileLanguage.UserLanguage) error {
	_, err := r.pool.Exec(ctx, queryAdd, userLanguage.UserLanguageID, userLanguage.ProfileID, userLanguage.LanguageCode)
	if err != nil {
		r.logger.Error("Error adding user-language to database",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddUserLanguage),
			zap.Error(err),
		)
		return fmt.Errorf("adding user-language to database with LabID: %w", err)
	}
	return nil
}

func (r *UserLanguageRepo) GetUserLanguages(ctx context.Context, profileID int64) ([]*language.Language, error) {
	r.logger.Info("Getting user-languages from database with LabID", zap.Int64("ProfileID", profileID))
	rows, err := r.pool.Query(ctx, queryGetUserLanguages, profileID)
	var userLanguages []*language.Language
	if err != nil {
		r.logger.Error("Error getting user-languages",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetUserLanguages),
			zap.Error(err),
		)
		return nil, fmt.Errorf("getting user-languages from database with LabID: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var languageCode string
		if err = rows.Scan(&languageCode); err != nil {
			r.logger.Error("Error getting user-languages",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetUserLanguages),
				zap.Error(err),
			)
			return nil, fmt.Errorf("getting user-languages from database with LabID: %w", err)
		}
		userLanguages = append(userLanguages, &language.Language{})
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("Error iterating rows",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetUserLanguages),
			zap.Error(err),
		)
		return nil, fmt.Errorf("getting user-languages from database with LabID: %w", err)
	}
	r.logger.Info("Successfully got user-languages",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetUserLanguages),
		zap.Int("user-languages", len(userLanguages)),
	)
	return userLanguages, nil
}
