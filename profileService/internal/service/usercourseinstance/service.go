package usercourseinstance

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/usercourseinstance"
	"go.uber.org/zap"
)

var _ usercourseinstance.Repository = (*Service)(nil)

type Service struct {
	repo   usercourseinstance.Repository
	logger *zap.Logger
}

func NewService(repo usercourseinstance.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

const (
	layer                   = "Service"
	getInstancesByProfileID = "GetInstancesByProfileID"
	addCourseInstance       = "AddCourseInstance"
)

func (s *Service) GetInstancesByProfileID(ctx context.Context, profileID int64) ([]int64, error) {
	if profileID <= 0 {
		s.logger.Error("profileID must be positive",
			zap.Int64("profileID", profileID),
		)
		return nil, fmt.Errorf("profileID must be positive. Id: %d", profileID)
	}
	ids, err := s.repo.GetInstancesByProfileID(ctx, profileID)
	if err != nil {
		s.logger.Error("Failed to get instances by profile_id",
			zap.String("layer", layer),
			zap.String("function", getInstancesByProfileID),
			zap.Int64("profile_id", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get instances by profile id %d: %v", profileID, err)
	}
	s.logger.Info("Successfully got instances by profile id",
		zap.String("layer", layer),
		zap.String("function", getInstancesByProfileID),
		zap.Int64("profile_id", profileID),
	)
	return ids, nil
}
func (s *Service) AddCourseInstance(ctx context.Context, instance *usercourseinstance.UserCourseInstance) error {
	if instance == nil {
		s.logger.Error("instance cannot be nil",
			zap.String("layer", layer),
			zap.String("function", addCourseInstance),
			zap.Any("instance", instance),
		)
		return fmt.Errorf("instance cannot be nil %v", instance)
	}
	if instance.CourseInstanceID <= 0 {
		s.logger.Error("instance.CourseInstanceID must be positive",
			zap.String("layer", layer),
			zap.String("function", addCourseInstance),
			zap.Int("instance_id", instance.CourseInstanceID))
		return fmt.Errorf("instance.CourseInstanceID must be positive. Id: %d", instance.CourseInstanceID)
	}
	if instance.ProfileID <= 0 {
		s.logger.Error("instance.ProfileID must be positive",
			zap.String("layer", layer),
			zap.String("function", addCourseInstance),
			zap.Int("instance_id", instance.ProfileID))
		return fmt.Errorf("instance.ProfileID must be positive. Id: %d", instance.ProfileID)
	}
	if instance.UserCourseID <= 0 {
		s.logger.Error("instance.UserCourseID must be positive",
			zap.String("layer", layer),
			zap.String("function", addCourseInstance),
			zap.Int("instance_id", instance.UserCourseID))
		return fmt.Errorf("instance.UserCourseID must be positive. Id: %d", instance.UserCourseID)
	}
	err := s.repo.AddCourseInstance(ctx, instance)
	if err != nil {
		s.logger.Error("Failed to add course instance",
			zap.String("layer", layer),
			zap.String("function", addCourseInstance),
			zap.Int("instance_id", instance.CourseInstanceID),
			zap.Error(err))
		return fmt.Errorf("failed to add course instance %d: %v", instance.CourseInstanceID, err)
	}
	return nil
}
