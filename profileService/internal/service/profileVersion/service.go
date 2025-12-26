package profileVersion

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileVersion.Service = (*Service)(nil)

type Service struct {
	repo   profileVersion.Repository
	logger *zap.Logger
}

func NewService(repo profileVersion.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetVersionByProfileID(
	ctx context.Context,
	profileID uuid.UUID,
	year int64,
) (*profileVersion.ProfileVersion, error) {
	return s.repo.GetVersionByProfileID(ctx, profileID, year)
}

func (s *Service) GetVersionByVersionID(ctx context.Context, versionID uuid.UUID) (*profileVersion.ProfileVersion, error) {
	return s.repo.GetVersionByVersionID(ctx, versionID)
}

func (s *Service) GetVersionIDByProfileID(ctx context.Context, profileID uuid.UUID, year int64) (uuid.UUID, error) {
	version, err := s.GetVersionByProfileID(ctx, profileID, year)
	if err != nil {
		s.logger.Error("Failed to get profile version",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetVersionIDByProfileID),
			zap.String("profileID", profileID.String()),
			zap.Error(err),
		)
		return uuid.Nil, err
	}
	return version.ProfileVersionId, nil
}

func (s *Service) AddProfileVersion(ctx context.Context, version *profileVersion.ProfileVersion) error {
	if err := s.repo.AddProfileVersion(ctx, version); err != nil {
		s.logger.Error("Failed to add profile version",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddProfileVersion),
			zap.String("profileID", version.ProfileVersionId.String()),
			zap.Error(err),
		)
		return fmt.Errorf("failed to add profile version: %w", err)
	}
	return nil
}
