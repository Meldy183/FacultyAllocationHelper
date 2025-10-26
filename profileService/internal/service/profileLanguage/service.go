package profileLanguage

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileLanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileLanguage.Service = (*Service)(nil)

type Service struct {
	repo   profileLanguage.Repository
	logger *zap.Logger
}

func NewService(repo profileLanguage.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) AddUserLanguage(ctx context.Context, userLanguage *profileLanguage.ProfileLanguage) error {
	err := s.repo.AddUserLanguage(ctx, userLanguage)
	if err != nil {
		s.logger.Error("User Institute Added to DB",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddUserLanguage))
		zap.Error(err)
		return fmt.Errorf("failed to add user language %w", err)
	}
	s.logger.Info("User Institute Added to DB",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddUserLanguage),
	)
	return nil
}
func (s *Service) GetProfileLanguages(ctx context.Context, profileID int64) ([]string, error) {
	if profileID <= 0 {
		s.logger.Error("profileID must be positive",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetUserLanguages),
		)
		return nil, fmt.Errorf("profileID must be positive: %d", profileID)
	}
	languages, err := s.repo.GetProfileLanguages(ctx, profileID)
	if err != nil {
		s.logger.Error("failed to get user languages",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetUserLanguages),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get user languages %w", err)
	}
	s.logger.Info("User Institute Added to DB",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetUserLanguages),
		zap.Int64("profileID", profileID),
	)
	return languages, nil
}
