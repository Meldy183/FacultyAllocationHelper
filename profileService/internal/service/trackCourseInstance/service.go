package trackcourseinstance

import (
	"context"
	"fmt"

	trackcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/trackCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ trackcourseinstance.Service = (*Service)(nil)

type Service struct {
	logger            *zap.Logger
	trackInstanceRepo trackcourseinstance.Repository
}

func (s *Service) AddTracksToCourseInstance(ctx context.Context, instanceID int64, trackID int64) error {
	return s.trackInstanceRepo.AddTracksToCourseInstance(ctx, instanceID, trackID)
}

func NewService(trackInstanceRepo trackcourseinstance.Repository, logger *zap.Logger) *Service {
	return &Service{
		logger:            logger,
		trackInstanceRepo: trackInstanceRepo,
	}
}

func (s *Service) GetTracksIDsOfCourseByInstanceID(ctx context.Context, instanceID int64) ([]int64, error) {
	trackCourseInstances, err := s.trackInstanceRepo.GetTracksIDsOfCourseByInstanceID(ctx, instanceID)
	if err != nil {
		s.logger.Error("Error getting track course instances by instance ids",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetTrackCourseByCourseID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting track course instances by instance ids: %w", err)
	}
	s.logger.Info("successfully got track course instances by instance ids",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetTrackCourseByCourseID),
	)
	return trackCourseInstances, nil
}
