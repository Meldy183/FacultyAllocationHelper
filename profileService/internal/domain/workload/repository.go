package workload

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetSemesterWorkloadByVersionID(ctx context.Context, profileVersionID int64, semesterID int64) (*Workload, error)
	AddSemesterWorkload(ctx context.Context, tx *pgx.Tx, workload *Workload) error
	UpdateSemesterWorkload(ctx context.Context, tx *pgx.Tx, workload *Workload) error
}
