package postgres

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userprofile"
	"go.uber.org/zap"
)

var _ userprofile.Repository = (*UserProfileVersionRepo)(nil)

type UserProfileVersionRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserProfileVersionRepo(pool *pgxpool.Pool, logger *zap.Logger) *UserProfileVersionRepo {
	return &UserProfileVersionRepo{pool: pool, logger: logger}
}

const (
	queryGetVersionByVersionID = `
		SELECT profile_version_id, profile_id, year, semester, lectures_count, tutorials_count, labs_count,
		elective_count, workload, maxload, position_id, employment_type, degree, mode 
		FROM user_profile_version
		WHERE profile_version_id = $1
	`

	queryInsertVersion = `
		INSERT INTO user_profile_version (
			(position_id)
		)
		VALUES ($1)
		RETURNING profile_version_id
	`

	queryUpdateVersion = `
		UPDATE user_profile_version
		SET profile_id = $1, year = $2, semester = $3, lectures_count = $4, tutorials_count = $5, labs_count = $6,
		elective_count = $7, workload = $8, maxload = $9, position_id = $10, employment_type = $11, degree = $12, mode = $13
		WHERE profile_version_id = $14
`
)

func (r *UserProfileRepo) GetProfileByID(ctx context.Context, profileID int64) (*userprofile.UserProfile, error) {
	row := r.pool.QueryRow(ctx, queryGetByProfileID, profileID)
	var userProfile userprofile.UserProfile
	err := row.Scan(
		&userProfile.ProfileID,
		&userProfile.Email,
		&userProfile.PositionID,
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
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetUserProfile failed: %w", err)
	}
	r.logger.Info("User profile found",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProfileByID),
		zap.Int64("profileID", profileID),
	)
	return &userProfile, err
}

func (r *UserProfileRepo) AddProfile(ctx context.Context, userProfile *userprofile.UserProfile) error {
	err := r.pool.QueryRow(ctx, queryInsertUserProfile,
		userProfile.Email,
		userProfile.PositionID,
		userProfile.EnglishName,
		userProfile.Alias,
	).Scan(&userProfile.ProfileID)
	if err != nil {
		r.logger.Error("Error creating user profile",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Int64("profileID", userProfile.ProfileID),
			zap.Error(err),
		)
		return fmt.Errorf("CreateUserProfile failed: %w", err)
	}
	r.logger.Info("User profile created",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddProfile),
		zap.Int64("profileId", userProfile.ProfileID),
	)
	return nil
}

func (r *UserProfileRepo) UpdateProfileByID(ctx context.Context, userProfile *userprofile.UserProfile) error {
	_, err := r.pool.Exec(ctx, queryUpdateUserProfile, 0, userProfile.Email, userProfile.PositionID,
		userProfile.EnglishName, userProfile.RussianName, userProfile.Alias, userProfile.EmploymentType,
		userProfile.Degree, userProfile.Mode, userProfile.StartDate, userProfile.EndDate, userProfile.MaxLoad,
		userProfile.ProfileID, userProfile.StudentType)
	if err != nil {
		r.logger.Error("Error updating user profile",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogUpdateFaculty),
			zap.Int64("profileId", userProfile.ProfileID),
			zap.Error(err),
		)
		return fmt.Errorf("UpdateUserProfile failed: %w", err)
	}
	r.logger.Info("User profile updated",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogUpdateFaculty),
		zap.Int64("profileId", userProfile.ProfileID),
	)
	return nil
}

func (r *UserProfileRepo) GetProfilesByFilter(ctx context.Context, institutes []int, positions []int) ([]int64, error) {
	if len(institutes) == 0 {
		instRepo, err := NewInstituteRepo(r.pool, r.logger).GetAllInstitutes(ctx)
		if err != nil {
			r.logger.Error("Error getting all institutes",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetProfilesByFilters),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetAllInstitutes failed: %w", err)
		}
		for _, inst := range instRepo {
			institutes = append(institutes, inst.InstituteID)
		}
		r.logger.Info("Institutes length is zero, filtering by all institutes",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfilesByFilters),
			zap.Int("institutes", len(institutes)),
		)
	}
	if len(positions) == 0 {
		posRepo, err := NewPositionRepo(r.pool, r.logger).GetAllPositions(ctx)
		if err != nil {
			r.logger.Error("Error getting all positions",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetProfilesByFilters),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetAllPositions failed: %w", err)
		}
		for _, pos := range posRepo {
			positions = append(positions, pos.PositionID)
		}
		r.logger.Info("Positions length is zero, filtering by all positions",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfilesByFilters),
			zap.Int("positions", len(positions)),
		)
	}
	rows, err := r.pool.Query(ctx, queryGetProfilesByFiler, institutes, positions)
	if err != nil {
		r.logger.Error("Error getting all profileIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfilesByFilters),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetAllProfileIDs failed: %w", err)
	}
	defer rows.Close()
	profileIDs := make([]int64, 0)
	for rows.Next() {
		var profileID int64
		err = rows.Scan(&profileID)
		if err != nil {
			r.logger.Error("Error getting single profileID",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetProfilesByFilters),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetAllProfileIDs failed: %w", err)
		}
		profileIDs = append(profileIDs, profileID)
	}
	r.logger.Info("ProfileIDs received successfully",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProfilesByFilters),
		zap.Int("profileIDs", len(profileIDs)),
	)
	return profileIDs, nil
}
