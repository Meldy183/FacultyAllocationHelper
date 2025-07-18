package workload

import (
	"context"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/facultyProfile"
)

type Service interface {
	GetSemesterWorkloadByVersionID(ctx context.Context, profileVersionID int64, semesterID int) (*Workload, error)
	AddSemesterWorkload(ctx context.Context, workload *Workload) error
	UpdateSemesterWorkload(ctx context.Context, workload *Workload) error
	GetYearWorkloadByVersionID(ctx context.Context, profileVersionID int64) ([]*facultyProfile.Classes, error)
}
