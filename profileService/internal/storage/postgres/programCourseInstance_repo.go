package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	programcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/programCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ programcourseinstance.Repository = (*ProgramCourseRepo)(nil)

type ProgramCourseRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewProgramCourseRepo(pool *pgxpool.Pool, logger *zap.Logger) *ProgramCourseRepo {
	return &ProgramCourseRepo{pool: pool, logger: logger}
}

const (
	queryProgramCourseByID          = `SELECT program_course_instance_id, program_id, instance_id FROM program_course_instance WHERE instance_id = $1`
	queryAddProgramToCourseInstance = `INSERT INTO program_course_instance (program_id, instance_id) VALUES ($1, $2) RETURNING program_course_instance_id`
)

func (r *ProgramCourseRepo) AddProgramToCourseInstance(ctx context.Context, programCourseInstance *programcourseinstance.ProgramCourseInstance) error {
	err := r.pool.QueryRow(ctx, queryAddProgramToCourseInstance, programCourseInstance.ProgramID, programCourseInstance.CourseInstanceID).Scan(&programCourseInstance.ProgramCourseID)
	if err != nil {
		r.logger.Error("Error adding program to courseInstance",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddNewProgram),
			zap.Error(err),
		)
		return fmt.Errorf("error adding program to courseInstance: %v", err)
	}
	return nil
}

func (r *ProgramCourseRepo) GetProgramCourseInstancesByCourseID(ctx context.Context, id int64) ([]*programcourseinstance.ProgramCourseInstance, error) {
	rows, err := r.pool.Query(ctx, queryProgramCourseByID, id)
	if err != nil {
		r.logger.Error("failed to Get ProgramCourses By course ID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetProgramCourseByID),
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetProgramCoursesByCourseIDs: %w", err)
	}

	defer rows.Close()
	var instances []*programcourseinstance.ProgramCourseInstance
	for rows.Next() {
		var instance programcourseinstance.ProgramCourseInstance
		err := rows.Scan(&instance.ProgramCourseID, &instance.ProgramID, &instance.CourseInstanceID)
		if err != nil {
			r.logger.Error("Error getting programCourses by courseIDs",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetCourseInstanceByProgramID),
				zap.String("course id", fmt.Sprintf("%v", id)),
				zap.Error(err),
			)
			return nil, fmt.Errorf("GetProgramCourseInstancesByCourseIDs failed: %w", err)
		}
		instances = append(instances, &instance)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error getting programCourses by courseIDs",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetCourseInstanceByProgramID),
			zap.String("course id", fmt.Sprintf("%v", id)),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetCourseInstanceByProgramIDs failed: %w", err)
	}
	r.logger.Info("programs found by course id",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetCourseInstanceByProgramID),
		zap.Int("instancesLen", len(instances)),
	)
	return instances, nil
}
