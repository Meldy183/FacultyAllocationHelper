package completeCourse

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/program"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/track"
	"go.uber.org/zap"
)

var _ CompleteCourse.Service = (*Service)(nil)

type Service struct {
	logger          *zap.Logger
	instanceService *courseInstance.Service
	courseService   *course.Service
	trackService    *track.Service
	program         *program.Service
}

func NewService(instance *courseInstance.Service, course *course.Service, track *track.Service, program *program.Service, logger *zap.Logger) *Service {
	return &Service{
		instanceService: instance,
		courseService:   course,
		trackService:    track,
		program:         program,
		logger:          logger,
	}
}

func (s *Service) GetFullCourseInfoByID(ctx context.Context, instanceID int64) (*CompleteCourse.FullCourse, error) {
	courseInstanceObj, err := s.instanceService.GetCourseInstanceByID(ctx, instanceID)
	if err != nil {
		s.logger.Error("failed to fetch courseInstanceObj",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetFullCourseInfoByID),
			zap.Int64("instanceID", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to fetch courseInstanceObj: %w", err)
	}
	courseObj, err := s.courseService.GetCourseByID(ctx, courseInstanceObj.CourseID)
	if err != nil {
		s.logger.Error("failed to fetch courseObj",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetFullCourseInfoByID),
			zap.Int64("instanceID", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to fetch courseObj: %w", err)
	}
	trackNames, err := s.trackService.GetTracksNamesOfCourseByCourseInstanceID(ctx, instanceID)
	if err != nil {
		s.logger.Error("failed to fetch trackNames",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetFullCourseInfoByID),
			zap.Int64("instanceID", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to fetch trackNames: %w", err)
	}
	studyProgramNames, err := s.program.GetProgramNamesByInstanceID(ctx, instanceID)
	if err != nil {
		s.logger.Error("failed to fetch studyProgramNames",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetFullCourseInfoByID),
			zap.Int64("instanceID", instanceID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to fetch studyProgramNames: %w", err)
	}
	fullCourse := &CompleteCourse.FullCourse{
		Course:         *courseObj,
		CourseInstance: *courseInstanceObj,
		Tracks:         trackNames,
		StudyPrograms:  studyProgramNames,
	}
	s.logger.Info("successfully fetched full course",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetFullCourseInfoByID),
		zap.Int64("instanceID", instanceID),
		zap.Int64("courseID", courseObj.CourseID),
		zap.Any("fullCourse", fullCourse),
	)
	return fullCourse, nil
}
func (s *Service) AddFullCourse(ctx context.Context, fullcourse *CompleteCourse.FullCourse) error {
	err := s.courseService.AddCourse(ctx, &fullcourse.Course)
	if err != nil {
		s.logger.Error("Error adding course",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddFullCourse),
			zap.Error(err),
		)
		return err
	}
	s.logger.Info("Added course successfully")
	err = s.instanceService.AddCourseInstance(ctx, &fullcourse.CourseInstance)
	if err != nil {
		s.logger.Error("Error adding courseInstance",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddFullCourse),
			zap.Error(err),
		)
		return err
	}
	s.logger.Info("Added course instance successfully")
	for _, track := range fullcourse.Tracks {
		ID, err := s.trackService.GetTrackIDByName(ctx, *track)
		if err != nil {
			s.logger.Error("error seeking track by name",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddFullCourse),
				zap.Error(err))
			continue
		}
	}
}
