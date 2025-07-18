package postgres

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"go.uber.org/zap"
)

var _ facultyProfile.Repository = (*FacultyProfileRepo)(nil)

type FacultyProfileRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewFacultyProfileRepo(pool *pgxpool.Pool, logger *zap.Logger) *FacultyProfileRepo {
	return &FacultyProfileRepo{pool: pool, logger: logger}
}

const (
	queryGetByProfileID = `
		SELECT profile_id, email, english_name, russian_name, alias, start_date, end_date
		FROM user_profile
		WHERE profile_id = $1
	`

	queryInsertUserProfile = `
		INSERT INTO user_profile (
			email, english_name, alias
		)
		VALUES ($1, $2, $3)
		RETURNING profile_id
	`

	queryUpdateUserProfile = `
		UPDATE user_profile
		SET email = $1, english_name = $2,
		    russian_name = $3, alias = $4, start_date = $5, end_date = $6
		WHERE profile_id = $7
	`
	queryGetProfileIDsByInstituteIDs = `SELECT profile_id FROM user_institute WHERE institute_id = ANY($1)
ORDER BY profile_id`
	queryGetProfileIDsByPositionIDs = `SELECT profile_id from user_profile_version where position_id = ANY($1)
ORDER BY profile_id`
)

func (r *FacultyProfileRepo) GetProfileByID(ctx context.Context, profileID int64) (*facultyProfile.UserProfile, error) {
	row := r.pool.QueryRow(ctx, queryGetByProfileID, profileID)
	var userProfile facultyProfile.UserProfile
	err := row.Scan(
		&userProfile.ProfileID,
		&userProfile.Email,
		&userProfile.EnglishName,
		&userProfile.RussianName,
		&userProfile.Alias,
		&userProfile.StartDate,
		&userProfile.EndDate,
	)
	if err != nil {
		r.logger.Error("Error getting user facultyProfile",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetUserProfile failed: %w", err)
	}
	r.logger.Info("User facultyProfile found",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProfileByID),
		zap.Int64("profileID", profileID),
	)
	return &userProfile, nil
}

func (r *FacultyProfileRepo) AddProfile(ctx context.Context, userProfile *facultyProfile.UserProfile) error {
	err := r.pool.QueryRow(ctx, queryInsertUserProfile,
		userProfile.Email,
		userProfile.EnglishName,
		userProfile.Alias,
	).Scan(&userProfile.ProfileID)
	if err != nil {
		r.logger.Error("Error creating user facultyProfile",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Int64("profileID", userProfile.ProfileID),
			zap.Error(err),
		)
		return fmt.Errorf("CreateUserProfile failed: %w", err)
	}
	r.logger.Info("User facultyProfile created",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddProfile),
		zap.Int64("profileId", userProfile.ProfileID),
	)
	return nil
}

func (r *FacultyProfileRepo) UpdateProfileByID(ctx context.Context, userProfile *facultyProfile.UserProfile) error {
	_, err := r.pool.Exec(ctx, queryUpdateUserProfile,
		userProfile.Email,
		userProfile.EnglishName,
		userProfile.RussianName,
		userProfile.Alias,
		userProfile.StartDate,
		userProfile.EndDate,
		userProfile.ProfileID,
	)
	if err != nil {
		r.logger.Error("Error updating user facultyProfile",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogUpdateFaculty),
			zap.Int64("profileId", userProfile.ProfileID),
			zap.Error(err),
		)
		return fmt.Errorf("UpdateUserProfile failed: %w", err)
	}
	r.logger.Info("User facultyProfile updated",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogUpdateFaculty),
		zap.Int64("profileId", userProfile.ProfileID),
	)
	return nil
}

func (r *FacultyProfileRepo) GetProfileIDsByInstituteIDs(ctx context.Context, instituteIDs []int) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetProfileIDsByInstituteIDs, instituteIDs)
	if err != nil {
		r.logger.Error("Error getting facultyProfile by instituteIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfileIDsByInstituteIDs),
			zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetProfileIDsByInstituteIDs failed: %w", err)
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			r.logger.Error("Error getting facultyProfile by instituteIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetProfileIDsByInstituteIDs),
				zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetProfileIDsByInstituteIDs failed: %w", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting facultyProfile by instituteIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfileIDsByInstituteIDs),
			zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetProfileIDsByInstituteIDs failed: %w", err)
	}
	r.logger.Info("facultyProfile by instituteIDs found",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProfileIDsByInstituteIDs),
		zap.Int64s("instituteIDs", ids),
	)
	return ids, nil
}
func (r *FacultyProfileRepo) GetProfileIDsByPositionIDs(ctx context.Context, positionIDs []int) ([]int64, error) {
	rows, err := r.pool.Query(ctx, queryGetProfileIDsByPositionIDs, positionIDs)
	if err != nil {
		r.logger.Error("Error getting facultyProfile by positionIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfileIDsByPositionIDs),
			zap.String("positionIDs", fmt.Sprintf("%v", positionIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetProfileIDsByPositionIDs failed: %w", err)
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			r.logger.Error("Error getting facultyProfile by positionIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetProfileIDsByPositionIDs),
				zap.String("positionIDs", fmt.Sprintf("%v", positionIDs)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetProfileIDsByPositionIDs failed: %w", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting facultyProfile by positionIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProfileIDsByPositionIDs),
			zap.String("positionIDs", fmt.Sprintf("%v", positionIDs)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetProfileIDsByPositionIDs failed: %w", err)
	}
	r.logger.Info("facultyProfile by positionIDs found",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetProfileIDsByPositionIDs),
		zap.Int64s("positionIDs", ids),
	)
	return ids, nil
}
