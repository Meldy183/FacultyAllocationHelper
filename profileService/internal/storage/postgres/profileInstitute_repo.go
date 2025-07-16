package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileInstitute.Repository = (*UserInstituteRepo)(nil)

type UserInstituteRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserInstituteRepo(pool *pgxpool.Pool, logger *zap.Logger) *UserInstituteRepo {
	return &UserInstituteRepo{pool: pool, logger: logger}
}

const (
	queryGetUserInstituteByID = `SELECT i.institute_id, i.name FROM user_institute ui JOIN institute i ON ui.institute_id = i.institute_id WHERE ui.profile_id = $1`
	queryAddUserInstitute     = `INSERT INTO user_institute (profile_id, institute_id, is_repr)
								 VALUES ($1, $2, $3)
								 RETURNING user_institute_id`
)

func (r *UserInstituteRepo) GetUserInstituteByID(ctx context.Context, profileID int64) (*institute.Institute, error) {
	row := r.pool.QueryRow(ctx, queryGetUserInstituteByID, profileID)
	var instituteById institute.Institute
	err := row.Scan(
		&instituteById.InstituteID,
		&instituteById.Name,
	)
	if err != nil {
		r.logger.Error("GetUserInstituteByID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetUserInstitute),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetUserInstituteByID: %w", err)
	}
	r.logger.Info("GetUserInstituteByID Success",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetUserInstitute),
		zap.Int64("profileID", profileID),
	)
	return &instituteById, nil
}

func (r *UserInstituteRepo) AddUserInstitute(ctx context.Context, userInstitute *profileInstitute.UserInstitute) error {
	err := r.pool.QueryRow(ctx, queryAddUserInstitute, userInstitute.ProfileID, userInstitute.InstituteID,
		userInstitute.IsRepresentative).Scan(&userInstitute.UserInstituteID)
	if err != nil {
		r.logger.Error("AddUserInstitute",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddUserInstitute),
			zap.Int64("profileID", userInstitute.ProfileID),
			zap.Error(err),
		)
		return fmt.Errorf("AddUserInstitute: %w", err)
	}
	r.logger.Info("Success of adding UserInstitute",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddUserInstitute),
		zap.Int64("profileID", userInstitute.ProfileID),
	)
	return nil
}
