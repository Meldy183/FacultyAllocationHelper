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
	queryGetStaffByInstanceID = ``
	queryAddStaff             = ``
	queryUpdateStaff          = ``
)

func (r *StaffRepo) GetStaffByInstanceID(ctx context.Context, instanceID int) ([]*staff.Staff, error) {
	rows, err := r.pool.Query(ctx, queryGetStaffByInstanceID, instanceID)
	if err != nil {
		r.logger.Error("failed to query staffs by instance id",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetStaffByInstanceID),
			zap.Int("instance", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to query staffs by instance id: %w", err)
	}
	staffs := make([]*staff.Staff, 0)
	for rows.Next() {
		var staffInstance staff.Staff
		
	}
}
func (r *StaffRepo) AddStaff(ctx context.Context, staff *staff.Staff) error {

}
func (r *StaffRepo) UpdateStaff(ctx context.Context, staff *staff.Staff) error {

}
