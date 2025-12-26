package track

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/track"
	trackcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/trackCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ track.Service = (*Service)(nil)

type Service struct {
	logger            *zap.Logger
	trackRepo         track.Repository
	trackInstanceRepo trackcourseinstance.Repository
}

func NewService(
	trackRepo track.Repository,
	trackInstanceRepo trackcourseinstance.Repository,
	logger *zap.Logger,
) *Service {
	return &Service{
		logger:            logger,
		trackRepo:         trackRepo,
		trackInstanceRepo: trackInstanceRepo,
	}
}

func (s *Service) GetAllTracks(ctx context.Context) ([]*track.Track, error) {
	tracks, err := s.trackRepo.GetAllTracks(ctx)
	if err != nil {
		s.logger.Error("Error getting all tracks",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetAllTracks),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting all tracks: %w", err)
	}
	return tracks, nil
}
func (s *Service) GetTrackNameByID(ctx context.Context, trackID uuid.UUID) (*string, error) {
	trackName, err := s.trackRepo.GetTrackNameByID(ctx, trackID)
	if err != nil {
		s.logger.Error("Error getting track name",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetTrackNameByID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting track name: %w", err)
	}
	return trackName, nil
}
func (s *Service) GetTrackIDByName(ctx context.Context, trackName string) (*uuid.UUID, error) {
	trackID, err := s.trackRepo.GetTrackIDByName(ctx, trackName)
	if err != nil {
		s.logger.Error("Error getting track name",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetTrackNameByID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting track name: %w", err)
	}
	return trackID, nil
}
func (s *Service) GetTracksOfCourseByInstanceID(ctx context.Context, instanceID uuid.UUID) ([]uuid.UUID, error) {
	tracks, err := s.trackInstanceRepo.GetTracksIDsOfCourseByInstanceID(ctx, instanceID)
	if err != nil {
		s.logger.Error("Error getting all tracks",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetTracksOfCourseByCourseID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting all tracks: %w", err)
	}
	return tracks, nil
}

func (s *Service) GetTracksNamesOfCourseByCourseInstanceID(ctx context.Context, instanceID uuid.UUID) ([]*string, error) {
	trackIds, err := s.GetTracksOfCourseByInstanceID(ctx, instanceID)
	if err != nil {
		s.logger.Error("Error getting all tracks",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetTracksOfCourseByCourseID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting all tracks: %w", err)
	}
	tracksNames := make([]*string, 0)
	for _, trackId := range trackIds {
		trackName, err := s.trackRepo.GetTrackNameByID(ctx, trackId)
		if err != nil {
			s.logger.Error("Error getting track name",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogGetTracksOfCourseByCourseID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error getting track name: %w", err)
		}
		tracksNames = append(tracksNames, trackName)
	}
	s.logger.Info("Successfully got all tracks",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetTracksOfCourseByCourseID),
		zap.Any("tracks", tracksNames),
	)
	return tracksNames, nil
}
