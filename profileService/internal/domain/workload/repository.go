package workload

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/transaction"
)

type Repository interface {
	GetSemesterWorkloadByVersionID(ctx context.Context, profileVersionID int64, semesterID int64) (*Workload, error)
	AddSemesterWorkload(ctx context.Context, tx transaction.Transaction, workload *Workload) error
	UpdateSemesterWorkload(ctx context.Context, tx transaction.Transaction, workload *Workload) error
}
