package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userprofile"
	"go.uber.org/zap"
)

var _ userprofile.Repository = (*UserProfileRepo)(nil)

type UserProfileRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserProfileRepo(pool *pgxpool.Pool, logger *zap.Logger) *UserProfileRepo {
	return &UserProfileRepo{pool: pool, logger: logger}
}

const (
	queryGetByUserID = `
		SELECT profile_id, user_id, email, position, english_name, russian_name, alias,
		       employment_type, degree, mode, start_date, end_date, maxload, student_type
		FROM user_profile
		WHERE user_id = $1
	`

	queryGetByProfileID = `
		SELECT profile_id, user_id, email, position, english_name, russian_name, alias,
		       employment_type, degree, mode, start_date, end_date, maxload, student_type
		FROM user_profile
		WHERE profile_id = $1
	`

	queryInsertUserProfile = `
		INSERT INTO user_profile (
			email, position, english_name, alias
		)
		VALUES ($1, $2, $3, $4)
		RETURNING profile_id
	`

	queryUpdateUserProfile = `
		UPDATE user_profile
		SET user_id = $1, email = $2, position = $3, english_name = $4,
		    russian_name = $5, alias = $6, employment_type = $7, degree = $8,
		    mode = $9, start_date = $10, end_date = $11, maxload = $12, student_type = $14
		WHERE profile_id = $13
	`
	logLayer          = "repository"
	logGetByUserID    = "GetByUserID"
	logGetByProfileID = "GetByProfileID"
	logCreate         = "Create"
	logUpdate         = "Update"
)

func (r *UserProfileRepo) GetByUserID(ctx context.Context, userID string) (*userprofile.UserProfile, error) {
	row := r.pool.QueryRow(ctx, queryGetByUserID, userID)
	var userProfile userprofile.UserProfile
	err := row.Scan(
		&userProfile.ProfileID,
		&userProfile.UserID,
		&userProfile.Email,
		&userProfile.Position,
		&userProfile.EnglishName,
		&userProfile.RussianName,
		&userProfile.Alias,
		&userProfile.EmploymentType,
		&userProfile.Degree,
		&userProfile.StartDate,
		&userProfile.EndDate,
		&userProfile.MaxLoad,
		&userProfile.StudentType)
	if err != nil {
		r.logger.Error("Error getting user profile",
			zap.String("layer", logLayer),
			zap.String("function", logGetByUserID),
			zap.String("userid", userID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetUserProfile failed: %w", err)
	}
	r.logger.Info("User profile found",
		zap.String("layer", logLayer),
		zap.String("function", logGetByUserID),
		zap.String("userId", userID),
	)
	return &userProfile, err
}

func (r *UserProfileRepo) GetByProfileID(ctx context.Context, profileID int64) (*userprofile.UserProfile, error) {
	row := r.pool.QueryRow(ctx, queryGetByProfileID, profileID)
	var userProfile userprofile.UserProfile
	err := row.Scan(
		&userProfile.ProfileID,
		&userProfile.UserID,
		&userProfile.Email,
		&userProfile.Position,
		&userProfile.EnglishName,
		&userProfile.RussianName,
		&userProfile.Alias,
		&userProfile.EmploymentType,
		&userProfile.Degree,
		&userProfile.StartDate,
		&userProfile.EndDate,
		&userProfile.StudentType)
	if err != nil {
		r.logger.Error("Error getting user profile",
			zap.String("layer", logLayer),
			zap.String("function", logGetByProfileID),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetUserProfile failed: %w", err)
	}
	r.logger.Info("User profile found",
		zap.String("layer", logLayer),
		zap.String("function", logGetByProfileID),
		zap.Int64("profileID", profileID),
	)
	return &userProfile, err
}

func (r *UserProfileRepo) Create(ctx context.Context, userProfile *userprofile.UserProfile) error {
	err := r.pool.QueryRow(ctx, queryInsertUserProfile,
		userProfile.Email,
		userProfile.Position,
		userProfile.EnglishName,
		userProfile.Alias,
	).Scan(&userProfile.ProfileID)

	if err != nil {
		r.logger.Error("Error creating user profile",
			zap.String("layer", logLayer),
			zap.String("function", logCreate),
			zap.Int64("profileID", userProfile.ProfileID),
			zap.Error(err),
		)
		return fmt.Errorf("CreateUserProfile failed: %w", err)
	}
	r.logger.Info("User profile created",
		zap.String("layer", logLayer),
		zap.String("function", logCreate),
		zap.Int64("profileId", userProfile.ProfileID),
	)
	return nil
}

func (r *UserProfileRepo) Update(ctx context.Context, userProfile *userprofile.UserProfile) error {
	_, err := r.pool.Exec(ctx, queryUpdateUserProfile, userProfile.UserID, userProfile.Email, userProfile.Position,
		userProfile.EnglishName, userProfile.RussianName, userProfile.Alias, userProfile.EmploymentType,
		userProfile.Degree, userProfile.Mode, userProfile.StartDate, userProfile.EndDate, userProfile.MaxLoad,
		userProfile.ProfileID, userProfile.StudentType)
	if err != nil {
		r.logger.Error("Error updating user profile",
			zap.String("layer", logLayer),
			zap.String("function", logUpdate),
			zap.Int64("profileId", userProfile.ProfileID),
			zap.Error(err),
		)
		return fmt.Errorf("UpdateUserProfile failed: %w", err)
	}
	r.logger.Info("User profile updated",
		zap.String("layer", logLayer),
		zap.String("function", logUpdate),
		zap.Int64("profileId", userProfile.ProfileID),
	)
	return nil
}
