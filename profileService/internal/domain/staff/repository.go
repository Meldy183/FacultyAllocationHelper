package staff

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetAllStaffByInstanceID(ctx context.Context, instanceID int64) ([]*Staff, error)
	GetStaffByInstanceAndVersionID(ctx context.Context, instanceID int64, versionID int64) (*Staff, error)
	AddStaff(ctx context.Context, tx *pgx.Tx, staff *Staff) error
	UpdateStaff(ctx context.Context, tx *pgx.Tx, staff *Staff) error
	DeleteStaff(ctx context.Context, tx *pgx.Tx, staffID int64) error
}
