package profileCourseInstance

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileCourseInstance.Service = (*Service)(nil)

type Service struct {
	repo   profileCourseInstance.Repository
	logger *zap.Logger
}

func NewService(repo profileCourseInstance.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetCourseInstancesByVersionID(ctx context.Context, profileVersionID uuid.UUID) ([]uuid.UUID, error) {
	if profileVersionID == uuid.Nil {
		s.logger.Error("profileID must not be nil",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByProfileID),
			zap.String("profileID", profileVersionID.String()),
		)
		return nil, fmt.Errorf("profileID must not be nil. Id: %s", profileVersionID.String())
	}
	ids, err := s.repo.GetCourseInstancesByVersionID(ctx, profileVersionID)
	if err != nil {
		s.logger.Error("Failed to get instances by profile_id",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstancesByProfileID),
			zap.String("profile_id", profileVersionID.String()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get instances by facultyProfile id %s: %w", profileVersionID.String(), err)
	}
	s.logger.Info("Successfully got instances by facultyProfile id",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstancesByProfileID),
		zap.String("profile_id", profileVersionID.String()),
	)
	return ids, nil
}

func (s *Service) AddCourseInstance(
	ctx context.Context,
	instance *profileCourseInstance.ProfileVersionCourseInstance,
) error {
	if instance == nil {
		s.logger.Error("instance cannot be nil",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.Any("instance", instance),
		)
		return fmt.Errorf("instance cannot be nil %v", instance)
	}
	if instance.CourseInstanceID == uuid.Nil {
		s.logger.Error("instance.CourseInstanceID must not be nil",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.String("instance_id", instance.CourseInstanceID.String()),
		)
		return fmt.Errorf("instance.CourseInstanceID must not be nil. Id: %s", instance.CourseInstanceID.String())
	}
	if instance.ProfileVersionID == uuid.Nil {
		s.logger.Error("instance.ProfileVersionID must not be nil",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.String("instance_id", instance.ProfileVersionID.String()),
		)
		return fmt.Errorf("instance.ProfileVersionID must not be nil. Id: %s", instance.ProfileVersionID.String())
	}
	if instance.ProfileCourseID == uuid.Nil {
		s.logger.Error("instance.ProfileCourseID must not be nil",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.String("instance_id", instance.ProfileCourseID.String()),
		)
		return fmt.Errorf("instance.ProfileCourseID must not be nil. Id: %s", instance.ProfileCourseID.String())
	}
	err := s.repo.AddCourseInstance(ctx, instance)
	if err != nil {
		s.logger.Error("Failed to add course instance",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddCourseInstance),
			zap.String("instance_id", instance.CourseInstanceID.String()),
			zap.Error(err),
		)
		return fmt.Errorf("failed to add course instance %s: %w", instance.CourseInstanceID.String(), err)
	}
	return nil
}
