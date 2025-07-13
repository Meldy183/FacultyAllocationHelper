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
	queryGetByProfileID = `
		SELECT profile_id, email, position, english_name, russian_name, alias,
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
		SET email = $1, position = $2, english_name = $3,
		    russian_name = $4, alias = $5, employment_type = $6, degree = $7,
		    mode = $8, start_date = $9, end_date = $10, maxload = $11, student_type = $13
		WHERE profile_id = $12
	`
	queryGetProfilesByFiler = ``

	logLayer              = "repository"
	logGetByProfileID     = "GetByProfileID"
	logCreate             = "Create"
	logUpdate             = "Update"
	logGetProfilesByFiler = "GetProfilesByFiler"
)

func (r *UserProfileRepo) GetByProfileID(ctx context.Context, profileID int64) (*userprofile.UserProfile, error) {
	row := r.pool.QueryRow(ctx, queryGetByProfileID, profileID)
	var userProfile userprofile.UserProfile
	err := row.Scan(
		&userProfile.ProfileID,
		&userProfile.Email,
		&userProfile.Position,
		&userProfile.EnglishName,
		&userProfile.RussianName,
		&userProfile.Alias,
		&userProfile.EmploymentType,
		&userProfile.StudentType,
		&userProfile.Degree,
		&userProfile.Mode,
		&userProfile.StartDate,
		&userProfile.EndDate,
		&userProfile.MaxLoad,
	)
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
	_, err := r.pool.Exec(ctx, queryUpdateUserProfile, 0, userProfile.Email, userProfile.Position,
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

func (r *UserProfileRepo) GetProfilesByFilter(ctx context.Context, institutes []int, positions []int) (error, []int64) {
	if len(institutes) == 0 {
		instRepo, err := NewInstituteRepo(r.pool, r.logger).GetAll(ctx)
		if err != nil {
			r.logger.Error("Error getting all institutes",
				zap.String("layer", logLayer),
				zap.String("function", logGetProfilesByFiler),
				zap.Error(err),
				)
		}
		for _, inst := range instRepo {
			institutes = append(institutes, inst.InstituteID)
		}
		r.logger.Info("Institutes length is zero, filtering by all institutes",
			zap.String("layer", logLayer),
			zap.String("function", logGetProfilesByFiler),
			zap.Int("institutes", len(institutes)),
			)
	}
	if len(positions) == 0 {
		posRepo, err :=
	}
}
