package track

import (
	"context"
	"fmt"
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

func NewService(trackRepo track.Repository, trackInstanceRepo trackcourseinstance.Repository, logger *zap.Logger) *Service {
	return &Service{
		logger:            logger,
		trackRepo:         trackRepo,
		trackInstanceRepo: trackInstanceRepo,
	}
}

func (s *Service) GetAllTracks(ctx context.Context) (*[]*track.Track, error) {
	tracks, err := s.trackRepo.GetAllTracks(ctx)
	if err != nil {
		s.logger.Error("Error getting all tracks",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetAllTracks),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting all tracks: %w", err)
	}
	return &tracks, nil
}
func (s *Service) GetTrackNameByID(ctx context.Context, trackID int64) (*string, error) {
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
func (s *Service) GetTracksOfCourseByInstanceID(ctx context.Context, instanceID int) ([]*trackcourseinstance.TrackCourseInstance, error) {
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
func (s *Service) ConvertTrackCourseInstanceToTrackNames(ctx context.Context, linkArray []*trackcourseinstance.TrackCourseInstance) ([]*string, error) {
	trackNames := make([]*string, 0)
	for _, elem := range linkArray {
		name, err := s.GetTrackNameByID(ctx, int64(elem.TrackID))
		if err != nil {
			s.logger.Error("Error getting track name",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogGetTrackNameByID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("error getting track name: %w", err)
		}
		trackNames = append(trackNames, name)
	}
	return trackNames, nil
}

func (s *Service) GetTracksNamesOfCourseByCourseID(ctx context.Context, instanceID int) ([]*string, error) {
	trackIds, err := s.GetTracksOfCourseByInstanceID(ctx, instanceID)
	if err != nil {
		s.logger.Error("Error getting all tracks",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetTracksOfCourseByCourseID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting all tracks: %w", err)
	}
	names, err := s.ConvertTrackCourseInstanceToTrackNames(ctx, trackIds)
	if err != nil {
		s.logger.Error("Error converting track course instance to track names",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetTracksOfCourseByCourseID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error converting track course instance to track names: %w", err)
	}
	return names, nil
}
