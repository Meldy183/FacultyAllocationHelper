package userlanguage

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userlanguage"
	"go.uber.org/zap"
)

var _ userlanguage.Repository = (*Service)(nil)

type Service struct {
	repo   userlanguage.Repository
	logger *zap.Logger
}

func NewService(repo userlanguage.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

const (
	logLayer            = "service"
	logAdd              = "Add"
	logGetUserLanguages = "getUserLanguages"
)

func (s *Service) Add(ctx context.Context, userLanguage *userlanguage.UserLanguage) error {
	err := s.repo.Add(ctx, userLanguage)
	if err != nil {
		s.logger.Error("User Institute Added to DB",
			zap.String(logLayer, logLayer),
			zap.String("function", logAdd))
		zap.Error(err)
		return fmt.Errorf("failed to add user language %w", err)
	}
	s.logger.Info("User Institute Added to DB",
		zap.String(logLayer, logLayer),
		zap.String("function", logAdd),
	)
	return nil
}
func (s *Service) GetUserLanguages(ctx context.Context, profileID int64) ([]*language.Language, error) {
	if profileID <= 0 {
		s.logger.Error("profileID must be positive",
			zap.String(logLayer, logLayer),
			zap.String("function", logGetUserLanguages),
		)
		return nil, fmt.Errorf("profileID must be positive: %d", profileID)
	}
	languages, err := s.repo.GetUserLanguages(ctx, profileID)
	if err != nil {
		s.logger.Error("failed to get user languages",
			zap.String(logLayer, logLayer),
			zap.String("function", logGetUserLanguages),
			zap.Int64("profileID", profileID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get user languages %w", err)
	}
	s.logger.Info("User Institute Added to DB",
		zap.String(logLayer, logLayer),
		zap.String("function", logGetUserLanguages),
		zap.Int64("profileID", profileID),
	)
	return languages, nil
}
