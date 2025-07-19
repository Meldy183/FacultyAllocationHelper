package courseInstance

import (
	"context"
	"fmt"

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
	if yearValid(courseInstance.Year) {
		s.logger.Error(
			"Invalid year",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid year: %v", courseInstance.Year)
	}
	if semesterIDValid(courseInstance.SemesterID) {
		s.logger.Error(
			"Invalid semester ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid year: %v", courseInstance.SemesterID)
	}
	if academicYearIDValid(courseInstance.AcademicYearID) {
		s.logger.Error(
			"Invalid academic year ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid academic year: %v", courseInstance.AcademicYearID)
	}
	if formValid(*courseInstance.Form) {
		s.logger.Error(
			"Invalid form",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid form: %v", courseInstance.Form)
	}
	if modeValid(*courseInstance.Mode) {
		s.logger.Error(
			"Invalid mode",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid mode: %v", courseInstance.Mode)
	}
	if statusValid(*courseInstance.PIAllocationStatus) {
		s.logger.Error(
			"Invalid PI allocation status",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid PI allocation status: %v", courseInstance.PIAllocationStatus)
	}
	if statusValid(*courseInstance.TIAllocationStatus) {
		s.logger.Error(
			"Invalid TI allocation status",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourseInstance),
		)
		return fmt.Errorf("invalid TI allocation status: %v", courseInstance.TIAllocationStatus)
	}
	err := s.repo.AddNewCourseInstance(ctx, courseInstance)
	if err != nil {
		s.logger.Error(
			"Invalid responsibleInstituteID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddNewCourse),
		)
		return fmt.Errorf("error creating new course: %v", err)
	}
	s.logger.Info("new course created",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddNewCourse),
	)
	return nil
}

func yearValid(year int) bool {
	return (year >= 2015)
}

func semesterIDValid(id int) bool {
	return (id >= 1 && id <= 3)
}

func academicYearIDValid(id int) bool {
	return (id >= 1 && id <= 8)
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

func (s *Service) GetCourseInstanceByID(ctx context.Context, id int64) (*courseInstance.CourseInstance, error) {
	CourseInstance, err := s.repo.GetCourseInstanceByID(ctx, id)
	if err != nil {
		s.logger.Error("error getting CourseInstance by ID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseInstanceByID),
			zap.Int64("CourseInstanceID", id),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course %w", err)
	}
	s.logger.Info("CourseInstance found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseInstanceByID),
		zap.Int64("CourseInstanceID", id),
		zap.Any("CourseInstance", CourseInstance),
	)
	return CourseInstance, nil
}

func (s *Service) UpdateCourseInstanceByID(ctx context.Context, course *courseInstance.CourseInstance) error {
	//TODO: implement me
	panic("implement me")
	return nil
}

func (s *Service) GetInstancesByInstitutes(ctx context.Context, instituteIDs []int64) ([]int64, error) {
	CourseInstances, err := s.repo.GetInstancesIDsByInstituteIDs(ctx, instituteIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by InstituteIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseInstanceByID),
			zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by InstituteIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseInstanceByID),
		zap.String("instituteIDs", fmt.Sprintf("%v", instituteIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesByAcademicYears(ctx context.Context, academicYearIDs []int64) ([]int64, error) {
	CourseInstances, err := s.repo.GetInstancesIDsByAcademicYearIDs(ctx, academicYearIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by academicYearIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseInstanceByID),
			zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by academicYearIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseInstanceByID),
		zap.String("academicYearIDs", fmt.Sprintf("%v", academicYearIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesBySemesters(ctx context.Context, semesterIDs []int64) ([]int64, error) {
	CourseInstances, err := s.repo.GetInstancesIDsBySemesterIDs(ctx, semesterIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by semesterIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseInstanceByID),
			zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by semesterIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseInstanceByID),
		zap.String("semesterIDs", fmt.Sprintf("%v", semesterIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}

func (s *Service) GetInstancesByPrograms(ctx context.Context, proframIDs []int64) ([]int64, error) {
	CourseInstances, err := s.repo.GetInstancesIDsByProgramIDs(ctx, proframIDs)
	if err != nil {
		s.logger.Error("error getting CourseInstances by proframIDs",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCourseInstanceByID),
			zap.String("proframIDs", fmt.Sprintf("%v", proframIDs)),
			zap.Error(err))
		return nil, fmt.Errorf("error getting course instances %w", err)
	}
	s.logger.Info("CourseInstances found by proframIDs",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCourseInstanceByID),
		zap.String("proframIDs", fmt.Sprintf("%v", proframIDs)),
		zap.String("CourseInstancesIDs", fmt.Sprintf("%v", CourseInstances)),
	)
	return CourseInstances, nil
}
