package workload

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetSemesterWorkloadByVersionID(ctx context.Context, profileVersionID uuid.UUID, semesterID uuid.UUID) (*Workload, error)
	AddSemesterWorkload(ctx context.Context, workload *Workload) error
	UpdateSemesterWorkload(ctx context.Context, workload *Workload) error
}
