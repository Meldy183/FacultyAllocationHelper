package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userlanguage"
	"go.uber.org/zap"
)

var _ userlanguage.Repository = (*UserLanguageRepo)(nil)

type UserLanguageRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserLanguageRepo(pool *pgxpool.Pool, logger *zap.Logger) *UserLanguageRepo {
	return &UserLanguageRepo{pool: pool, logger: logger}
}

const (
	queryAdd              = `INSERT INTO user_language (user_language_id, profile_id, language_code) VALUES ($1, $2, $3)`
	queryGetUserLanguages = `SELECT language_code FROM user_language WHERE profile_id = $1`
)

func (r *UserLanguageRepo) Add(ctx context.Context, userLanguage *userlanguage.UserLanguage) error {
	r.logger.Info("Adding user-language to database with ID", zap.Int("ID", userLanguage.ProfileID))
	_, err := r.pool.Exec(ctx, queryAdd, userLanguage.UserLanguageID, userLanguage.ProfileID, userLanguage.LanguageCode)
	if err != nil {
		r.logger.Error("Error adding user-language to database", zap.Error(err))
		return fmt.Errorf("adding user-language to database with ID: %w", err)
	}
	return nil
}

func (r *UserLanguageRepo) GetUserLanguages(ctx context.Context, profileID int64) ([]*language.Language, error) {
	r.logger.Info("Getting user-languages from database with ID", zap.Int64("ProfileID", profileID))
	row, err := r.pool.Query(ctx, queryGetUserLanguages, profileID)
	var userLanguages []*language.Language
	if err != nil {
		r.logger.Error("Error getting user-languages", zap.Error(err))
		return nil, fmt.Errorf("getting user-languages from database with ID: %w", err)
	}
	for row.Next() {
		var languageCode string
		if err := row.Scan(&languageCode); err != nil {
			r.logger.Error("Error getting user-languages", zap.Error(err))
			return nil, fmt.Errorf("getting user-languages from database with ID: %w", err)
		}
		userLanguages = append(userLanguages, &language.Language{})
	}
	if err := row.Err(); err != nil {
		r.logger.Error("Error iterating rows", zap.Error(err))
		return nil, fmt.Errorf("getting user-languages from database with ID: %w", err)
	}
	return userLanguages, nil
}
