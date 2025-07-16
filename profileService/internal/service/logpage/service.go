package logpage

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/logpage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ logpage.Repository = (*Service)(nil)

type Service struct {
	repo   logpage.Repository
	logger *zap.Logger
}

func NewService(repo logpage.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}
func (s *Service) GetLogpages(ctx context.Context, lastID string, limit string) ([]*logpage.LogPage, error) {
	logpages, err := s.repo.GetLogpages(ctx, lastID, limit)
	if err != nil {
		s.logger.Error("Error getting logpages",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetLogpages),
			zap.Error(err))
		return nil, err
	}
	s.logger.Info("Succesfully got logpages",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetLogpages))
	return logpages, err
}
func (s *Service) AddLogpage(ctx context.Context, logpage *logpage.LogPage) error {
	err := s.repo.AddLogpage(ctx, logpage)
	if err != nil {
		s.logger.Error("Error adding logpage",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddLogpage),
			zap.Error(err))
		return err
	}
	s.logger.Info("Succesfully added logpage",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddLogpage))
	return err
}
