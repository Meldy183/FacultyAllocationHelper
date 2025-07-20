package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ staff.Repository = (*StaffRepo)(nil)

type StaffRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewStaffRepo(pool *pgxpool.Pool, logger *zap.Logger) *StaffRepo {
	return &StaffRepo{pool: pool, logger: logger}
}

const (
	queryGetStaffByInstanceID = `SELECT assignment_id, instance_id, profile_version_id, position_type,
    groups_assigned, is_confirmed, lectures_count, tutorials_count, labs_count
	FROM staff WHERE instance_id = $1`
	queryAddStaff = `INSERT INTO staff (instance_id, profile_version_id, position_type,
    groups_assigned, is_confirmed, lectures_count, tutorials_count, labs_count)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	queryUpdateStaff = ``
)

func (r *StaffRepo) GetAllStaffByInstanceID(ctx context.Context, instanceID int64) ([]*staff.Staff, error) {
	rows, err := r.pool.Query(ctx, queryGetStaffByInstanceID, instanceID)
	if err != nil {
		r.logger.Error("failed to query staffs by instance id",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetStaffByInstanceID),
			zap.Int64("instance", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to query staffs by instance id: %w", err)
	}
	defer rows.Close()
	staffs := make([]*staff.Staff, 0)
	for rows.Next() {
		var staffInstance staff.Staff
		if rows.Err() != nil {
			r.logger.Error("failed to query staffs by instance id",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetStaffByInstanceID),
				zap.Int64("instance", instanceID),
				zap.Error(rows.Err()),
			)
			return nil, fmt.Errorf("failed to query staffs by instance id: %w", err)
		}
		err = rows.Scan(
			&staffInstance.AssignmentID,
			&staffInstance.InstanceID,
			&staffInstance.ProfileVersionID,
			&staffInstance.PositionType,
			&staffInstance.GroupsAssigned,
			&staffInstance.IsConfirmed,
			&staffInstance.LecturesCount,
			&staffInstance.TutorialsCount,
			&staffInstance.LabsCount,
		)
		if err != nil {
			r.logger.Error("failed to scan staffs by instance id",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetStaffByInstanceID),
				zap.Int64("instance", instanceID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("failed to scan staffs by instance id: %w", err)
		}
		staffs = append(staffs, &staffInstance)
	}
	r.logger.Info("successfully fetched staffs by instance id",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetStaffByInstanceID),
		zap.Int64("instance", instanceID),
	)
	return staffs, nil
}
func (r *StaffRepo) AddStaff(ctx context.Context, staff *staff.Staff) error {
	err := r.pool.QueryRow(ctx, queryAddStaff,
		staff.InstanceID,
		staff.ProfileVersionID,
		staff.PositionType,
		staff.GroupsAssigned,
		staff.IsConfirmed,
		staff.LecturesCount,
		staff.TutorialsCount,
		staff.LabsCount,
	).Scan(&staff.AssignmentID)
	if err != nil {
		r.logger.Error("failed to add staff",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddStaff),
			zap.Int64("instance", staff.InstanceID),
			zap.Error(err),
		)
		return fmt.Errorf("failed to add staff: %w", err)
	}
	return nil
}
func (r *StaffRepo) UpdateStaff(ctx context.Context, staff *staff.Staff) error {
	//TODO: implement me
	panic("implement me")
	return nil
}
