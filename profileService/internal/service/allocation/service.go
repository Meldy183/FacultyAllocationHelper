package allocationService

import (
	"context"
	"fmt"
	"time"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/allocation"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/transaction"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ allocation.Service = (*Service)(nil)

type Service struct {
	uow                transaction.UnitOfWork
	logger             *zap.Logger
	profileVersionRepo profileVersion.Repository
	courseStaffRepo    staff.Repository
	courseInstanceRepo courseInstance.Repository
	workloadRepo       workload.Repository
}

func NewService(uow transaction.UnitOfWork, logger *zap.Logger,
	profileVersionRepo profileVersion.Repository,
	courseStaffRepo staff.Repository,
	courseInstanceRepo courseInstance.Repository,
	workloadRepo workload.Repository) *Service {
	return &Service{
		uow:                uow,
		logger:             logger,
		profileVersionRepo: profileVersionRepo,
		courseStaffRepo:    courseStaffRepo,
		courseInstanceRepo: courseInstanceRepo,
		workloadRepo:       workloadRepo,
	}
}
func (s *Service) AllocateFaculty(ctx context.Context, courseID int64, profileID int64, positionType *string, groupsAssigned *int64) (*staff.Staff, error) {
	s.logger.Info("allocation started")
	var assignment *staff.Staff
	err := s.uow.Do(ctx, func(tx transaction.Transaction) error {
		curYear := int64(time.Now().Year())
		profileVer, err := s.profileVersionRepo.GetVersionByProfileID(ctx, profileID, curYear)
		if err != nil || profileVer == nil {
			s.logger.Error("error getting Profile Version",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Int64("profileID", profileID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}
		inst, err := s.courseInstanceRepo.GetCourseInstanceByID(ctx, courseID)
		if err != nil || inst == nil {
			s.logger.Error("error getting Course Instance",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Int64("courseInstanceID", courseID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}
		courseStaff, err := s.courseStaffRepo.GetStaffByInstanceAndVersionID(ctx, inst.InstanceID, profileVer.ProfileVersionId)
		if err != nil {
			s.logger.Error("error getting Staff",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Int64("versionID", profileVer.ProfileVersionId),
				zap.Int64("courseInstanceID", courseID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}

		if courseStaff == nil {
			courseStaff = staff.NewStaff(inst.InstanceID, profileVer.ProfileVersionId, positionType, groupsAssigned)
			err := s.courseStaffRepo.AddStaff(ctx, tx, courseStaff)
			if err != nil {
				s.logger.Error("Error adding Course Staff",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogAllocateFaculty),
					zap.Error(err),
				)
				return fmt.Errorf("Internal server error")
			}
			s.logger.Info("Course Staff added successfully",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty))
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
			switch *positionType {
			case "PI":
				if courseStaff.LecturesCount == nil {
					courseStaff.LecturesCount = new(int64)
				}
				*courseStaff.LecturesCount = 15
				if *courseStaff.PositionType != *positionType {
					courseStaff.PositionType = positionType
				}
			case "TI":
				if courseStaff.TutorialsCount == nil {
					courseStaff.TutorialsCount = new(int64)
				}
				*courseStaff.TutorialsCount = 15
				if *courseStaff.PositionType != "PI" {
					courseStaff.PositionType = positionType
				}
			case "TA":
				if courseStaff.LabsCount == nil {
					courseStaff.LabsCount = new(int64)
				}
				*courseStaff.LabsCount = 15 * int64(*groupsAssigned)
			}
			err := s.courseStaffRepo.UpdateStaff(ctx, tx, courseStaff)
			if err != nil {
				s.logger.Error("Error updating Course Staff",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogAllocateFaculty),
					zap.Error(err),
				)
				return fmt.Errorf("Internal server error")
			}
			s.logger.Info("Course Staff updated successfully",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty))
		}
		// update course instance logic:
		// if *positionType == "TA" {
		// 	if *inst.GroupsTaken+*groupsAssigned > inst.GroupsNeeded {
		// 		s.logger.Error("Too much groups for this course",
		// 			zap.String("layer", logctx.LogServiceLayer),
		// 			zap.String("function", logctx.LogAllocateFaculty),
		// 			zap.Int64("courseInstanceID", courseID),
		// 			zap.Int64("groups needed", inst.GroupsNeeded),
		// 			zap.Int64("groups taken", *inst.GroupsTaken),
		// 			zap.Error(err))
		// 		return fmt.Errorf("Too much groups")

		// 	} else {
		// 		*inst.GroupsTaken += *groupsAssigned
		// 		err := s.courseInstanceRepo.UpdateCourseInstanceByID(ctx, &tx, inst.InstanceID, inst)
		// 		if err != nil {
		// 			s.logger.Error("error updating course instance",
		// 				zap.String("layer", logctx.LogServiceLayer),
		// 				zap.String("function", logctx.LogAllocateFaculty),
		// 				zap.Int64("courseInstanceID", courseID),
		// 				zap.Error(err))
		// 			return fmt.Errorf("Internal server error")
		// 		}
		// 	}
		// } else {
		// 	if positionOccupied(inst, *positionType) {
		// 		s.logger.Error("Position already occupied for this course",
		// 			zap.String("layer", logctx.LogServiceLayer),
		// 			zap.String("function", logctx.LogAllocateFaculty),
		// 			zap.Int64("courseInstanceID", courseID),
		// 			zap.String("position", *positionType),
		// 			zap.Error(err))
		// 		return fmt.Errorf("Position already occupied")
		// 	} else {
		// 		switch *positionType {
		// 		case "PI":
		// 			*inst.PIAllocationStatus = courseInstance.Status("allocated")
		// 		case "TI":
		// 			*inst.TIAllocationStatus = courseInstance.Status("allocated")
		// 		default:
		// 			s.logger.Error("Invalid position",
		// 				zap.String("layer", logctx.LogServiceLayer),
		// 				zap.String("function", logctx.LogAllocateFaculty),
		// 				zap.String("position", *positionType),
		// 				zap.Error(err))
		// 			return fmt.Errorf("Invalid position")
		// 		}
		// 		err := s.courseInstanceRepo.UpdateCourseInstanceByID(ctx, &tx, inst.InstanceID, inst)
		// 		if err != nil {
		// 			s.logger.Error("error updating course instance",
		// 				zap.String("layer", logctx.LogServiceLayer),
		// 				zap.String("function", logctx.LogAllocateFaculty),
		// 				zap.Int64("courseInstanceID", courseID),
		// 				zap.Error(err))
		// 			return fmt.Errorf("Internal server error")
		// 		}
		// 	}
		// }

		load, err := s.workloadRepo.GetSemesterWorkloadByVersionID(ctx, profileVer.ProfileVersionId, inst.SemesterID)
		if err != nil {
			s.logger.Error("error getting Course Instance",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Int64("courseInstanceID", courseID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}
		if load == nil {
			load := workload.NewWorkload(profileVer.ProfileVersionId,
				inst.SemesterID,
				*courseStaff.LecturesCount,
				*courseStaff.TutorialsCount,
				*courseStaff.LabsCount)
			err := s.workloadRepo.AddSemesterWorkload(ctx, tx, load)
			if err != nil {
				s.logger.Error("error adding workload",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogAllocateFaculty),
					zap.Error(err))
				return fmt.Errorf("Internal server error")
			}
			s.logger.Info("Course Staff found successfully",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty))
		} else {
			s.logger.Info("workload found successfully",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty))
			switch *positionType {
			case "PI":
				load.LecturesCount += 15
			case "TI":
				load.TutorialsCount += 15
			case "TA":
				load.LabsCount += 15 * int64(*groupsAssigned)
			}
			err := s.workloadRepo.UpdateSemesterWorkload(ctx, tx, load)
			if err != nil {
				s.logger.Error("Error updating Workload",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogAllocateFaculty),
					zap.Error(err),
				)
				return fmt.Errorf("Internal server error")
			}
			s.logger.Info("workload updated successfully",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty))
		}
		assignment = courseStaff
		return nil
	})
	if err != nil {
		return nil, err
	}
	return assignment, err
}
func (s *Service) DeallocateFaculty(ctx context.Context, courseID int64, profileID int64, positionType *string, groupsAssigned *int64) error {
	s.logger.Info("Deallocation started")
	err := s.uow.Do(ctx, func(tx transaction.Transaction) error {
		curYear := int64(time.Now().Year())
		profileVer, err := s.profileVersionRepo.GetVersionByProfileID(ctx, profileID, curYear)
		if err != nil || profileVer == nil {
			s.logger.Error("error getting Profile Version",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogDeallocateFaculty),
				zap.Int64("profileID", profileID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}
		inst, err := s.courseInstanceRepo.GetCourseInstanceByID(ctx, courseID)
		if err != nil || inst == nil {
			s.logger.Error("error getting Course Instance",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogDeallocateFaculty),
				zap.Int64("courseInstanceID", courseID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}
		courseStaff, err := s.courseStaffRepo.GetStaffByInstanceAndVersionID(ctx, inst.InstanceID, profileVer.ProfileVersionId)
		if err != nil {
			s.logger.Error("error getting Staff",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogDeallocateFaculty),
				zap.Int64("versionID", profileVer.ProfileVersionId),
				zap.Int64("courseInstanceID", inst.InstanceID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}
		if courseStaff == nil {
			s.logger.Info("No staff found",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogDeallocateFaculty),
				zap.Int64("versionID", profileVer.ProfileVersionId),
				zap.Int64("courseInstanceID", inst.InstanceID))
			return fmt.Errorf("Staff is not allocated to this course")
		} else {
			s.logger.Info("Course Staff found successfully",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty))
			switch *positionType {
			case "PI":
				if courseStaff.LecturesCount == nil || *courseStaff.LecturesCount == 0 {
					s.logger.Info("User not allocated as PI on this course",
						zap.String("layer", logctx.LogServiceLayer),
						zap.String("function", logctx.LogDeallocateFaculty),
						zap.Int64("versionID", profileVer.ProfileVersionId),
						zap.Int64("courseInstanceID", inst.InstanceID))
					return fmt.Errorf("User not allocated as PI on this course")
				}
				if courseStaff.TutorialsCount != nil && *courseStaff.TutorialsCount != 0 {
					*courseStaff.PositionType = "TI"
				} else if courseStaff.LabsCount != nil && *courseStaff.LabsCount != 0 {
					*courseStaff.PositionType = "TA"
				} else {
					courseStaff.PositionType = nil
				}
				courseStaff.LecturesCount = nil

			case "TI":
				if courseStaff.TutorialsCount == nil || *courseStaff.TutorialsCount == 0 {
					s.logger.Info("User not allocated as TI on this course",
						zap.String("layer", logctx.LogServiceLayer),
						zap.String("function", logctx.LogDeallocateFaculty),
						zap.Int64("versionID", profileVer.ProfileVersionId),
						zap.Int64("courseInstanceID", inst.InstanceID))
					return fmt.Errorf("User not allocated as TI on this course")
				}
				if courseStaff.LabsCount != nil && *courseStaff.LabsCount != 0 {
					*courseStaff.PositionType = "TA"
				} else {
					courseStaff.PositionType = nil
				}
				courseStaff.TutorialsCount = nil
			case "TA":
				if courseStaff.LabsCount == nil || *courseStaff.LabsCount == 0 {
					s.logger.Info("User not allocated as TA on this course",
						zap.String("layer", logctx.LogServiceLayer),
						zap.String("function", logctx.LogDeallocateFaculty),
						zap.Int64("versionID", profileVer.ProfileVersionId),
						zap.Int64("courseInstanceID", inst.InstanceID))
					return fmt.Errorf("User not allocated as TA on this course")
				}
				*courseStaff.LabsCount = 15 * int64(*groupsAssigned)
				if *courseStaff.LabsCount == 0 {
					courseStaff.LabsCount = nil
				}
				courseStaff.PositionType = nil
			}
			if courseStaff.PositionType != nil {
				err := s.courseStaffRepo.UpdateStaff(ctx, tx, courseStaff)
				if err != nil {
					s.logger.Error("Error updating Course Staff",
						zap.String("layer", logctx.LogServiceLayer),
						zap.String("function", logctx.LogDeallocateFaculty),
						zap.Error(err),
					)
					return fmt.Errorf("Internal server error")
				}
				s.logger.Info("Course Staff updated successfully",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogAllocateFaculty))
			} else {
				err := s.courseStaffRepo.DeleteStaff(ctx, tx, courseStaff.AssignmentID)
				if err != nil {
					s.logger.Error("Error deleting Course Staff",
						zap.String("layer", logctx.LogServiceLayer),
						zap.String("function", logctx.LogDeallocateFaculty),
						zap.Error(err),
					)
					return fmt.Errorf("Internal server error")
				}
				s.logger.Info("Course Staff deleted successfully",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogAllocateFaculty))
			}
		}
		load, err := s.workloadRepo.GetSemesterWorkloadByVersionID(ctx, profileVer.ProfileVersionId, inst.SemesterID)
		if err != nil {
			s.logger.Error("error getting Course Instance",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Int64("courseInstanceID", courseID),
				zap.Error(err))
			return fmt.Errorf("Internal server error")
		}
		if load == nil {
			s.logger.Info("No workload found for this course",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty),
				zap.Int64("courseInstanceID", inst.InstanceID))
			return fmt.Errorf("No workload found for this user and course")
		} else {
			s.logger.Info("workload found successfully",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAllocateFaculty))
			switch *positionType {
			case "PI":
				load.LecturesCount -= 15
			case "TI":
				load.TutorialsCount -= 15
			case "TA":
				load.LabsCount -= 15 * int64(*groupsAssigned)
			}
			err := s.workloadRepo.UpdateSemesterWorkload(ctx, tx, load)
			if err != nil {
				s.logger.Error("Error updating Workload",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogAllocateFaculty),
					zap.Error(err),
				)
				return fmt.Errorf("Internal server error")
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func positionOccupied(instance *courseInstance.CourseInstance, positionType string) bool {
	return *instance.PIAllocationStatus == courseInstance.Status("allocated") && positionType == "PI" ||
		*instance.TIAllocationStatus == courseInstance.Status("allocated") && positionType == "TI"
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
