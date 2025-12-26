package courseInstance

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ courseInstance.Service = (*Service)(nil)

type Service struct {
	repo   courseInstance.Repository
	logger *zap.Logger
}

func NewService(repo courseInstance.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) AddCourseInstance(ctx context.Context, courseInstance *courseInstance.CourseInstance) error {
	if !yearValid(courseInstance.Year) {
		s.logger.Error(
			"Invalid year",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
			zap.Int64("year", courseInstance.Year),
		)
		return fmt.Errorf("invalid year: %v", courseInstance.Year)
	}
	if !semesterIDValid(courseInstance.SemesterID) {
		s.logger.Error(
			"Invalid semester ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid semester: %v", courseInstance.SemesterID)
	}
	if !academicYearIDValid(courseInstance.AcademicYearID) {
		s.logger.Error(
			"Invalid academic year ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid academic year: %v", courseInstance.AcademicYearID)
	}
	s.logger.Info("academic year ok")
	if courseInstance.Form != nil && !formValid(*courseInstance.Form) {
		s.logger.Error(
			"Invalid form",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid form: %v", courseInstance.Form)
	}
	s.logger.Info("form ok")

	if courseInstance.Mode != nil && !modeValid(*courseInstance.Mode) {
		s.logger.Error(
			"Invalid mode",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid mode: %v", courseInstance.Mode)
	}
	s.logger.Info("mode ok")

	if courseInstance.PIAllocationStatus != nil && !statusValid(*courseInstance.PIAllocationStatus) {
		s.logger.Error(
			"Invalid PI allocation status",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid PI allocation status: %v", courseInstance.PIAllocationStatus)
	}
	s.logger.Info("pi status ok")

	if courseInstance.TIAllocationStatus != nil && !statusValid(*courseInstance.TIAllocationStatus) {
		s.logger.Error(
			"Invalid TI allocation status",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid TI allocation status: %v", courseInstance.TIAllocationStatus)
	}
	s.logger.Info("ti status ok")

	err := s.repo.AddNewCourseInstance(ctx, courseInstance)
	if err != nil {
		s.logger.Error(
			"Invalid responsibleInstituteID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourse),
		)
		return fmt.Errorf("error creating new course: %w", err)
	}
	s.logger.Info("new course created",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddNewCourse),
	)
	return nil
}

func yearValid(year int64) bool {
	return year >= 2015
}

func semesterIDValid(id uuid.UUID) bool {
	return id != uuid.Nil
}

func academicYearIDValid(id uuid.UUID) bool {
	return id != uuid.Nil
}

func formValid(form courseInstance.Form) bool {
	switch form {
	case courseInstance.FormBlock1, courseInstance.FormBlock2, courseInstance.FormFull:
		return true
	default:
		return false
	}
}

func modeValid(mode courseInstance.Mode) bool {
	switch mode {
	case courseInstance.ModeOnsite, courseInstance.ModeMixed, courseInstance.ModeRemote:
		return true
	default:
		return false
	}
}

func statusValid(mode courseInstance.Status) bool {
	switch mode {
	case courseInstance.StatusAllocated, courseInstance.StatusNotNeeded, courseInstance.StatusNotAllocated:
		return true
	default:
		return false
	}
}

func (s *Service) GetCourseInstanceByID(ctx context.Context, id uuid.UUID) (*courseInstance.CourseInstance, error) {
	CourseInstance, err := s.repo.GetCourseInstanceByID(ctx, id)
	if err != nil {
		s.logger.Error("error getting CourseInstance by ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseInstanceByID),
			zap.String("CourseInstanceID", id.String()),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course %w", err)
	}
	s.logger.Info("CourseInstance found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseInstanceByID),
		zap.String("CourseInstanceID", id.String()),
		zap.Any("CourseInstance", CourseInstance),
	)
	return CourseInstance, nil
}

func (s *Service) UpdateCourseInstanceByID(
	ctx context.Context,
	id uuid.UUID,
	courseInstance *courseInstance.CourseInstance,
) error {
	if !semesterIDValid(courseInstance.SemesterID) {
		s.logger.Error(
			"Invalid semester ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
		)
		return fmt.Errorf("invalid semester: %v", courseInstance.SemesterID)
	}
	if !academicYearIDValid(courseInstance.AcademicYearID) {
		s.logger.Error(
			"Invalid academic year ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
		)
		return fmt.Errorf("invalid academic year: %v", courseInstance.AcademicYearID)
	}
	s.logger.Info("academic year ok")
	if courseInstance.Form != nil && !formValid(*courseInstance.Form) {
		s.logger.Error(
			"Invalid form",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
		)
		return fmt.Errorf("invalid form: %v", courseInstance.Form)
	}
	s.logger.Info("form ok")

	if courseInstance.Mode != nil && !modeValid(*courseInstance.Mode) {
		s.logger.Error(
			"Invalid mode",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
		)
		return fmt.Errorf("invalid mode: %v", courseInstance.Mode)
	}
	s.logger.Info("mode ok")

	if courseInstance.PIAllocationStatus != nil && !statusValid(*courseInstance.PIAllocationStatus) {
		s.logger.Error(
			"Invalid PI allocation status",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
		)
		return fmt.Errorf("invalid PI allocation status: %v", courseInstance.PIAllocationStatus)
	}
	s.logger.Info("pi status ok")

	if courseInstance.TIAllocationStatus != nil && !statusValid(*courseInstance.TIAllocationStatus) {
		s.logger.Error(
			"Invalid TI allocation status",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
		)
		return fmt.Errorf("invalid TI allocation status: %v", courseInstance.TIAllocationStatus)
	}
	s.logger.Info("ti status ok")

	err := s.repo.UpdateCourseInstanceByID(ctx, id, courseInstance)
	if err != nil {
		s.logger.Error(
			"could not update course instance by id",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateCourseInstanceByID),
		)
		return fmt.Errorf("error updating course instance: %w", err)
	}
	s.logger.Info("course updated",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogUpdateCourseInstanceByID),
		zap.String("id", id.String()),
	)
	return nil
}

func (s *Service) GetInstancesByInstituteIDs(ctx context.Context, instituteIDs []uuid.UUID) ([]uuid.UUID, error) {
	var actualIDs []uuid.UUID
	if len(instituteIDs) == 0 {
		actualIDs = []uuid.UUID{}
	} else {
		actualIDs = instituteIDs
	}
	CourseInstances, err := s.repo.GetInstancesIDsByInstituteIDs(ctx, actualIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by InstituteIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByInstituteIDs),
			zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
			zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by InstituteIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByInstituteIDs),
		zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
		zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesByAcademicYearIDs(ctx context.Context, academicYearIDs []uuid.UUID) ([]uuid.UUID, error) {
	var actualIDs []uuid.UUID
	if len(academicYearIDs) == 0 {
		actualIDs = []uuid.UUID{}
	} else {
		actualIDs = academicYearIDs
	}
	CourseInstances, err := s.repo.GetInstancesIDsByAcademicYearIDs(ctx, actualIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by academicYearIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByAcademicYearIDs),
			zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
			zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by academicYearIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByAcademicYearIDs),
		zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
		zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesBySemesterIDs(ctx context.Context, semesterIDs []uuid.UUID) ([]uuid.UUID, error) {
	var actualIDs []uuid.UUID
	if len(semesterIDs) == 0 {
		actualIDs = []uuid.UUID{}
	} else {
		actualIDs = semesterIDs
	}
	CourseInstances, err := s.repo.GetInstancesIDsBySemesterIDs(ctx, actualIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by semesterIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesBySemesterIDs),
			zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
			zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by semesterIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesBySemesterIDs),
		zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
		zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesByProgramIDs(ctx context.Context, programIDs []uuid.UUID) ([]uuid.UUID, error) {
	var actualIDs []uuid.UUID
	if len(programIDs) == 0 {
		actualIDs = []uuid.UUID{}
	} else {
		actualIDs = programIDs
	}
	CourseInstances, err := s.repo.GetInstancesIDsByProgramIDs(ctx, actualIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by programIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByProgramIDs),
			zap.String("programIDs", fmt.Sprintf("%v", programIDs)),
			zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by programIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByProgramIDs),
		zap.String("programIDs", fmt.Sprintf("%v", programIDs)),
		zap.String("actualIDs", fmt.Sprintf("%v", actualIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesByAllocationStatus(ctx context.Context, allocNotFinished bool) ([]uuid.UUID, error) {
	var CourseInstances []uuid.UUID
	var err error
	if allocNotFinished {
		CourseInstances, err = s.repo.GetInstancesByAllocationStatus(ctx)
	} else {
		CourseInstances, err = s.repo.GetAllInstancesIDs(ctx)
	}
	if err != nil {
		s.logger.Error("error getting CourseInstances by AllocationStatus",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByAllocationStatus),
			zap.String("allocNotFinished", strconv.FormatBool(allocNotFinished)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by AllocationStatus",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByAllocationStatus),
		zap.String("allocNotFinished", strconv.FormatBool(allocNotFinished)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil //
}

func (s *Service) GetInstancesByYear(ctx context.Context, year int64) ([]uuid.UUID, error) {
	CourseInstances, err := s.repo.GetInstancesByYear(ctx, year)
	if err != nil {
		s.logger.Error("error getting CourseInstances by year",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByYear),
			zap.Int64("year", year),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by year",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByYear),
		zap.Int64("year", year),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesByVersionID(ctx context.Context, versionID uuid.UUID) ([]uuid.UUID, error) {
	CourseInstances, err := s.repo.GetInstancesByVersionID(ctx, versionID)
	if err != nil {
		s.logger.Error("error getting CourseInstances by versionID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByVersionID),
			zap.String("versionID", versionID.String()),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by versionID",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByVersionID),
		zap.String("versionID", versionID.String()),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}
