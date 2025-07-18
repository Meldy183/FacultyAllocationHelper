package profileCourseInstance

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileCourseInstance.Repository = (*Service)(nil)

type Service struct {
	repo   profileCourseInstance.Repository
	logger *zap.Logger
}

func NewService(repo profileCourseInstance.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetCourseInstancesByProfileID(ctx context.Context, profileVersionID int64) ([]int64, error) {
	if profileVersionID <= 0 {
		s.logger.Error("profileID must be positive",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByProfileID),
			zap.Int64("profileID", profileVersionID),
		)
		return nil, fmt.Errorf("profileID must be positive. Id: %d", profileVersionID)
	}
	ids, err := s.repo.GetCourseInstancesByProfileID(ctx, profileVersionID)
	if err != nil {
		s.logger.Error("Failed to get instances by profile_id",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByProfileID),
			zap.Int64("profile_id", profileVersionID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get instances by facultyProfile id %d: %v", profileVersionID, err)
	}
	s.logger.Info("Successfully got instances by facultyProfile id",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByProfileID),
		zap.Int64("profile_id", profileVersionID),
	)
	return ids, nil
}

func (s *Service) AddCourseInstance(ctx context.Context, instance *profileCourseInstance.ProfileCourseInstance) error {
	if instance == nil {
		s.logger.Error("instance cannot be nil",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.Any("instance", instance),
		)
		return fmt.Errorf("instance cannot be nil %v", instance)
	}
	if instance.CourseInstanceID <= 0 {
		s.logger.Error("instance.CourseInstanceID must be positive",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.Int("instance_id", instance.CourseInstanceID),
		)
		return fmt.Errorf("instance.CourseInstanceID must be positive. Id: %d", instance.CourseInstanceID)
	}
	if instance.ProfileVersionID <= 0 {
		s.logger.Error("instance.ProfileVersionID must be positive",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.Int("instance_id", instance.ProfileVersionID),
		)
		return fmt.Errorf("instance.ProfileVersionID must be positive. Id: %d", instance.ProfileVersionID)
	}
	if instance.ProfileCourseID <= 0 {
		s.logger.Error("instance.ProfileCourseID must be positive",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.Int("instance_id", instance.ProfileCourseID),
		)
		return fmt.Errorf("instance.ProfileCourseID must be positive. Id: %d", instance.ProfileCourseID)
	}
	err := s.repo.AddCourseInstance(ctx, instance)
	if err != nil {
		s.logger.Error("Failed to add course instance",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.Int("instance_id", instance.CourseInstanceID),
			zap.Error(err),
		)
		return fmt.Errorf("failed to add course instance %d: %v", instance.CourseInstanceID, err)
	}
	return nil
}
