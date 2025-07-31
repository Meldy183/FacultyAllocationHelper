package workload

import "context"

type Repository interface {
	GetSemesterWorkloadByVersionID(ctx context.Context, profileVersionID int64, semesterID int64) (*Workload, error)
	AddSemesterWorkload(ctx context.Context, workload *Workload) error
	UpdateSemesterWorkload(ctx context.Context, workload *Workload) error
}
