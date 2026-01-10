package profileVersion

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileVersion.Service = (*Service)(nil)

type Service struct {
	pool   *pgxpool.Pool // MOCK to avoid import error
	repo   profileVersion.Repository
	logger *zap.Logger
}

func NewService(repo profileVersion.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetVersionByProfileID(
	ctx context.Context,
	profileID int64,
	year int64,
) (*profileVersion.ProfileVersion, error) {
	return s.repo.GetVersionByProfileID(ctx, profileID, year)
}

func (s *Service) GetVersionByVersionID(ctx context.Context, versionID int64) (*profileVersion.ProfileVersion, error) {
	return s.repo.GetVersionByVersionID(ctx, versionID)
}

func (s *Service) GetVersionIDByProfileID(ctx context.Context, profileID int64, year int64) (int64, error) {
	version, err := s.GetVersionByProfileID(ctx, profileID, year)
	if err != nil {
		s.logger.Error("Failed to get profile version",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetVersionIDByProfileID),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return 0, err
	}
	return version.ProfileVersionId, nil
}

func (s *Service) AddProfileVersion(ctx context.Context, version *profileVersion.ProfileVersion) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		s.logger.Error("error starting transaction",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddSemesterWorkload),
			zap.Error(err))
		return err
	}
	defer func() {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			s.logger.Error("error rolling back transaction",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddSemesterWorkload),
				zap.Error(rollbackErr))
		}
	}()
	if err = s.repo.AddProfileVersion(ctx, &tx, version); err != nil {
		s.logger.Error("Failed to add profile version",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddProfileVersion),
			zap.Int64("profileID", version.ProfileVersionId),
			zap.Error(err),
		)
		return fmt.Errorf("failed to add profile version: %w", err)
	}
	return nil
}
