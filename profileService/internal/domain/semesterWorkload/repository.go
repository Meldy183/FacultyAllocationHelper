package semesterworkload

import "context"

type Repository interface {
	GetSemesterWorkloadByProfileVersionID(ctx context.Context, profileVersionID int64) (*SemesterWorkload, error)
}
