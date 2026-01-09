package allocationService

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/allocation"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"go.uber.org/zap"
)

var _ allocation.Service = (*Service)(nil)

type Service struct {
	logger                    *zap.Logger
	profileCourseInstanceRepo profileCourseInstance.Repository
	profileVersionRepo        profileVersion.Repository
	courseStaffRepo           staff.Repository
	courseInstanceRepo        courseInstance.Repository
	workloadRepo              workload.Repository
}

func NewService(logger *zap.Logger,
	profileCourseInstanceRepo profileCourseInstance.Repository,
	profileVersionRepo profileVersion.Repository,
	courseStaffRepo staff.Repository,
	courseInstanceRepo courseInstance.Repository,
	workloadRepo workload.Repository) *Service {
	return &Service{
		logger:                    logger,
		profileCourseInstanceRepo: profileCourseInstanceRepo,
		profileVersionRepo:        profileVersionRepo,
		courseStaffRepo:           courseStaffRepo,
		courseInstanceRepo:        courseInstanceRepo,
		workloadRepo:              workloadRepo,
	}
}
func (s *Service) AllocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType string) error {
	panic("Implement Me!")
}
func (s *Service) DeallocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType string) error {
	panic("Implement Me!")
}
