package allocationService

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/allocation"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
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
	var profileInstance profileCourseInstance.ProfileVersionCourseInstance
	profileInstance.CourseInstanceID = courseInstanceID
	profileInstance.ProfileVersionID = profileID
	err := s.profileCourseInstanceRepo.AddCourseInstance(ctx, &profileInstance)
	if err != nil {
		s.logger.Error("Error adding Course Instance to Profile link",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Error(err),
		)
		return err
	}
	s.logger.Info("Course Instance to Profile link added successfully",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAllocateFaculty),
	)
	courseStaff, err := s.courseStaffRepo.GetStaffByInstanceAndVersionID(ctx, courseInstanceID, profileID)
	if err != nil {
		s.logger.Error("error getting Staff",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Int64("profileID", profileID),
			zap.Int64("courseInstanceID", courseInstanceID),
			zap.Error(err))
		return err
	}
	if courseStaff == nil {
		var staffAssignment staff.Staff
		staffAssignment.InstanceID = courseInstanceID
		staffAssignment.ProfileVersionID = profileID
		staffAssignment.PositionType = &positionType
		// TODO: create default staff workload fields or pass them as parameters
		err := s.courseStaffRepo.AddStaff(ctx, &staffAssignment)
		if err != nil {
			s.logger.Error("Error adding Course Staff",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Error(err),
			)
			return err
		}
	} else {
		s.logger.Info("Course Staff found successfully",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty))
	}
	// Staff assignment already exists
	//TODO: staff update logic
	inst, err := s.courseInstanceRepo.GetCourseInstanceByID(ctx, courseInstanceID)
	if err != nil {
		s.logger.Error("error getting Course Instance",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Int64("courseInstanceID", courseInstanceID),
			zap.Error(err))
		return err
	}
	load, err := s.workloadRepo.GetSemesterWorkloadByVersionID(ctx, profileID, inst.SemesterID)
	if err != nil {
		s.logger.Error("error getting Course Instance",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Int64("courseInstanceID", courseInstanceID),
			zap.Error(err))
		return err
	}
	if load == nil {
		var workloadEntry workload.Workload
		workloadEntry.ProfileVersionID = profileID
		workloadEntry.SemesterID = inst.SemesterID
		// TODO: create default workload fields or pass them as parameters
		err := s.workloadRepo.AddSemesterWorkload(ctx, &workloadEntry)
		if err != nil {
			return err
		}
	} else {
		s.logger.Info("workload found successfully",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty))
	}
	// TODO: workload update logic
	return nil

}
func (s *Service) DeallocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType string) error {
	panic("Implement Me!")
}
