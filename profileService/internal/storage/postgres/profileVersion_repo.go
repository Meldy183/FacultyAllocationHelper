package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"go.uber.org/zap"
)

var _ profileVersion.Repository = (*ProfileVersionRepo)(nil)

type ProfileVersionRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewUserProfileVersionRepo(pool *pgxpool.Pool, logger *zap.Logger) *ProfileVersionRepo {
	return &ProfileVersionRepo{pool: pool, logger: logger}
}

const (
	queryGetVersionByVersionID = `
		SELECT profile_version_id, profile_id, year, semester_id, lectures_count, tutorials_count, labs_count,
		elective_count, workload, maxload, position_id, employment_type, degree, mode 
		FROM user_profile_version
		WHERE profile_version_id = $1
	`
	queryInsertVersion = `
		INSERT INTO user_profile_version (
			(position_id, profile_id, semester_id, year)
		)
		VALUES ($1, $2, $3, $4)
		RETURNING profile_version_id
	`
	queryUpdateVersion = `
		UPDATE user_profile_version
		SET profile_id = $1, year = $2, semester = $3, lectures_count = $4, tutorials_count = $5, labs_count = $6,
		elective_count = $7, workload = $8, maxload = $9, position_id = $10, employment_type = $11, degree = $12, mode = $13
		WHERE profile_version_id = $14
`
)

func (r *ProfileVersionRepo) AddProfileVersion(ctx context.Context,
	profileVersion *profileVersion.ProfileVersion) error {

}
