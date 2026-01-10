package allocationService

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/allocation"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ allocation.Service = (*Service)(nil)

type Service struct {
	pool               *pgxpool.Pool
	logger             *zap.Logger
	profileVersionRepo profileVersion.Repository
	courseStaffRepo    staff.Repository
	courseInstanceRepo courseInstance.Repository
	workloadRepo       workload.Repository
}

func NewService(logger *zap.Logger,
	profileVersionRepo profileVersion.Repository,
	courseStaffRepo staff.Repository,
	courseInstanceRepo courseInstance.Repository,
	workloadRepo workload.Repository) *Service {
	return &Service{
		logger:             logger,
		profileVersionRepo: profileVersionRepo,
		courseStaffRepo:    courseStaffRepo,
		courseInstanceRepo: courseInstanceRepo,
		workloadRepo:       workloadRepo,
	}
}
func (s *Service) AllocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType *string, groupsAssigned *int64) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		s.logger.Error("error starting transaction",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Error(err))
		return err
	}
	defer func() {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			s.logger.Error("error rolling back transaction",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Error(rollbackErr))
		}
	}()
	curYear := int64(time.Now().Year())
	profileVer, err := s.profileVersionRepo.GetVersionByProfileID(ctx, profileID, curYear)
	if err != nil || profileVer == nil {
		s.logger.Error("error getting Profile Version",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Int64("profileID", profileID),
			zap.Error(err))
		return err
	}
	courseStaff, err := s.courseStaffRepo.GetStaffByInstanceAndVersionID(ctx, courseInstanceID, profileVer.ProfileVersionId)
	if err != nil {
		s.logger.Error("error getting Staff",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Int64("versionID", profileVer.ProfileVersionId),
			zap.Int64("courseInstanceID", courseInstanceID),
			zap.Error(err))
		return err
	}
	if courseStaff == nil {
		courseStaff = staff.NewStaff(courseInstanceID, profileVer.ProfileVersionId, positionType, groupsAssigned)
		err := s.courseStaffRepo.AddStaff(ctx, &tx, courseStaff)
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
		if !staffIsValid(courseStaff, *positionType) {
			s.logger.Error("Faculty is already allocated to this position",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Int64("assignmentID", courseStaff.AssignmentID))
			return fmt.Errorf("Faculty alredy allocated")
		}
	}
	// Staff assignment already exists
	//TODO: staff update logic
	inst, err := s.courseInstanceRepo.GetCourseInstanceByID(ctx, courseInstanceID)
	if err != nil || inst == nil {
		s.logger.Error("error getting Course Instance",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Int64("courseInstanceID", courseInstanceID),
			zap.Error(err))
		return err
	}
	load, err := s.workloadRepo.GetSemesterWorkloadByVersionID(ctx, profileVer.ProfileVersionId, inst.SemesterID)
	if err != nil {
		s.logger.Error("error getting Course Instance",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty),
			zap.Int64("courseInstanceID", courseInstanceID),
			zap.Error(err))
		return err
	}
	if load == nil {
		load := workload.NewWorkload(profileVer.ProfileVersionId,
			inst.SemesterID,
			*courseStaff.LecturesCount,
			*courseStaff.TutorialsCount,
			*courseStaff.LabsCount)
		err := s.workloadRepo.AddSemesterWorkload(ctx, &tx, load)
		if err != nil {
			return err
		}
	} else {
		s.logger.Info("workload found successfully",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAllocateFaculty))
	}
	// TODO: workload update logic
	return tx.Commit(ctx)

}
func (s *Service) DeallocateFaculty(ctx context.Context, courseInstanceID int64, profileID int64, positionType *string) error {
	panic("Implement Me!")
}
func staffIsValid(staffMember *staff.Staff, positionType string) bool {
	switch positionType {
	case "PI":
		return staffMember.LecturesCount == nil || *staffMember.LecturesCount == 0
	case "TI":
		return staffMember.TutorialsCount == nil || *staffMember.TutorialsCount == 0
	case "TA":
		return staffMember.LabsCount == nil || *staffMember.LabsCount == 0 && staffMember.GroupsAssigned == nil || *staffMember.GroupsAssigned == 0
	default:
		return false
	}

}
