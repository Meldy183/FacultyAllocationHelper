package workload

import "context"

type Service interface {
	GetSemesterWorkloadByVersionID(ctx context.Context, profileVersionID int64, semesterID int) (*Workload, error)
	AddSemesterWorkload(ctx context.Context, workload *Workload) error
	UpdateSemesterWorkload(ctx context.Context, workload *Workload) error
}
